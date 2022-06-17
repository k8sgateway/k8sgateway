package register

import (
	"context"
	"fmt"
	"os"

	"github.com/solo-io/gloo/pkg/cliutil"
	linkedversion "github.com/solo-io/gloo/pkg/version"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/install"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/version"
	"github.com/solo-io/skv2/pkg/multicluster/kubeconfig"
	"github.com/solo-io/skv2/pkg/multicluster/register"
	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func installCrdsToRemote(context string) error {
	helmClient := install.DefaultHelmClient()

	chartObj, err := helmClient.DownloadChart("https://storage.googleapis.com/solo-public-helm/charts/gloo-" + linkedversion.Version + ".tgz")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "\nGloo failed to install CRDs! Detailed logs available at %s.\n", cliutil.GetLogsPath())
		return err
	}
	chartObj.Templates = nil // explicitly remove teamplates, since we only care about installing CRDs

	helmInstall, _, err := helmClient.NewInstall("default", "gloo-automatic-crd-application", false, context)
	if err != nil {
		return err
	}

	_, err = helmInstall.Run(chartObj, nil)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "\nGloo failed to install CRDs! Detailed logs available at %s.\n", cliutil.GetLogsPath())
		return err
	}
	return nil
}

func Register(opts *options.Options) error {
	ctx := context.Background()
	registerOpts := opts.Cluster.Register
	mgmtKubeCfg, err := kubeconfig.GetClientConfigWithContext("", "", "")
	if err != nil {
		return err
	}
	remoteKubeCfg, err := kubeconfig.GetClientConfigWithContext(registerOpts.RemoteKubeConfig, registerOpts.RemoteContext, "")
	if err != nil {
		return err
	}

	// check to see if Gloo is installed onto RemoteContext.  If not, install CRDs to prevent a gloo-fed crash, per
	// https://github.com/solo-io/gloo/issues/5832
	serverVersion, err := version.NewKube(opts.Metadata.GetNamespace(), registerOpts.RemoteContext).Get(ctx)
	if err != nil {
		return err
	}
	if serverVersion == nil {
		fmt.Printf("No `gloo` install detected on %s.  Installing OSS CRDs.", registerOpts.RemoteContext)
		installCrdsToRemote(registerOpts.RemoteContext)
	}

	clusterRegisterOpts := register.RegistrationOptions{
		KubeCfg:          mgmtKubeCfg,
		RemoteKubeCfg:    remoteKubeCfg,
		RemoteCtx:        registerOpts.RemoteContext,
		APIServerAddress: registerOpts.LocalClusterDomainOverride,
		ClusterName:      registerOpts.ClusterName,
		Namespace:        opts.Cluster.FederationNamespace,
		RemoteNamespace:  registerOpts.RemoteNamespace,
		ClusterRoles: []*v1.ClusterRole{
			{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: registerOpts.RemoteNamespace,
					Name:      "gloo-federation-controller",
				},
				Rules: glooFederationPolicyRules,
			},
		},
		Roles: []*v1.Role{
			{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: registerOpts.RemoteNamespace,
					Name:      "gloo-federation-controller",
				},
				Rules: glooFederationReadConfigPolicyRules,
			},
		},
	}

	return clusterRegisterOpts.RegisterCluster(ctx)
}
