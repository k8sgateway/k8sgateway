# Envoyinit

The running instance of Envoy. In Gloo Edge, this is commonly referred to as the `gateway-proxy` component.

## Background

The [Envoy Proxy](https://www.envoyproxy.io/) is a cloud-native, high-performance, service proxy. 

### Source Code

The Gloo Edge service proxies provide all the functionality of the [open source Envoy Proxy](https://github.com/solo-io/envoy-gloo), in addition to some custom extensions. The source code for these proxies is maintained at [envoy-gloo](https://github.com/solo-io/envoy-gloo)

### Versioning

In the [Makefile](../../Makefile), the `ENVOY_GLOO_IMAGE` value defines the version of `envoy-gloo` that Gloo Edge depends on.

Envoy publishes new minor releases [each quarter](https://www.envoyproxy.io/docs/envoy/latest/version_history/version_history#). Gloo attempts to follow this cadence, and increment our minor version of `envoy-gloo` as well.

## Build

*All make targets are currently defined in the [Makefile](../../Makefile) at the root of the repository.*

The `VERSION` env variable determines the name of the tag for the image.

You may either inject the version yourself:
```bash
VERSION=<version name> make gloo-envoy-wrapper-docker
```

Or rely on the auto-generated version:
```bash
make gloo-envoy-wrapper-docker
```

## Release

During a Gloo Edge release, the `gloo-envoy-wrapper` image is published to the [Google Cloud Registry](https://console.cloud.google.com/gcr/images/gloo-edge/GLOBAL) and the [Quay repository](https://quay.io/repository/solo-io/gloo-envoy-wrapper).

## Configuration

Envoy [configuration](https://www.envoyproxy.io/docs/envoy/latest/configuration/configuration) can be provided either statically or dynamically. In Gloo Edge, we leverage both of these patterns: at initialization time, the proxy is provided with basic boostrap configuration, and then all future updates are provided dynamically by the control-plane.

### Static

Envoy is configured with [Bootstrap configuration](https://www.envoyproxy.io/docs/envoy/latest/configuration/overview/bootstrap). This provides static server configuration and configures Envoy to access dynamic configuration from the Gloo control plane.

In Gloo Edge, Envoy configuration is processed in the following order:

1. The bootstrap configuration is defined in a [ConfigMap](https://github.com/solo-io/gloo/blob/master/install/helm/gloo/templates/9-gateway-proxy-configmap.yaml)
2. The ConfigMap is mounted as a volume on the Pod.
3. At [initialization](./cmd/main.go), the container reads the configuration, and transforms it using the [Kubernetes Downward API](https://kubernetes.io/docs/tasks/inject-data-application/downward-api-volume-expose-pod-information/#the-downward-api)
4. The transformed configuration is provided to the Envoy executable

### Dynamic

Envoy receives dynamic configuration via the [xDS protocol](https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol#xds-protocol). The Gloo [xDS](https://github.com/solo-io/gloo/tree/master/projects/gloo/pkg/xds) package contains relevant code for serving dynamic configuration.

## Debug

It can be useful to run the Envoy proxy, without the control-plane, as a way of validating behavior. [hack/envoy.yaml](./hack/envoy.yaml) provides example bootstrap that can be used

`docker run \
    --rm \
    -ti \
    -p 8000:8000 \
    -p 19000:19000 \
    -v $(pwd)/hack/envoy.yaml:/etc/envoy/envoy.yaml:ro \
    -l trace \
    quay.io/solo-io/gloo-envoy-wrapper:1.11.11
`

After running this, you should see a lot of Envoy logs:
```text
[2022-05-19 23:21:59.465][8][info][main] [external/envoy/source/server/server.cc:381] initializing epoch 0 (base id=0, hot restart version=11.104)
[2022-05-19 23:21:59.465][8][info][main] [external/envoy/source/server/server.cc:383] statically linked extensions:
```

Envoy exposes an [administration interface](https://www.envoyproxy.io/docs/envoy/latest/operations/admin) which can be used to query and modify different aspects of the server. The address of this interface is defined in the [bootstrap API](https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/bootstrap/v3/bootstrap.proto#envoy-v3-api-msg-config-bootstrap-v3-admin), though it is commonly found at port `19000`. 
If the above command succeeded, you should be able to visit [port 19000 in your browser](http://localhost:19000/) to explore the admin interface.


## Useful Information

### Other resources

To discover further information about Envoy, there are a number of great sources:
 - [Hoot YouTube series](https://www.youtube.com/watch?v=KsO4pw4tEGA)
 - [Envoy docs](https://www.envoyproxy.io/docs.html)
 - [Envoy slack](https://envoyproxy.io/slack)

### Determine the underlying version of Envoy
`docker run --entrypoint=envoy gcr.io/gloo-edge/gloo-envoy-wrapper:1.11.11 --version`

```text
envoy  version: 1f606cca72a8cd5f712803a732d0dd97828bd860/1.21.1/Distribution/RELEASE/BoringSSL
```


