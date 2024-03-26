package deployer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo/pkg/version"
	"github.com/solo-io/gloo/projects/gateway2/helm"
	"github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1"
	v1alpha1kube "github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1/kube"
	"github.com/solo-io/gloo/projects/gloo/constants"
	"github.com/solo-io/gloo/projects/gloo/pkg/bootstrap"
	"golang.org/x/exp/slices"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/storage"
	"helm.sh/helm/v3/pkg/storage/driver"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/yaml"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	api "sigs.k8s.io/gateway-api/apis/v1"
)

var (
	NoGatewayClassError = func(gw *api.Gateway) error {
		return eris.Errorf("gateway %s.%s does not contain a gatewayClassName", gw.Namespace, gw.Name)
	}
	GetGatewayClassError = func(err error, gw *api.Gateway, gatewayClassName string) error {
		return eris.Wrapf(err, "could not retrieve gatewayclass %s for gateway %s.%s", gatewayClassName, gw.Namespace, gw.Name)
	}
	UnsupportedParametersRefKind = func(gatewayClassName string, parametersRef *api.ParametersReference) error {
		return eris.Errorf("parametersRef for gatewayclass %s points to an unsupported kind: %v", gatewayClassName, parametersRef)
	}
	GetDataPlaneConfigError = func(err error, gatewayClassName string, dpcNamespace string, dpcName string) error {
		return eris.Wrapf(err, "could not retrieve dataplaneconfig (%s.%s) for gatewayclass %s", dpcNamespace, dpcName, gatewayClassName)
	}
)

type gatewayPort struct {
	Port       uint16 `json:"port"`
	Protocol   string `json:"protocol"`
	Name       string `json:"name"`
	TargetPort uint16 `json:"targetPort"`
}

// A Deployer is responsible for deploying proxies
type Deployer struct {
	chart *chart.Chart
	cli   client.Client

	inputs *Inputs
}

// Inputs is the set of options used to configure the gateway deployer deployment
type Inputs struct {
	ControllerName string
	Dev            bool
	IstioValues    bootstrap.IstioValues
}

// NewDeployer creates a new gateway deployer
func NewDeployer(cli client.Client, inputs *Inputs) (*Deployer, error) {
	helmChart, err := loadFs(helm.GlooGatewayHelmChart)
	if err != nil {
		return nil, err
	}
	// simulate what `helm package` in the Makefile does
	if version.Version != version.UndefinedVersion {
		helmChart.Metadata.AppVersion = version.Version
		helmChart.Metadata.Version = version.Version
	}

	return &Deployer{
		cli:    cli,
		chart:  helmChart,
		inputs: inputs,
	}, nil
}

// GetGvksToWatch returns the list of GVKs that the deployer will watch for
func (d *Deployer) GetGvksToWatch(ctx context.Context) ([]schema.GroupVersionKind, error) {
	fakeGw := &api.Gateway{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "default",
			Namespace: "default",
		},
	}

	// these are the minimal values that render the Deployment, Service, ServiceAccount, and ConfigMap
	vals := map[string]any{
		"gateway": map[string]any{
			"serviceAccount": map[string]any{
				"create": true,
			},
		},
	}

	objs, err := d.renderChartToObjects(ctx, fakeGw, vals)
	if err != nil {
		return nil, err
	}
	var ret []schema.GroupVersionKind
	for _, obj := range objs {
		gvk := obj.GetObjectKind().GroupVersionKind()
		if !slices.Contains(ret, gvk) {
			ret = append(ret, gvk)
		}
	}

	log.FromContext(ctx).V(1).Info("watching GVKs", "GVKs", ret)
	return ret, nil
}

func jsonConvert2(in *v1alpha1.DataPlaneConfig, out interface{}) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, out)
}

func jsonConvert3(in *v1alpha1kube.Autoscaling, out interface{}) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, out)
}

func (d *Deployer) renderChartToObjects(ctx context.Context, gw *api.Gateway, vals map[string]any) ([]client.Object, error) {
	logger := log.FromContext(ctx)
	logger.Info("rendering helm chart", "vals", vals)

	objs, err := d.Render(ctx, gw.Name, gw.Namespace, vals)
	if err != nil {
		return nil, err
	}

	for _, obj := range objs {
		obj.SetNamespace(gw.Namespace)
	}

	return objs, nil
}

// Gets the DataPlaneConfig object (if any) associated with a given Gateway.
func (d *Deployer) getDataPlaneConfigForGateway(ctx context.Context, gw *api.Gateway) (*v1alpha1.DataPlaneConfig, error) {
	logger := log.FromContext(ctx)

	// Get the GatewayClass for the Gateway
	gwClassName := gw.Spec.GatewayClassName
	if gwClassName == "" {
		// this shouldn't happen as the gatewayClassName field is required, but throw an error in this case
		return nil, NoGatewayClassError(gw)
	}

	gwc := &api.GatewayClass{}
	err := d.cli.Get(ctx, client.ObjectKey{Name: string(gwClassName)}, gwc)
	if err != nil {
		return nil, GetGatewayClassError(err, gw, string(gwClassName))
	}

	// Get the DataPlaneConfig from the GatewayClass
	paramsRef := gwc.Spec.ParametersRef
	if paramsRef == nil {
		// there is no custom data plane config (just use default values)
		logger.V(1).Info("no parametersRef found for GatewayClass", "GatewayClass", gwClassName)
		return nil, nil
	}

	if string(paramsRef.Group) != v1alpha1.DataPlaneConfigGVK.Group ||
		string(paramsRef.Kind) != v1alpha1.DataPlaneConfigGVK.Kind {
		return nil, UnsupportedParametersRefKind(string(gwClassName), paramsRef)
	}

	dpc := &v1alpha1.DataPlaneConfig{}
	err = d.cli.Get(ctx, client.ObjectKey{Namespace: string(*paramsRef.Namespace), Name: paramsRef.Name}, dpc)
	if err != nil {
		return nil, GetDataPlaneConfigError(err, string(gwClassName), string(*paramsRef.Namespace), paramsRef.Name)
	}

	return dpc, nil
}

func (d *Deployer) getValues(ctx context.Context, gw *api.Gateway) (map[string]any, error) {
	dpc, err := d.getDataPlaneConfigForGateway(ctx, gw)
	if err != nil {
		return nil, err
	}

	fmt.Printf("xxxxx got DataPlaneConfig: %v\n", dpc)
	portsVals, err := GetPortsValues(gw)
	if err != nil {
		return nil, err
	}

	xdsPort, err := GetDefaultXdsPort(ctx, d.cli)
	if err != nil {
		return nil, err
	}

	kubeProxyConfig := dpc.Spec.GetProxyConfig().GetKube()
	deployConfig := kubeProxyConfig.GetDeployment()
	podConfig := deployConfig.GetPodTemplate()
	envoyContainerConfig := deployConfig.GetEnvoyContainer()
	svcConfig := kubeProxyConfig.GetService()

	var replicas any
	autoscalingVals := GetAutoscalingValues(kubeProxyConfig.GetAutoscaling())
	if autoscalingVals == nil && deployConfig.GetReplicas() != nil {
		replicas = deployConfig.GetReplicas().GetValue()
	}

	vals := map[string]any{
		"gateway": map[string]any{
			"name":        gw.Name,
			"gatewayName": gw.Name,

			// deployment/service values
			"replicaCount": replicas,
			"autoscaling":  autoscalingVals,
			"ports":        portsVals,
			"service": map[string]any{
				// convert the service type enum to its string representation;
				// if type is not set, it will default to 0 ("ClusterIP")
				"type":             v1alpha1kube.Service_ServiceType_name[int32(svcConfig.GetType())],
				"clusterIP":        svcConfig.GetClusterIP(),
				"extraAnnotations": svcConfig.GetExtraAnnotations(),
				"extraLabels":      svcConfig.GetExtraLabels(),
			},

			// pod template values
			"extraPodAnnotations": podConfig.GetExtraAnnotations(),
			"extraPodLabels":      podConfig.GetExtraLabels(),
			//"imagePullSecrets": podConfig.GetImagePullSecrets(),
			//"podSecurityContext": podConfig.GetSecurityContext(),
			//"nodeSelector":podConfig.GetNodeSelector(),
			//"affinity":podConfig.GetAffinity(),
			//"tolerations":podConfig.GetTolerations(),

			// envoy container values
			"logLevel":          envoyContainerConfig.GetLogLevel(),
			"componentLogLevel": envoyContainerConfig.GetComponentLogLevel(),
			"image":             getDeployerImageValues(), // envoyContainerConfig.GetImage()
			//"resources": envoyContainerConfig.GetResources(),
			//"securityContext": envoyContainerConfig.GetSecurityContext(),

			// istio values
			"istioSDS": map[string]any{
				"enabled": d.inputs.IstioValues.SDSEnabled,
			},

			// xds values
			"xds": map[string]any{
				// The xds host/port MUST map to the Service definition for the Control Plane
				// This is the socket address that the Proxy will connect to on startup, to receive xds updates
				"host": GetDefaultXdsHost(),
				"port": xdsPort,
			},
		},
	}

	return vals, nil
}

func (d *Deployer) Render(ctx context.Context, name, ns string, vals map[string]any) ([]client.Object, error) {
	mem := driver.NewMemory()
	mem.SetNamespace(ns)
	cfg := &action.Configuration{
		Releases: storage.Init(mem),
	}
	client := action.NewInstall(cfg)
	client.Namespace = ns
	client.ReleaseName = name
	client.ClientOnly = true
	release, err := client.RunWithContext(ctx, d.chart, vals)
	if err != nil {
		return nil, fmt.Errorf("failed to render helm chart: %w", err)
	}

	fmt.Printf("xxxxxx release manifest: %v\n", release.Manifest)
	objs, err := ConvertYAMLToObjects(d.cli.Scheme(), []byte(release.Manifest))
	if err != nil {
		return nil, fmt.Errorf("failed to convert yaml to objects: %w", err)
	}
	return objs, nil
}

func (d *Deployer) GetObjsToDeploy(ctx context.Context, gw *api.Gateway) ([]client.Object, error) {
	logger := log.FromContext(ctx)

	vals, err := d.getValues(ctx, gw)
	if err != nil {
		return nil, fmt.Errorf("failed to get values to render objects: %w", err)
	}
	logger.V(1).Info("got deployer helm values", "values", vals)

	objs, err := d.renderChartToObjects(ctx, gw, vals)
	if err != nil {
		return nil, fmt.Errorf("failed to get objects to deploy: %w", err)
	}

	// Set owner ref
	trueVal := true
	for _, obj := range objs {
		obj.SetOwnerReferences([]metav1.OwnerReference{{
			Kind:       gw.Kind,
			APIVersion: gw.APIVersion,
			Controller: &trueVal,
			UID:        gw.UID,
			Name:       gw.Name,
		}})
	}

	return objs, nil
}

func (d *Deployer) DeployObjs(ctx context.Context, objs []client.Object) error {
	logger := log.FromContext(ctx)
	for _, obj := range objs {
		logger.V(1).Info("deploying object", "kind", obj.GetObjectKind(), "namespace", obj.GetNamespace(), "name", obj.GetName())
		if err := d.cli.Patch(ctx, obj, client.Apply, client.ForceOwnership, client.FieldOwner(d.inputs.ControllerName)); err != nil {
			return fmt.Errorf("failed to apply object %s %s: %w", obj.GetObjectKind().GroupVersionKind().String(), obj.GetName(), err)
		}
	}
	return nil
}

func loadFs(filesystem fs.FS) (*chart.Chart, error) {
	var bufferedFiles []*loader.BufferedFile
	entries, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}
	if len(entries) != 1 {
		return nil, fmt.Errorf("expected exactly one entry in the chart folder, got %v", entries)
	}

	root := entries[0].Name()
	err = fs.WalkDir(filesystem, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		data, readErr := fs.ReadFile(filesystem, path)
		if readErr != nil {
			return readErr
		}

		relativePath, relErr := filepath.Rel(root, path)
		if relErr != nil {
			return relErr
		}

		bufferedFile := &loader.BufferedFile{
			Name: relativePath,
			Data: data,
		}

		bufferedFiles = append(bufferedFiles, bufferedFile)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return loader.LoadFiles(bufferedFiles)
}

func ConvertYAMLToObjects(scheme *runtime.Scheme, yamlData []byte) ([]client.Object, error) {
	var objs []client.Object

	// Split the YAML manifest into separate documents
	decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(yamlData), 4096)
	for {
		var obj unstructured.Unstructured
		if err := decoder.Decode(&obj); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		// try to translate to real objects, so they are easier to query later
		gvk := obj.GetObjectKind().GroupVersionKind()
		if realObj, err := scheme.New(gvk); err == nil {
			if realObj, ok := realObj.(client.Object); ok {
				if err := runtime.DefaultUnstructuredConverter.FromUnstructured(obj.Object, realObj); err == nil {
					objs = append(objs, realObj)
					continue
				}
			}
		} else if len(obj.Object) == 0 {
			// This can happen with an "empty" document
			continue
		}

		objs = append(objs, &obj)
	}

	return objs, nil
}

func getDeployerImageValues() map[string]any {
	image := os.Getenv(constants.GlooGatewayDeployerImage)
	defaultImageValues := map[string]any{
		// If tag is not defined, we fall back to the default behavior, which is to use that Chart version
		"tag": "",
	}

	if image == "" {
		// If the env is not defined, return the default
		return defaultImageValues
	}

	imageParts := strings.Split(image, ":")
	if len(imageParts) != 2 {
		// If the user provided an invalid override, fallback to the default
		return defaultImageValues
	}
	return map[string]any{
		"repository": imageParts[0],
		"tag":        imageParts[1],
	}
}
