package install

import (
	"fmt"
	"os"

	"github.com/solo-io/gloo/pkg/cliutil"
	"github.com/solo-io/gloo/pkg/cliutil/install"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
)

func UninstallGloo(opts *options.Options, cli install.KubeCli) error {
	if err := uninstallGloo(opts, cli); err != nil {
		fmt.Fprintf(os.Stderr, "Uninstall failed. Detailed logs available at %s.\n", cliutil.GetLogsPath())
		return err
	}
	return nil
}

func uninstallGloo(opts *options.Options, cli install.KubeCli) error {
	if opts.Uninstall.DeleteNamespace || opts.Uninstall.DeleteAll {
		deleteNamespace(cli, opts.Uninstall.Namespace)
	} else {
		deleteGlooSystem(cli, opts.Uninstall.Namespace)
	}

	if opts.Uninstall.DeleteCrds || opts.Uninstall.DeleteAll {
		deleteGlooCrds(cli)
	}

	if opts.Uninstall.DeleteAll {
		deleteRbac(cli)
	}

	uninstallKnativeIfNecessary()

	return nil
}

func deleteRbac(cli install.KubeCli) {
	fmt.Printf("Removing Gloo RBAC configuration...\n")
	for _, rbacKind := range GlooRbacKinds {
		if err := cli.Kubectl(nil, "delete", rbacKind, "-l", "app=gloo"); err != nil {
			fmt.Printf("Deleting rbac %s failed\n", rbacKind)
		}
	}
}

func deleteGlooSystem(cli install.KubeCli, namespace string) {
	fmt.Printf("Removing Gloo system components from namespace %s...\n", namespace)
	for _, kind := range GlooSystemKinds {
		if err := cli.Kubectl(nil, "delete", kind, "-l", "app=gloo", "-n", namespace); err != nil {
			fmt.Printf("Deleting gloo system %s failed\n", kind)
		}
	}
}

func deleteGlooCrds(cli install.KubeCli) {
	fmt.Printf("Removing Gloo CRDs...\n")
	args := []string{"delete", "crd"}
	for _, crd := range GlooCrdNames {
		args = append(args, crd)
	}
	if err := cli.Kubectl(nil, args...); err != nil {
		fmt.Printf("Deleting Gloo CRDs failed\n")
	}
}

func deleteNamespace(cli install.KubeCli, namespace string) {
	fmt.Printf("Removing namespace %s...\n", namespace)
	if err := cli.Kubectl(nil, "delete", "namespace", namespace); err != nil {
		fmt.Printf("Deleting namespace %s failed\n", namespace)
	}
}

func uninstallKnativeIfNecessary() {
	_, installOpts, err := checkKnativeInstallation()
	if err != nil {
		fmt.Printf("Finding knative installation\n")
		return
	}
	if installOpts != nil {
		fmt.Printf("Removing knative components installed by Gloo %#v...\n", installOpts)
		manifests, err := RenderKnativeManifests(*installOpts)
		if err != nil {
			fmt.Printf("Rendering knative manifests\n")
			return
		}
		if err := install.KubectlDelete([]byte(manifests), "--ignore-not-found"); err != nil {
			fmt.Printf("Deleting knative failed\n")
		}
	}
}
