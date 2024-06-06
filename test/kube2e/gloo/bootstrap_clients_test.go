package gloo_test

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/solo-io/gloo/test/kube2e"
	kubetestclients "github.com/solo-io/gloo/test/kubernetes/testutils/clients"
	"github.com/solo-io/go-utils/testutils"

	"github.com/onsi/gomega/gstruct"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kubesecret"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/vault"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	skhelpers "github.com/solo-io/solo-kit/test/helpers"
	corev1 "k8s.io/api/core/v1"

	"github.com/hashicorp/consul/api"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/gloo/test/services"
	skclients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	corecache "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/cache"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/solo-io/gloo/test/gomega"

	vaultapi "github.com/hashicorp/vault/api"
	"github.com/solo-io/gloo/projects/gloo/pkg/bootstrap/clients"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Kubernetes tests for clients generated from projects/gloo/pkg/bootstrap/clients
var _ = Describe("Bootstrap Clients", func() {

	Context("Kube Client Factory", func() {

		It("should set kube rate limits", func() {
			var cfg *rest.Config
			settings := &v1.Settings{
				ConfigSource: &v1.Settings_KubernetesConfigSource{
					KubernetesConfigSource: &v1.Settings_KubernetesCrds{},
				},
				Kubernetes: &v1.Settings_KubernetesConfiguration{
					RateLimits: &v1.Settings_KubernetesConfiguration_RateLimits{
						QPS:   100.5,
						Burst: 1000,
					},
				},
			}
			params := clients.NewConfigFactoryParams(
				settings,
				nil,
				nil,
				&cfg,
				nil,
			)

			kubefactory, err := clients.ConfigFactoryForSettings(params, v1.UpstreamCrd)

			Expect(err).ToNot(HaveOccurred())
			Expect(cfg).ToNot(BeNil())
			Expect(kubefactory.(*factory.KubeResourceClientFactory).Cfg).To(Equal(cfg))

			Expect(cfg.QPS).To(Equal(float32(100.5)))
			Expect(cfg.Burst).To(Equal(1000))
		})

	})

	Context("Artifact Client", func() {

		var (
			testNamespace string

			cfg           *rest.Config
			kubeClient    kubernetes.Interface
			kubeCoreCache corecache.KubeCoreCache
		)

		BeforeEach(func() {
			var err error

			cfg = kubetestclients.MustRestConfig()
			kubeClient = resourceClientset.KubeClients()

			testNamespace = skhelpers.RandString(8)
			_, err = kubeClient.CoreV1().Namespaces().Create(ctx, &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: testNamespace,
				},
			}, metav1.CreateOptions{})
			Expect(err).NotTo(HaveOccurred())
			kubeCoreCache, err = corecache.NewKubeCoreCacheWithOptions(ctx, kubeClient, time.Hour, []string{testNamespace})
			Expect(err).NotTo(HaveOccurred())

		})

		AfterEach(func() {
			err := kubeClient.CoreV1().Namespaces().Delete(ctx, testNamespace, metav1.DeleteOptions{})
			Expect(err).NotTo(HaveOccurred())
		})

		Context("artifacts as config maps", func() {

			var (
				artifactClient v1.ArtifactClient
			)

			BeforeEach(func() {
				_, err := kubeClient.CoreV1().ConfigMaps(testNamespace).Create(ctx,
					&corev1.ConfigMap{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "cfg",
							Namespace: testNamespace,
						},
						Data: map[string]string{
							"test": "data",
						},
					}, metav1.CreateOptions{})
				Expect(err).NotTo(HaveOccurred())

				settings := &v1.Settings{
					ArtifactSource:  &v1.Settings_KubernetesArtifactSource{},
					WatchNamespaces: []string{testNamespace},
				}

				factory, err := clients.ArtifactFactoryForSettings(ctx,
					settings,
					nil,
					&cfg,
					&kubeClient,
					&kubeCoreCache,
					&api.Client{},
					"artifacts")
				Expect(err).NotTo(HaveOccurred())
				artifactClient, err = v1.NewArtifactClient(ctx, factory)
				Expect(err).NotTo(HaveOccurred())
			})

			It("should work with artifacts", func() {
				artifact, err := artifactClient.Read(testNamespace, "cfg", skclients.ReadOpts{Ctx: ctx})
				Expect(err).NotTo(HaveOccurred())
				Expect(artifact.GetMetadata().Name).To(Equal("cfg"))
				Expect(artifact.Data["test"]).To(Equal("data"))
			})
		})
	})

	Context("Secret Client", func() {
		const (
			kubeSecretName  = "kubesecret"
			vaultSecretName = "vaultsecret"
		)
		var (
			vaultInstance  *services.VaultInstance
			secretForVault *v1.Secret

			testNamespace      string
			cfg                *rest.Config
			kubeClient         kubernetes.Interface
			kubeCoreCache      corecache.KubeCoreCache
			secretClient       v1.SecretClient
			settings           *v1.Settings
			vaultClientInitMap map[int]clients.VaultClientInitFunc

			testCtx    context.Context
			testCancel context.CancelFunc
		)

		// setupKubeSecret will
		// - initiate kube clients
		// - create a namespace
		// - create a kubeCoreCache
		// - create a new secret
		// - wait up to 5 seconds to confirm the existence of the secret
		//
		// as-is, this function is not idempotent and should be run only once
		setupKubeSecret := func() {
			var err error
			cfg = kubetestclients.MustRestConfig()
			kubeClient = resourceClientset.KubeClients()

			_, err = kubeClient.CoreV1().Namespaces().Create(ctx, &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: testNamespace,
				},
			}, metav1.CreateOptions{})
			Expect(err).NotTo(HaveOccurred())
			kubeCoreCache, err = corecache.NewKubeCoreCacheWithOptions(ctx, kubeClient, time.Hour, []string{testNamespace})
			Expect(err).NotTo(HaveOccurred())

			kubeSecret := helpers.GetKubeSecret(kubeSecretName, testNamespace)
			_, err = kubeClient.CoreV1().Secrets(testNamespace).Create(ctx,
				kubeSecret,
				metav1.CreateOptions{})
			Expect(err).NotTo(HaveOccurred())

			kubeSecretMatcher := gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"ObjectMeta": gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
					"Name":      Equal(kubeSecretName),
					"Namespace": Equal(testNamespace),
				}),
			})
			Eventually(func(g Gomega) error {
				l, err := kubeClient.CoreV1().Secrets(testNamespace).List(ctx, metav1.ListOptions{})
				if err != nil {
					return err
				}
				g.Expect(l.Items).To(ContainElement(kubeSecretMatcher))
				return nil
			}, "5s", "500ms").ShouldNot(HaveOccurred())
		}

		// setupVaultSecret will
		// - initiate vault instance
		// - create a new secret
		// - wait up to 5 seconds to confirm the existence of the secret
		//
		// as-is, this function is not idempotent and should be run only once
		setupVaultSecret := func() {
			vaultInstance = vaultFactory.MustVaultInstance()
			vaultInstance.Run(testCtx)

			secretForVault = &v1.Secret{
				Kind: &v1.Secret_Tls{},
				Metadata: &core.Metadata{
					Name:      vaultSecretName,
					Namespace: testNamespace,
				},
			}

			vaultInstance.WriteSecret(secretForVault)
			Eventually(func(g Gomega) error {
				// https://developer.hashicorp.com/vault/docs/commands/kv/get
				s, err := vaultInstance.Exec("kv", "get", "-mount=secret", fmt.Sprintf("gloo/gloo.solo.io/v1/Secret/%s/%s", testNamespace, vaultSecretName))
				if err != nil {
					return err
				}
				g.Expect(s).NotTo(BeEmpty())
				return nil
			}, "5s", "500ms").ShouldNot(HaveOccurred())
		}

		getVaultSecrets := func(vi *services.VaultInstance) *v1.Settings_VaultSecrets {
			return &v1.Settings_VaultSecrets{
				Address: vi.Address(),
				AuthMethod: &v1.Settings_VaultSecrets_AccessToken{
					AccessToken: vi.Token(),
				},
			}
		}

		setVaultClientInitMap := func(idx int, vaultSettings *v1.Settings_VaultSecrets) {
			vaultClientInitMap[idx] = func(ctx context.Context) *vaultapi.Client {
				c, err := clients.VaultClientForSettings(ctx, vaultSettings)
				Expect(err).NotTo(HaveOccurred())
				return c
			}
		}

		appendSourceToOptions := func(source *v1.Settings_SecretOptions_Source) {
			secretOpts := settings.GetSecretOptions()
			if secretOpts == nil {
				secretOpts = &v1.Settings_SecretOptions{}
			}
			sources := secretOpts.GetSources()
			if sources == nil {
				sources = make([]*v1.Settings_SecretOptions_Source, 0)
			}
			sources = append(sources, source)

			secretOpts.Sources = sources
			settings.SecretOptions = secretOpts
		}

		BeforeEach(func() {
			testCtx, testCancel = context.WithCancel(ctx)

			testNamespace = skhelpers.RandString(8)
			settings = &v1.Settings{
				WatchNamespaces: []string{testNamespace},
			}
			vaultClientInitMap = make(map[int]clients.VaultClientInitFunc)
		})

		AfterEach(func() {
			testCancel()
		})

		JustBeforeEach(func() {
			factory, err := clients.SecretFactoryForSettings(ctx,
				clients.SecretFactoryParams{
					Settings:           settings,
					SharedCache:        nil,
					Cfg:                &cfg,
					Clientset:          &kubeClient,
					KubeCoreCache:      &kubeCoreCache,
					VaultClientInitMap: vaultClientInitMap,
					PluralName:         "secrets",
				})
			Expect(err).NotTo(HaveOccurred())
			secretClient, err = v1.NewSecretClient(ctx, factory)
			Expect(err).NotTo(HaveOccurred())
		})

		listSecret := func(g Gomega, secretName string) {
			l, err := secretClient.List(testNamespace, skclients.ListOpts{})
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(l).NotTo(BeNil())
			kubeSecret, err := l.Find(testNamespace, secretName)
			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(kubeSecret).NotTo(BeNil())
		}

		When("using secretSource API", func() {
			When("using a kubernetes secret source", func() {
				BeforeEach(func() {
					setupKubeSecret()
					settings.SecretSource = &v1.Settings_KubernetesSecretSource{}
				})
				It("lists secrets", func() {
					Expect(secretClient.BaseClient()).To(BeAssignableToTypeOf(&kubesecret.ResourceClient{}))
					Eventually(func(g Gomega) {
						listSecret(g, kubeSecretName)
					}, DefaultEventuallyTimeout, DefaultEventuallyPollingInterval).Should(Succeed())
				})
			})
			When("using a vault secret source", func() {
				BeforeEach(func() {
					setupVaultSecret()
					vaultSettings := getVaultSecrets(vaultInstance)
					settings.SecretSource = &v1.Settings_VaultSecretSource{
						VaultSecretSource: vaultSettings,
					}
					setVaultClientInitMap(clients.SecretSourceAPIVaultClientInitIndex, vaultSettings)
				})
				It("lists secrets", func() {
					Expect(secretClient.BaseClient()).To(BeAssignableToTypeOf(&vault.ResourceClient{}))
					Eventually(func(g Gomega) {
						listSecret(g, vaultSecretName)
					}, DefaultEventuallyTimeout, DefaultEventuallyPollingInterval).Should(Succeed())
				})
			})

		})
		When("using secretOptions API", func() {
			When("using a single kubernetes secret source", func() {
				BeforeEach(func() {
					setupKubeSecret()
					appendSourceToOptions(
						&v1.Settings_SecretOptions_Source{
							Source: &v1.Settings_SecretOptions_Source_Kubernetes{
								Kubernetes: &v1.Settings_KubernetesSecrets{},
							},
						})
				})
				It("lists secrets", func() {
					Expect(secretClient.BaseClient()).To(BeAssignableToTypeOf(&kubesecret.ResourceClient{}))
					Eventually(func(g Gomega) {
						listSecret(g, kubeSecretName)
					}, DefaultEventuallyTimeout, DefaultEventuallyPollingInterval).Should(Succeed())
				})
			})

			When("using a single vault secret source", func() {
				BeforeEach(func() {
					setupVaultSecret()
					vaultSettings := getVaultSecrets(vaultInstance)
					appendSourceToOptions(
						&v1.Settings_SecretOptions_Source{
							Source: &v1.Settings_SecretOptions_Source_Vault{
								Vault: vaultSettings,
							},
						})
					setVaultClientInitMap(len(settings.GetSecretOptions().GetSources())-1, vaultSettings)

				})
				It("lists secrets", func() {
					Expect(secretClient.BaseClient()).To(BeAssignableToTypeOf(&vault.ResourceClient{}))
					Eventually(func(g Gomega) {
						listSecret(g, vaultSecretName)
					}, DefaultEventuallyTimeout, DefaultEventuallyPollingInterval).Should(Succeed())
				})
			})
			When("using a kubernetes+vault secret source", func() {
				BeforeEach(func() {
					setupKubeSecret()
					appendSourceToOptions(
						&v1.Settings_SecretOptions_Source{
							Source: &v1.Settings_SecretOptions_Source_Kubernetes{
								Kubernetes: &v1.Settings_KubernetesSecrets{},
							},
						})

					setupVaultSecret()
					vaultSettings := getVaultSecrets(vaultInstance)
					appendSourceToOptions(
						&v1.Settings_SecretOptions_Source{
							Source: &v1.Settings_SecretOptions_Source_Vault{
								Vault: vaultSettings,
							},
						})
					setVaultClientInitMap(len(settings.GetSecretOptions().GetSources())-1, vaultSettings)
				})
				It("lists secrets", func() {
					Expect(secretClient.BaseClient()).To(BeAssignableToTypeOf(&clients.MultiSecretResourceClient{}))
					Eventually(func(g Gomega) {
						listSecret(g, kubeSecretName)
					}, DefaultEventuallyTimeout, DefaultEventuallyPollingInterval).Should(Succeed())
					Eventually(func(g Gomega) {
						listSecret(g, vaultSecretName)
					}, DefaultEventuallyTimeout, DefaultEventuallyPollingInterval).Should(Succeed())
				})
			})
		})
	})

	FContext("Retry leader election failure", func() {
		AfterEach(func() {
			ModifyDeploymentEnv(resourceClientset, "gloo", 0, corev1.EnvVar{
				Name:  "RECOVER_FROM_LEADER_ELECTION_FAILURE",
				Value: "false",
			})
		})

		It("does not recover by default", func() {
			waitUntilLeaseAcquired()
			simulateKubeAPIServerUnavailability()

			Eventually(func(g Gomega) {
				logs := getGlooDeploymentLogs()
				g.Expect(logs).To(ContainSubstring("lost leadership, quitting app"))
			}, "30s", "1s")
		})

		It("recovers when RECOVER_FROM_LEADER_ELECTION_FAILURE=true", func() {
			ModifyDeploymentEnv(resourceClientset, "gloo", 0, corev1.EnvVar{
				Name:  "RECOVER_FROM_LEADER_ELECTION_FAILURE",
				Value: "true",
			})

			waitUntilLeaseAcquired()
			simulateKubeAPIServerUnavailability()

			Eventually(func(g Gomega) {
				logs := getGlooDeploymentLogs()
				g.Expect(logs).To(ContainSubstring("Leader election cycle 0 lost. Trying again"))
				g.Expect(logs).To(ContainSubstring("recovered from lease renewal failure"))
				g.Expect(logs).NotTo(ContainSubstring("lost leadership, quitting app"))
			}, "30s", "1s")
		})
	})
})

func ModifyDeploymentEnv(resourceClientset *kube2e.KubeResourceClientSet, deploymentName string, containerIndex int, envVar corev1.EnvVar) {
	deploymentClient := resourceClientset.KubeClients().AppsV1().Deployments(namespace)

	d, err := deploymentClient.Get(ctx, deploymentName, metav1.GetOptions{})
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	// make sure we are referencing a valid container
	ExpectWithOffset(1, len(d.Spec.Template.Spec.Containers)).To(BeNumerically(">", containerIndex))

	// if an env var with the given name already exists, modify it
	exists := false
	for i, env := range d.Spec.Template.Spec.Containers[containerIndex].Env {
		if env.Name == envVar.Name {
			d.Spec.Template.Spec.Containers[containerIndex].Env[i].Value = envVar.Value
			exists = true
			break
		}
	}
	// otherwise add a new env var
	if !exists {
		d.Spec.Template.Spec.Containers[containerIndex].Env = append(d.Spec.Template.Spec.Containers[containerIndex].Env, envVar)
	}
	_, err = deploymentClient.Update(ctx, d, metav1.UpdateOptions{})
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	WaitForRolloutWithOffset(1, deploymentName, namespace, "60s", "1s")
}

// WaitForRollout waits for the specified deployment to be rolled out successfully.
func WaitForRollout(deploymentName string, deploymentNamespace string, intervals ...interface{}) {
	WaitForRolloutWithOffset(1, deploymentName, deploymentNamespace, intervals...)
}

func WaitForRolloutWithOffset(offset int, deploymentName string, deploymentNamespace string, intervals ...interface{}) {
	EventuallyWithOffset(offset+1, func() (bool, error) {
		out, err := testutils.KubectlOut("rollout", "status", "-n", deploymentNamespace, fmt.Sprintf("deployment/%s", deploymentName))
		fmt.Println(out)
		return strings.Contains(out, "successfully rolled out"), err
	}, "30s", "1s").Should(BeTrue())
}

func simulateKubeAPIServerUnavailability() {
	out, err := testutils.KubectlOut("apply", "-f", testHelper.RootDir+"/test/kube2e/gloo/artifacts/block.yaml")
	Expect(err).ToNot(HaveOccurred())
	fmt.Println(out)
	time.Sleep(15 * time.Second)
	out, err = testutils.KubectlOut("delete", "-f", testHelper.RootDir+"/test/kube2e/gloo/artifacts/block.yaml")
	fmt.Println(out)
	Expect(err).ToNot(HaveOccurred())
}

func getGlooDeploymentLogs() string {
	out, err := testutils.KubectlOut("-n", "gloo-system", "logs", "deploy/gloo")
	Expect(err).ToNot(HaveOccurred())
	fmt.Println(out)
	return out
}

func waitUntilLeaseAcquired() {
	Eventually(func(g Gomega) {
		out := getGlooDeploymentLogs()
		g.Expect(out).To(ContainSubstring("successfully acquired lease gloo-system/gloo"))
	}, "30s", "2s")
	time.Sleep(30 * time.Second)
}
