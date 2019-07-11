---
title: "glooctl install gateway"
weight: 5
---
## glooctl install gateway

install the Gloo Gateway on kubernetes

### Synopsis

requires kubectl to be installed

```
glooctl install gateway [flags]
```

### Options

```
  -d, --dry-run            Dump the raw installation yaml instead of applying it to kubernetes
  -f, --file string        Install Gloo from this Helm chart archive file rather than from a release
  -h, --help               help for gateway
  -n, --namespace string   namespace to install gloo into (default "gloo-system")
```

### Options inherited from parent commands

```
  -i, --interactive   use interactive mode
  -v, --verbose       If true, output from kubectl commands will print to stdout/stderr
```

### SEE ALSO

* [glooctl install](../glooctl_install)	 - install gloo on different platforms

