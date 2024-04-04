package deployer_test

import (
	"context"
	"fmt"

	envoy_config_bootstrap "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/upstreams/http/v3"
	"github.com/golang/protobuf/ptypes/wrappers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/pkg/version"
	"github.com/solo-io/gloo/projects/gateway2/controller/scheme"
	"github.com/solo-io/gloo/projects/gateway2/deployer"
	"github.com/solo-io/gloo/projects/gateway2/extensions"
	v1 "github.com/solo-io/gloo/projects/gateway2/pkg/api/external/kubernetes/api/core/v1"
	gw2_v1alpha1 "github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1"
	"github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1/kube"
	"github.com/solo-io/gloo/projects/gateway2/wellknown"
	"github.com/solo-io/gloo/projects/gloo/pkg/bootstrap"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/utils/protoutils"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	api "sigs.k8s.io/gateway-api/apis/v1"
)

// testBootstrap implements resources.Resource in order to use protoutils.UnmarshalYAML
// this is hacky but it seems more stable/concise than map-casting all the way down
// to the field we need.
type testBootstrap struct {
	envoy_config_bootstrap.Bootstrap
}

func (t *testBootstrap) SetMetadata(meta *core.Metadata) {
	return
}

func (t *testBootstrap) Equal(_ any) bool {
	return false
}
func (t *testBootstrap) GetMetadata() *core.Metadata {
	return nil
}

type clientObjects []client.Object

func (objs *clientObjects) findDeployment(namespace, name string) *appsv1.Deployment {
	for _, obj := range *objs {
		if dep, ok := obj.(*appsv1.Deployment); ok {
			if dep.Name == name && dep.Namespace == namespace {
				return dep
			}
		}
	}
	return nil
}
func (objs *clientObjects) findServiceAccount(namespace, name string) *corev1.ServiceAccount {
	for _, obj := range *objs {
		if sa, ok := obj.(*corev1.ServiceAccount); ok {
			if sa.Name == name && sa.Namespace == namespace {
				return sa
			}
		}
	}
	return nil
}
func (objs *clientObjects) findService(namespace, name string) *corev1.Service {
	for _, obj := range *objs {
		if svc, ok := obj.(*corev1.Service); ok {
			if svc.Name == name && svc.Namespace == namespace {
				return svc
			}
		}
	}
	return nil
}
func (objs *clientObjects) findConfigMap(namespace, name string) *corev1.ConfigMap {
	for _, obj := range *objs {
		if cm, ok := obj.(*corev1.ConfigMap); ok {
			if cm.Name == name && cm.Namespace == namespace {
				return cm
			}
		}
	}
	return nil
}

func (objs *clientObjects) getEnvoyConfig(namespace, name string) *testBootstrap {
	cm := objs.findConfigMap(namespace, name).Data
	var bootstrapCfg testBootstrap
	err := protoutils.UnmarshalYAML([]byte(cm["envoy.yaml"]), &bootstrapCfg)
	ExpectWithOffset(1, err).ToNot(HaveOccurred())
	return &bootstrapCfg
}

var _ = Describe("Deployer", func() {
	const (
		defaultNamespace = "default"
	)
	var (
		d *deployer.Deployer

		gwc *api.GatewayClass
	)
	BeforeEach(func() {
		var err error

		gwc = &api.GatewayClass{
			ObjectMeta: metav1.ObjectMeta{
				Name: wellknown.GatewayClassName,
			},
			Spec: api.GatewayClassSpec{
				ControllerName: wellknown.GatewayControllerName,
			},
		}
		d, err = deployer.NewDeployer(newFakeClientWithObjs(gwc), &deployer.Inputs{
			ControllerName: wellknown.GatewayControllerName,
			Dev:            false,
			ControlPlane: bootstrap.ControlPlane{
				Kube: bootstrap.KubernetesControlPlaneConfig{XdsHost: "something.cluster.local", XdsPort: 1234},
			},
		})
		Expect(err).NotTo(HaveOccurred())
	})

	It("should get gvks", func() {
		gvks, err := d.GetGvksToWatch(context.Background())
		Expect(err).NotTo(HaveOccurred())
		Expect(gvks).NotTo(BeEmpty())
	})

	It("support segmenting by release", func() {
		mgr, err := ctrl.NewManager(&rest.Config{}, ctrl.Options{})
		Expect(err).NotTo(HaveOccurred())
		k8sGatewayExt, err := extensions.NewK8sGatewayExtensions(mgr)
		Expect(err).NotTo(HaveOccurred())

		d1, err := deployer.NewDeployer(newFakeClientWithObjs(gwc), &deployer.Inputs{
			ControllerName: wellknown.GatewayControllerName,
			Dev:            false,
			ControlPlane: bootstrap.ControlPlane{
				Kube: bootstrap.KubernetesControlPlaneConfig{XdsHost: "something.cluster.local", XdsPort: 1234},
			},
			Extensions: k8sGatewayExt,
		})
		Expect(err).NotTo(HaveOccurred())

		d2, err := deployer.NewDeployer(newFakeClientWithObjs(gwc), &deployer.Inputs{
			ControllerName: wellknown.GatewayControllerName,
			Dev:            false,
			ControlPlane: bootstrap.ControlPlane{
				Kube: bootstrap.KubernetesControlPlaneConfig{XdsHost: "something.cluster.local", XdsPort: 1234},
			},
			Extensions: k8sGatewayExt,
		})
		Expect(err).NotTo(HaveOccurred())

		gw1 := &api.Gateway{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo",
				Namespace: "default",
				UID:       "1235",
			},
			TypeMeta: metav1.TypeMeta{
				Kind:       "Gateway",
				APIVersion: "gateway.solo.io/v1beta1",
			},
			Spec: api.GatewaySpec{
				GatewayClassName: wellknown.GatewayClassName,
			},
		}

		gw2 := &api.Gateway{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "bar",
				Namespace: "default",
				UID:       "1235",
			},
			TypeMeta: metav1.TypeMeta{
				Kind:       "Gateway",
				APIVersion: "gateway.solo.io/v1beta1",
			},
			Spec: api.GatewaySpec{
				GatewayClassName: wellknown.GatewayClassName,
			},
		}

		proxyName := func(name string) string {
			return fmt.Sprintf("gloo-proxy-%s", name)
		}
		var objs1, objs2 clientObjects
		objs1, err = d1.GetObjsToDeploy(context.Background(), gw1)
		Expect(err).NotTo(HaveOccurred())
		Expect(objs1).NotTo(BeEmpty())
		Expect(objs1.findDeployment(defaultNamespace, proxyName(gw1.Name))).ToNot(BeNil())
		Expect(objs1.findService(defaultNamespace, proxyName(gw1.Name))).ToNot(BeNil())
		Expect(objs1.findConfigMap(defaultNamespace, proxyName(gw1.Name))).ToNot(BeNil())
		// Expect(objs1.findServiceAccount("default")).ToNot(BeNil())
		objs2, err = d2.GetObjsToDeploy(context.Background(), gw2)
		Expect(err).NotTo(HaveOccurred())
		Expect(objs2).NotTo(BeEmpty())
		Expect(objs2.findDeployment(defaultNamespace, proxyName(gw2.Name))).ToNot(BeNil())
		Expect(objs2.findService(defaultNamespace, proxyName(gw2.Name))).ToNot(BeNil())
		Expect(objs2.findConfigMap(defaultNamespace, proxyName(gw2.Name))).ToNot(BeNil())
		// Expect(objs2.findServiceAccount("default")).ToNot(BeNil())

		for _, obj := range objs1 {
			Expect(obj.GetName()).To(Equal("gloo-proxy-foo"))
		}
		for _, obj := range objs2 {
			Expect(obj.GetName()).To(Equal("gloo-proxy-bar"))
		}

	})

	Context("Single gwc and gw", func() {
		type input struct {
			dInputs        *deployer.Inputs
			gw             *api.Gateway
			gwp            *gw2_v1alpha1.GatewayParameters
			arbitrarySetup func()
		}

		type expectedOutput struct {
			err            error
			validationFunc func(objs clientObjects, inp *input) error
		}

		var (
			defaultGwpName        = "default-gateway-params"
			defaultDeployerInputs = func() *deployer.Inputs {
				mgr, err := ctrl.NewManager(&rest.Config{}, ctrl.Options{})
				Expect(err).NotTo(HaveOccurred())
				k8sGatewayExt, err := extensions.NewK8sGatewayExtensions(mgr)
				Expect(err).NotTo(HaveOccurred())
				return &deployer.Inputs{
					ControllerName: wellknown.GatewayControllerName,
					Dev:            false,
					ControlPlane: bootstrap.ControlPlane{
						Kube: bootstrap.KubernetesControlPlaneConfig{XdsHost: "something.cluster.local", XdsPort: 1234},
					},
					Extensions: k8sGatewayExt,
				}
			}
			defaultDeployerInputsWithSds = func() *deployer.Inputs {
				inp := defaultDeployerInputs()
				inp.IstioValues.SDSEnabled = true
				return inp
			}
			defaultGateway = func() *api.Gateway {
				return &api.Gateway{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "foo",
						Namespace: defaultNamespace,
						UID:       "1235",
					},
					TypeMeta: metav1.TypeMeta{
						Kind:       "Gateway",
						APIVersion: "gateway.solo.io/v1beta1",
					},
					Spec: api.GatewaySpec{
						GatewayClassName: wellknown.GatewayClassName,
						Listeners: []api.Listener{
							{
								Name: "listener-1",
								Port: 80,
							},
						},
					},
				}
			}
			defaultGatewayParams = func() *gw2_v1alpha1.GatewayParameters {
				return &gw2_v1alpha1.GatewayParameters{
					TypeMeta: metav1.TypeMeta{
						Kind: gw2_v1alpha1.GatewayParametersGVK.Kind,
						// The parsing expects GROUP/VERSION format in this field
						APIVersion: fmt.Sprintf("%s/%s", gw2_v1alpha1.GatewayParametersGVK.Group, gw2_v1alpha1.GatewayParametersGVK.Version),
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      defaultGwpName,
						Namespace: defaultNamespace,
						UID:       "1236",
					},
					Spec: gw2_v1alpha1.GatewayParametersSpec{
						ProxyConfig: &gw2_v1alpha1.ProxyConfig{
							EnvironmentType: &gw2_v1alpha1.ProxyConfig_Kube{
								Kube: &gw2_v1alpha1.KubernetesProxyConfig{
									WorkloadType: &gw2_v1alpha1.KubernetesProxyConfig_Deployment{
										Deployment: &gw2_v1alpha1.ProxyDeployment{
											Replicas: &wrappers.UInt32Value{Value: 3},
										},
									},
									EnvoyContainer: &gw2_v1alpha1.EnvoyContainer{
										Bootstrap: &gw2_v1alpha1.EnvoyBootstrap{
											LogLevel:          "debug",
											ComponentLogLevel: "router:info,listener:warn",
										},
										Image: &kube.Image{
											Registry:   "foo",
											Repository: "bar",
											Tag:        "bat",
											PullPolicy: kube.Image_Always,
										},
									},
									PodTemplate: &kube.Pod{
										ExtraAnnotations: map[string]string{
											"foo": "bar",
										},
										SecurityContext: &v1.PodSecurityContext{
											RunAsUser:  func() *int64 { var i int64 = 1; return &i }(),
											RunAsGroup: func() *int64 { var i int64 = 2; return &i }(),
										},
									},
									Service: &kube.Service{
										Type:      kube.Service_ClusterIP,
										ClusterIP: "99.99.99.99",
										ExtraAnnotations: map[string]string{
											"foo": "bar",
										},
									},
								},
							},
						},
					},
				}
			}
			defaultGatewayWithGatewayParams = func(gwpName string) *api.Gateway {
				gw := defaultGateway()
				gw.Annotations = map[string]string{
					wellknown.GatewayParametersAnnotationName: gwpName,
				}

				return gw
			}
			defaultInput = func() *input {
				return &input{
					dInputs: defaultDeployerInputs(),
					gw:      defaultGateway(),
					gwp:     defaultGatewayParams(),
				}
			}
			defaultDeploymentName     = fmt.Sprintf("gloo-proxy-%s", defaultGateway().Name)
			defaultConfigMapName      = defaultDeploymentName
			defaultServiceName        = defaultDeploymentName
			defaultServiceAccountName = defaultDeploymentName
		)
		DescribeTable("create and validate objs", func(inp *input, expected *expectedOutput) {
			// run break-glass setup
			if inp.arbitrarySetup != nil {
				inp.arbitrarySetup()
			}

			// Catch a nil gwp so the fake client doesn't choke on the nil obj
			gwp := inp.gwp
			if gwp == nil {
				gwp = &gw2_v1alpha1.GatewayParameters{}
			}

			d, err := deployer.NewDeployer(newFakeClientWithObjs(gwc, gwp), inp.dInputs)
			// We don't have any interesting error cases in the NewDeployer to test for but if we get
			// some then we will need to handle those outside the table or be more clever about expected
			// errors
			Expect(err).NotTo(HaveOccurred())

			objs, err := d.GetObjsToDeploy(context.Background(), inp.gw)
			if expected.err != nil {
				Expect(err).To(MatchError(expected.err))
				// return here since we matched our expected error
				return
			} else {
				Expect(err).NotTo(HaveOccurred())
			}

			// handle custom test validation func
			Expect(expected.validationFunc(objs, inp)).NotTo(HaveOccurred())
		},
			Entry("No GatewayParameters", &input{
				dInputs: defaultDeployerInputs(),
				gw:      defaultGateway(),
			}, &expectedOutput{
				validationFunc: func(objs clientObjects, inp *input) error {
					Expect(objs).NotTo(BeEmpty())
					// Check we have Deployment, ConfigMap, ServiceAccount, Service
					Expect(objs).To(HaveLen(4))
					cm := objs.findConfigMap(defaultNamespace, defaultConfigMapName)
					Expect(cm).ToNot(BeNil())

					dep := objs.findDeployment(defaultNamespace, defaultDeploymentName)
					Expect(dep).ToNot(BeNil())
					Expect(dep.Spec.Replicas).ToNot(BeNil())
					Expect(*dep.Spec.Replicas).To(Equal(int32(1)))
					gatewayContainer := dep.Spec.Template.Spec.Containers[0]
					Expect(gatewayContainer.Name).To(Equal("gloo-gateway"))
					Expect(gatewayContainer.Ports).To(ContainElement(HaveField("Name", "readiness")))
					for _, port := range gatewayContainer.Ports {
						if port.Name == "readiness" {
							// default from values.yaml
							Expect(port.ContainerPort).To(Equal(int32(8082)))
						}
					}
					Expect(gatewayContainer.Image).To(ContainSubstring("quay.io/solo-io/gloo-envoy-wrapper"))

					svc := objs.findService(defaultNamespace, defaultServiceName)
					Expect(svc).ToNot(BeNil())

					sa := objs.findServiceAccount(defaultNamespace, defaultServiceAccountName)
					Expect(sa).ToNot(BeNil())

					return nil
				},
			}),
			Entry("GatewayParameters overrides", &input{
				dInputs: defaultDeployerInputs(),
				gw:      defaultGatewayWithGatewayParams(defaultGwpName),
				gwp:     defaultGatewayParams(),
			}, &expectedOutput{
				validationFunc: func(objs clientObjects, inp *input) error {
					Expect(objs).NotTo(BeEmpty())
					// Check we have Deployment, ConfigMap, ServiceAccount, Service
					Expect(objs).To(HaveLen(4))
					dep := objs.findDeployment(defaultNamespace, defaultDeploymentName)
					Expect(dep).ToNot(BeNil())
					Expect(dep.Spec.Replicas).ToNot(BeNil())
					Expect(*dep.Spec.Replicas).To(Equal(int32(inp.gwp.Spec.ProxyConfig.GetKube().GetDeployment().Replicas.GetValue())))
					Expect(dep.Spec.Template.Spec.Containers[0].Image).To(Equal("foo/bar:bat"))
					Expect(string(dep.Spec.Template.Spec.Containers[0].ImagePullPolicy)).To(Equal(inp.gwp.Spec.GetProxyConfig().GetKube().GetEnvoyContainer().GetImage().GetPullPolicy().String()))
					Expect(dep.Spec.Template.Annotations["foo"]).To(Equal("bar"))
					Expect(*dep.Spec.Template.Spec.SecurityContext.RunAsUser).To(Equal(inp.gwp.Spec.GetProxyConfig().GetKube().GetPodTemplate().GetSecurityContext().GetRunAsUser()))
					Expect(*dep.Spec.Template.Spec.SecurityContext.RunAsGroup).To(Equal(inp.gwp.Spec.GetProxyConfig().GetKube().GetPodTemplate().GetSecurityContext().GetRunAsGroup()))

					svc := objs.findService(defaultNamespace, defaultServiceName)
					Expect(svc).ToNot(BeNil())
					Expect(svc.GetAnnotations()).ToNot(BeNil())
					Expect(svc.Annotations["foo"]).To(Equal("bar"))
					Expect(string(svc.Spec.Type)).To(Equal(inp.gwp.Spec.GetProxyConfig().GetKube().GetService().GetType().String()))
					Expect(svc.Spec.ClusterIP).To(Equal(inp.gwp.Spec.GetProxyConfig().GetKube().GetService().GetClusterIP()))

					sa := objs.findServiceAccount(defaultNamespace, defaultServiceAccountName)
					Expect(sa).ToNot(BeNil())

					cm := objs.findConfigMap(defaultNamespace, defaultConfigMapName)
					Expect(cm).ToNot(BeNil())
					Expect(objs.findDeployment(defaultNamespace, defaultDeploymentName).Spec.Template.Spec.Containers[0].Args).To(ContainElements(
						"--log-level",
						inp.gwp.Spec.GetProxyConfig().GetKube().GetEnvoyContainer().GetBootstrap().GetLogLevel(),
						"--component-log-level",
						inp.gwp.Spec.GetProxyConfig().GetKube().GetEnvoyContainer().GetBootstrap().GetComponentLogLevel(),
					))
					return nil
				},
			}),
			Entry("correct deployment with sds enabled", &input{
				dInputs: defaultDeployerInputsWithSds(),
				gw:      defaultGateway(),
				gwp:     defaultGatewayParams(),
			}, &expectedOutput{
				validationFunc: func(objs clientObjects, inp *input) error {
					containers := objs.findDeployment(defaultNamespace, defaultDeploymentName).Spec.Template.Spec.Containers
					Expect(containers).To(HaveLen(3))
					foundGw, foundSds, foundIstioProxy := false, false, false
					for _, container := range containers {
						switch container.Name {
						case "sds":
							foundSds = true
						case "istio-proxy":
							foundIstioProxy = true
						case "gloo-gateway":
							foundGw = true
						default:
							Fail("unknown container name " + container.Name)
						}

					}
					Expect(foundGw).To(BeTrue())
					Expect(foundSds).To(BeTrue())
					Expect(foundIstioProxy).To(BeTrue())

					bootstrapCfg := objs.getEnvoyConfig(defaultNamespace, defaultConfigMapName)
					clusters := bootstrapCfg.GetStaticResources().GetClusters()
					Expect(clusters).ToNot(BeNil())
					Expect(clusters).To(ContainElement(HaveField("Name", "gateway_proxy_sds")))

					return nil
				},
			}),
			Entry("no listeners on gateway", &input{
				dInputs: defaultDeployerInputs(),
				gw: &api.Gateway{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "foo",
						Namespace: defaultNamespace,
						UID:       "1235",
					},
					TypeMeta: metav1.TypeMeta{
						Kind:       "Gateway",
						APIVersion: "gateway.solo.io/v1beta1",
					},
					Spec: api.GatewaySpec{
						GatewayClassName: "gloo-gateway",
					},
				},
				gwp: defaultGatewayParams(),
			}, &expectedOutput{
				validationFunc: func(objs clientObjects, inp *input) error {
					Expect(objs).NotTo(BeEmpty())
					return nil
				},
			}),
			Entry("port offset", defaultInput(), &expectedOutput{
				validationFunc: func(objs clientObjects, inp *input) error {
					svc := objs.findService(defaultNamespace, defaultServiceName)
					Expect(svc).NotTo(BeNil())

					port := svc.Spec.Ports[0]
					Expect(port.Port).To(Equal(int32(80)))
					Expect(port.TargetPort.IntVal).To(Equal(int32(8080)))
					return nil
				},
			}),
			Entry("duplicate ports", &input{
				dInputs: defaultDeployerInputs(),
				gw: &api.Gateway{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "foo",
						Namespace: defaultNamespace,
						UID:       "1235",
					},
					TypeMeta: metav1.TypeMeta{
						Kind:       "Gateway",
						APIVersion: "gateway.solo.io/v1beta1",
					},
					Spec: api.GatewaySpec{
						GatewayClassName: "gloo-gateway",
						Listeners: []api.Listener{
							{
								Name: "listener-1",
								Port: 80,
							},
							{
								Name: "listener-2",
								Port: 80,
							},
						},
					},
				},
				gwp: defaultGatewayParams(),
			}, &expectedOutput{
				validationFunc: func(objs clientObjects, inp *input) error {
					svc := objs.findService(defaultNamespace, defaultServiceName)
					Expect(svc).NotTo(BeNil())

					Expect(svc.Spec.Ports).To(HaveLen(1))
					port := svc.Spec.Ports[0]
					Expect(port.Port).To(Equal(int32(80)))
					Expect(port.TargetPort.IntVal).To(Equal(int32(8080)))
					return nil
				},
			}),
			Entry("propagates version.Version to deployment", &input{
				dInputs:        defaultDeployerInputs(),
				gw:             defaultGateway(),
				arbitrarySetup: func() { version.Version = "testversion" },
				gwp:            defaultGatewayParams(),
			}, &expectedOutput{
				validationFunc: func(objs clientObjects, inp *input) error {
					dep := objs.findDeployment(defaultNamespace, defaultDeploymentName)
					Expect(dep).NotTo(BeNil())
					Expect(dep.Spec.Template.Spec.Containers).NotTo(BeEmpty())
					for _, c := range dep.Spec.Template.Spec.Containers {
						Expect(c.Image).To(HaveSuffix(":testversion"))
					}
					return nil
				},
			}),
			Entry("object owner refs are set", defaultInput(), &expectedOutput{
				validationFunc: func(objs clientObjects, inp *input) error {
					Expect(objs).NotTo(BeEmpty())

					gw := defaultGateway()

					for _, obj := range objs {
						ownerRefs := obj.GetOwnerReferences()
						Expect(ownerRefs).To(HaveLen(1))
						Expect(ownerRefs[0].Name).To(Equal(gw.Name))
						Expect(ownerRefs[0].UID).To(Equal(gw.UID))
						Expect(ownerRefs[0].Kind).To(Equal(gw.Kind))
						Expect(ownerRefs[0].APIVersion).To(Equal(gw.APIVersion))
						Expect(*ownerRefs[0].Controller).To(BeTrue())
					}
					return nil
				},
			}),
			Entry("envoy yaml is valid", defaultInput(), &expectedOutput{
				validationFunc: func(objs clientObjects, inp *input) error {
					gw := defaultGateway()
					Expect(objs).NotTo(BeEmpty())

					cm := objs.findConfigMap(defaultNamespace, defaultConfigMapName)
					Expect(cm).NotTo(BeNil())

					envoyYaml := cm.Data["envoy.yaml"]
					Expect(envoyYaml).NotTo(BeEmpty())

					// make sure it's valid yaml
					var envoyConfig map[string]any
					err := yaml.Unmarshal([]byte(envoyYaml), &envoyConfig)
					Expect(err).NotTo(HaveOccurred(), "envoy config is not valid yaml: %s", envoyYaml)

					// make sure the envoy node metadata looks right
					node := envoyConfig["node"].(map[string]any)
					Expect(node).To(HaveKeyWithValue("metadata", map[string]any{
						"gateway": map[string]any{
							"name":      gw.Name,
							"namespace": gw.Namespace,
						},
					}))
					return nil
				},
			}),
			Entry("failed to get GatewayParameters", &input{
				dInputs: defaultDeployerInputs(),
				gw:      defaultGatewayWithGatewayParams("bad-gwp"),
			}, &expectedOutput{
				err: deployer.GetGatewayParametersError,
			}),
		)
	})
})

// initialize a fake controller-runtime client with the given list of objects
func newFakeClientWithObjs(objs ...client.Object) client.Client {
	s := scheme.NewScheme()
	return fake.NewClientBuilder().WithScheme(s).WithObjects(objs...).Build()
}
