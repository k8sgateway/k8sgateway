---
title: Gloo Edge and Istio
menuTitle: Configure your Gloo Edge gateway to run an Istio sidecar 
weight: 1
---

You can configure your Gloo Edge gateway with an Istio sidecar to secure the connection between your gateway and the services in your Istio service mesh. The sidecar in your Gloo Edge gateway uses mutual TLS (mTLS) to prove its identity to the services in the mesh and vice versa.

## Before you begin

Complete the following tasks before configuring an Istio sidecar for your Gloo Edge gateway: 

1. Create or use an existing cluster that runs Kubernetes a version supported by both [your version of Edge]({{< versioned_link_path fromRoot="/reference/support" >}}) and the [version of Istio you intend to install](https://istio.io/latest/docs/releases/supported-releases/). 
2. [Install Istio in your cluster](https://istio.io/latest/docs/setup/getting-started/). Istio versions 1.13 through 1.18 are supported in Gloo Edge 1.15. See the [support matrix]({{< versioned_link_path fromRoot="/reference/support" >}}) for more details.
3. Set up a service mesh for your cluster. For example, you can use [Gloo Mesh Enterprise](https://docs.solo.io/gloo-mesh-enterprise/latest/getting_started/managed_kubernetes/) to configure a service mesh that is based on Envoy and Istio, and that can span across multiple service meshes and clusters. 
4. Install an application in your mesh, such as Bookinfo. 
   ```shell
   kubectl label namespace default istio-injection=enabled
   kubectl apply -f /path/to/istio/samples/bookinfo/platform/kube/bookinfo.yaml
   ```
   
5. Install [Helm version 3](https://helm.sh/docs/intro/install/) on your local machine.

## Configure the Gloo Edge gateway with an Istio sidecar

Install the Gloo Edge gateway and inject it with an Istio sidecar. 

1. Add the Gloo Edge Helm repo. 
   ```shell
   helm repo add gloo https://storage.googleapis.com/solo-public-helm
   ```
   
2. Update the repo. 
   ```shell
   helm repo update
   ```
      
3. Create a `value-overrides.yaml` file with the following content. To configure your gateway with an Istio sidecar, make sure to add the `istioIntegration` section and set the `enableIstioSidecarOnGateway` option to `true`. You can optionally add the `global.istioSDS.enabled` option to your overrides file to automatically renew the certificate that the sidecar uses before it expires. 
Be sure to specify valid image fields under `global.glooMtls.istioProxy.image` and `global.glooMtls.sds.image`. The default Istio version is 1.18.2.
   ```yaml
   global:
     istioIntegration:
       enableIstioSidecarOnGateway: true
       disableAutoInjection: true
     istioSDS:
       enabled: true
     glooMtls:
       istioProxy:
         image:
           registry: docker.io/istio
           repository: proxyv2
           tag: 1.18.2
       sds:
         image:
           repository: sds
           tag: 1.15.7
   gatewayProxies:
     gatewayProxy:
       podTemplate: 
         httpPort: 8080
         httpsPort: 8443
   ```
   
4. Install or upgrade Gloo Edge. 
   {{< tabs >}} 
   {{< tab name="Install Gloo Edge">}}

   Install Gloo Edge with the settings in the `value-overrides.yaml` file. This command creates the `gloo-system` namespace and installs the Gloo Edge components into it.
   ```shell
    helm install gloo gloo/gloo --namespace gloo-system --create-namespace -f value-overrides.yaml
   ```
   {{< /tab >}}
   {{< tab name="Upgrade Gloo Edge">}}
   Upgrade Gloo Edge with the settings in the `value-overrides.yaml` file.
   ```shell
   helm upgrade gloo gloo/gloo --namespace gloo-system -f value-overrides.yaml
   ```
   {{< /tab >}}
   {{< /tabs >}}   
5. [Verify your setup]({{< versioned_link_path fromRoot="/installation/gateway/kubernetes/#verify-your-installation" >}}). 
   
7. Get the pods for your gateway proxy deployment. You now see three containers in the gateway-proxy pod. 
   ```shell
   kubectl get pods -n gloo-system
   ```
    
   Example output: 
   ```
   NAME                             READY   STATUS      RESTARTS   AGE
   discovery-6dcc8ddc58-q4zv7       1/1     Running     0          39s
   gateway-certgen-xzr7t            0/1     Completed   0          43s
   gateway-proxy-7bc5c97449-n9498   3/3     Running     0          39s
   gloo-d8cfbf86b-v59j4             1/1     Running     0          39s
   gloo-resource-rollout-hhvf9      0/1     Completed   0          38s
   ```
    
9. Describe the `gateway-proxy` pod to verify that the `istio-proxy` and `sds` containers are running. 
   ```shell
   kubectl describe <gateway-pod-name> -n gloo-system
   ```

Congratulations! You successfully configured an Istio sidecar for your Gloo Edge gateway. 

## Verify the mTLS connection 

To verify that you can connect to your app via mutual TLS (mTLS), you can install the Bookinfo app in your cluster and set up an upstream and a virtual service to route incoming requests to that app. 

1. If you haven't already, install the Bookinfo app in your cluster as part of the Istio mesh. 
   ```shell
   kubectl label namespace default istio-injection=enabled
   kubectl apply -f /path/to/istio/samples/bookinfo/platform/kube/bookinfo.yaml
   ```
   
   Example output: 
   ```
   service/details created
   serviceaccount/bookinfo-details created
   deployment.apps/details-v1 created
   service/ratings created
   serviceaccount/bookinfo-ratings created
   deployment.apps/ratings-v1 created
   service/reviews created
   serviceaccount/bookinfo-reviews created
   deployment.apps/reviews-v1 created
   deployment.apps/reviews-v2 created
   deployment.apps/reviews-v3 created
   service/productpage created
   serviceaccount/bookinfo-productpage created
   deployment.apps/productpage-v1 created
   ```

2. Create a virtual service to set up the routing rules for your Bookinfo app. In the following example, you instruct the Gloo Edge gateway to route incoming requests on the `/productpage` path to be routed to the `productpage` service in your cluster. 
   ```yaml
   kubectl apply -f- <<EOF
   apiVersion: gateway.solo.io/v1
   kind: VirtualService
   metadata:
     name: vs
     namespace: gloo-system
   spec:
     virtualHost:
       domains:
       - 'www.example.com'
     routes:
     - matchers:
       - prefix: /productpage
       routeAction:
         single:
           upstream:
             name: default-productpage-9080
             namespace: gloo-system
   EOF
   ```
   
3. At this point we can send a request to the product page. 
   ```shell
   curl -viks -H "Host: www.example.com" "$(glooctl proxy url)/productpage" --output /dev/null 
   ```
   A 200 response indicates success, however if we stop here, traffic to the productpage Upstream will not be encrypted with mTLS.

4. In order to require all traffic in the Mesh uses mTLS, apply the following STRICT PeerAuthentication policy:
   ```yaml
   kubectl apply -f - <<EOF
   apiVersion: "security.istio.io/v1beta1"
   kind: "PeerAuthentication"
   metadata:
     name: "test"
     namespace: "istio-system"
   spec:
     mtls:
       mode: STRICT
   EOF
   ```
   Now if we make the same curl request we will get a 503 response, as our upstream is not configured for Istio mTLS.

5. Use `glooctl` to configure the upstream for Istio mTLS: 
   ```shell
   glooctl istio enable-mtls --upstream default-productpage-9080
   ```
   Now the request will once again succeed, with a 200 response, with traffic encrypted using mTLS.

{{% notice note %}} 
If you use Gloo Mesh Enterprise for your service mesh, you can configure your Gloo Edge upstream resource to point to the Gloo Mesh `ingress-gateway`. For a request to reach the Bookinfo app in remote workload clusters, your virtual service must be configured to route traffic to the Gloo Mesh `east-west` gateway. 
{{% /notice %}}
