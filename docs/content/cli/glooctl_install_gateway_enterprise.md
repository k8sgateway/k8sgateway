---
title: "glooctl install gateway enterprise"
weight: 5
---
## glooctl install gateway enterprise

install the Gloo Enterprise Gateway on kubernetes

### Synopsis

requires kubectl to be installed

```
glooctl install gateway enterprise [flags]
```

### Options

```
  -h, --help                 help for enterprise
      --license-key string   License key to activate GlooE features
```

### Options inherited from parent commands

```
  -c, --config string        set the path to the glooctl config file (default "<home_directory>/.gloo/glooctl-config.yaml")
  -d, --dry-run              Dump the raw installation yaml instead of applying it to kubernetes
  -f, --file string          Install Gloo from this Helm chart archive file rather than from a release
  -i, --interactive          use interactive mode
      --kubeconfig string    kubeconfig to use, if not standard one
  -n, --namespace string     namespace to install gloo into (default "gloo-system")
  -u, --upgrade              Upgrade an existing v1 gateway installation to use v2 CRDs. Set this when upgrading from v0.17.x or earlier versions of gloo
      --values strings       List of files with value overrides for the Gloo Helm chart, (e.g. --values file1,file2 or --values file1 --values file2)
  -v, --verbose              If true, output from kubectl commands will print to stdout/stderr
      --with-admin-console   install gloo and a read-only version of its admin console
```

### SEE ALSO

* [glooctl install gateway](../glooctl_install_gateway)	 - install the Gloo Gateway on kubernetes

