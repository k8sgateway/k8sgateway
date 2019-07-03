package install

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

var _ = Describe("Knative", func() {
	knativeInstallOpts := options.Knative{
		InstallKnativeVersion:    "0.7.0",
		InstallKnativeBuild:      true,
		InstallKnativeEventing:   true,
		InstallKnativeMonitoring: true,
	}
	Context("RenderKnativeManifests", func() {
		It("renders manifests for each knative component", func() {
			manifests, err := RenderKnativeManifests(knativeInstallOpts)
			Expect(err).NotTo(HaveOccurred())
			Expect(manifests).To(Equal(expected070Manifests))
		})
	})
	Context("checkKnativeInstallation", func() {
		It("returns true, opts if knative was installed by us", func() {
			optsJson, err := json.Marshal(knativeInstallOpts)
			Expect(err).NotTo(HaveOccurred())
			kc := fake.NewSimpleClientset()
			_, err = kc.CoreV1().Namespaces().Create(&v1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:   "knative-serving",
					Labels: map[string]string{installedByUsAnnotationKey: string(optsJson)},
				},
			})
			Expect(err).NotTo(HaveOccurred())

			installed, opts, err := checkKnativeInstallation(kc)
			Expect(err).NotTo(HaveOccurred())
			Expect(installed).To(BeTrue())
			Expect(opts).To(Equal(&knativeInstallOpts))
		})
		It("returns true, nil if knative was installed but not by us", func() {
			kc := fake.NewSimpleClientset()
			_, err := kc.CoreV1().Namespaces().Create(&v1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: "knative-serving",
				},
			})
			Expect(err).NotTo(HaveOccurred())

			installed, opts, err := checkKnativeInstallation(kc)
			Expect(err).NotTo(HaveOccurred())
			Expect(installed).To(BeTrue())
			Expect(opts).To(BeNil())
		})
		It("returns false, nil if knative was not installed", func() {
			kc := fake.NewSimpleClientset()

			installed, opts, err := checkKnativeInstallation(kc)
			Expect(err).NotTo(HaveOccurred())
			Expect(installed).To(BeFalse())
			Expect(opts).To(BeNil())
		})
	})
})

// does not contain any networking.knative.dev/ingress-provider=istio resources
const expected070Manifests = `apiVersion: v1
kind: Namespace
metadata:
  labels:
    istio-injection: enabled
    serving.knative.dev/release: "v0.7.0"
  name: knative-serving
---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    networking.knative.dev/certificate-provider: cert-manager
    serving.knative.dev/controller: "true"
    serving.knative.dev/release: "v0.7.0"
  name: knative-serving-certmanager
rules:
- apiGroups:
  - certmanager.k8s.io
  resources:
  - certificates
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    autoscaling.knative.dev/metric-provider: custom-metrics
    serving.knative.dev/release: "v0.7.0"
  name: custom-metrics-server-resources
rules:
- apiGroups:
  - custom.metrics.k8s.io
  resources:
  - '*'
  verbs:
  - '*'
---

aggregationRule:
  clusterRoleSelectors:
  - matchLabels:
      serving.knative.dev/controller: "true"
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: knative-serving-admin
rules: []---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    serving.knative.dev/controller: "true"
    serving.knative.dev/release: "v0.7.0"
  name: knative-serving-core
rules:
- apiGroups:
  - ""
  resources:
  - pods
  - namespaces
  - secrets
  - configmaps
  - endpoints
  - services
  - events
  - serviceaccounts
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
- apiGroups:
  - ""
  resources:
  - endpoints/restricted
  verbs:
  - create
- apiGroups:
  - apps
  resources:
  - deployments
  - deployments/finalizers
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
- apiGroups:
  - admissionregistration.k8s.io
  resources:
  - mutatingwebhookconfigurations
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
- apiGroups:
  - autoscaling
  resources:
  - horizontalpodautoscalers
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
- apiGroups:
  - serving.knative.dev
  - autoscaling.internal.knative.dev
  - networking.internal.knative.dev
  resources:
  - '*'
  - '*/status'
  - '*/finalizers'
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - deletecollection
  - patch
  - watch
- apiGroups:
  - caching.internal.knative.dev
  resources:
  - images
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
---

apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: controller
  namespace: knative-serving
---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    autoscaling.knative.dev/metric-provider: custom-metrics
    serving.knative.dev/release: "v0.7.0"
  name: custom-metrics:system:auth-delegator
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:auth-delegator
subjects:
- kind: ServiceAccount
  name: controller
  namespace: knative-serving
---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    autoscaling.knative.dev/metric-provider: custom-metrics
    serving.knative.dev/release: "v0.7.0"
  name: hpa-controller-custom-metrics
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: custom-metrics-server-resources
subjects:
- kind: ServiceAccount
  name: horizontal-pod-autoscaler
  namespace: kube-system
---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: knative-serving-controller-admin
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: knative-serving-admin
subjects:
- kind: ServiceAccount
  name: controller
  namespace: knative-serving
---

apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  labels:
    autoscaling.knative.dev/metric-provider: custom-metrics
    serving.knative.dev/release: "v0.7.0"
  name: custom-metrics-auth-reader
  namespace: kube-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: extension-apiserver-authentication-reader
subjects:
- kind: ServiceAccount
  name: controller
  namespace: knative-serving
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    knative.dev/crd-install: "true"
    serving.knative.dev/release: "v0.7.0"
  name: certificates.networking.internal.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .status.conditions[?(@.type=="Ready")].status
    name: Ready
    type: string
  - JSONPath: .status.conditions[?(@.type=="Ready")].reason
    name: Reason
    type: string
  group: networking.internal.knative.dev
  names:
    categories:
    - all
    - knative-internal
    - networking
    kind: Certificate
    plural: certificates
    shortNames:
    - kcert
    singular: certificate
  scope: Namespaced
  subresources:
    status: {}
  version: v1alpha1
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    knative.dev/crd-install: "true"
    serving.knative.dev/release: "v0.7.0"
  name: clusteringresses.networking.internal.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .status.conditions[?(@.type=='Ready')].status
    name: Ready
    type: string
  - JSONPath: .status.conditions[?(@.type=='Ready')].reason
    name: Reason
    type: string
  group: networking.internal.knative.dev
  names:
    categories:
    - all
    - knative-internal
    - networking
    kind: ClusterIngress
    plural: clusteringresses
    singular: clusteringress
  scope: Cluster
  subresources:
    status: {}
  version: v1alpha1
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    knative.dev/crd-install: "true"
    serving.knative.dev/release: "v0.7.0"
  name: configurations.serving.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .status.latestCreatedRevisionName
    name: LatestCreated
    type: string
  - JSONPath: .status.latestReadyRevisionName
    name: LatestReady
    type: string
  - JSONPath: .status.conditions[?(@.type=='Ready')].status
    name: Ready
    type: string
  - JSONPath: .status.conditions[?(@.type=='Ready')].reason
    name: Reason
    type: string
  group: serving.knative.dev
  names:
    categories:
    - all
    - knative
    - serving
    kind: Configuration
    plural: configurations
    shortNames:
    - config
    - cfg
    singular: configuration
  scope: Namespaced
  subresources:
    status: {}
  versions:
  - name: v1alpha1
    served: true
    storage: true
  - name: v1beta1
    served: true
    storage: false
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    knative.dev/crd-install: "true"
  name: images.caching.internal.knative.dev
spec:
  group: caching.internal.knative.dev
  names:
    categories:
    - all
    - knative-internal
    - caching
    kind: Image
    plural: images
    shortNames:
    - img
    singular: image
  scope: Namespaced
  subresources:
    status: {}
  version: v1alpha1
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    knative.dev/crd-install: "true"
    serving.knative.dev/release: "v0.7.0"
  name: ingresses.networking.internal.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .status.conditions[?(@.type=='Ready')].status
    name: Ready
    type: string
  - JSONPath: .status.conditions[?(@.type=='Ready')].reason
    name: Reason
    type: string
  group: networking.internal.knative.dev
  names:
    categories:
    - all
    - knative-internal
    - networking
    kind: Ingress
    plural: ingresses
    shortNames:
    - ing
    singular: ingress
  scope: Namespaced
  subresources:
    status: {}
  version: v1alpha1
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    knative.dev/crd-install: "true"
    serving.knative.dev/release: "v0.7.0"
  name: podautoscalers.autoscaling.internal.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .status.conditions[?(@.type=='Ready')].status
    name: Ready
    type: string
  - JSONPath: .status.conditions[?(@.type=='Ready')].reason
    name: Reason
    type: string
  group: autoscaling.internal.knative.dev
  names:
    categories:
    - all
    - knative-internal
    - autoscaling
    kind: PodAutoscaler
    plural: podautoscalers
    shortNames:
    - kpa
    singular: podautoscaler
  scope: Namespaced
  subresources:
    status: {}
  version: v1alpha1
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    knative.dev/crd-install: "true"
    serving.knative.dev/release: "v0.7.0"
  name: revisions.serving.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .status.serviceName
    name: Service Name
    type: string
  - JSONPath: .metadata.labels['serving\.knative\.dev/configurationGeneration']
    name: Generation
    type: string
  - JSONPath: .status.conditions[?(@.type=='Ready')].status
    name: Ready
    type: string
  - JSONPath: .status.conditions[?(@.type=='Ready')].reason
    name: Reason
    type: string
  group: serving.knative.dev
  names:
    categories:
    - all
    - knative
    - serving
    kind: Revision
    plural: revisions
    shortNames:
    - rev
    singular: revision
  scope: Namespaced
  subresources:
    status: {}
  versions:
  - name: v1alpha1
    served: true
    storage: true
  - name: v1beta1
    served: true
    storage: false
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    knative.dev/crd-install: "true"
    serving.knative.dev/release: "v0.7.0"
  name: routes.serving.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .status.url
    name: URL
    type: string
  - JSONPath: .status.conditions[?(@.type=='Ready')].status
    name: Ready
    type: string
  - JSONPath: .status.conditions[?(@.type=='Ready')].reason
    name: Reason
    type: string
  group: serving.knative.dev
  names:
    categories:
    - all
    - knative
    - serving
    kind: Route
    plural: routes
    shortNames:
    - rt
    singular: route
  scope: Namespaced
  subresources:
    status: {}
  versions:
  - name: v1alpha1
    served: true
    storage: true
  - name: v1beta1
    served: true
    storage: false
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    knative.dev/crd-install: "true"
    serving.knative.dev/release: "v0.7.0"
  name: services.serving.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .status.url
    name: URL
    type: string
  - JSONPath: .status.latestCreatedRevisionName
    name: LatestCreated
    type: string
  - JSONPath: .status.latestReadyRevisionName
    name: LatestReady
    type: string
  - JSONPath: .status.conditions[?(@.type=='Ready')].status
    name: Ready
    type: string
  - JSONPath: .status.conditions[?(@.type=='Ready')].reason
    name: Reason
    type: string
  group: serving.knative.dev
  names:
    categories:
    - all
    - knative
    - serving
    kind: Service
    plural: services
    shortNames:
    - kservice
    - ksvc
    singular: service
  scope: Namespaced
  subresources:
    status: {}
  versions:
  - name: v1alpha1
    served: true
    storage: true
  - name: v1beta1
    served: true
    storage: false
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    knative.dev/crd-install: "true"
    serving.knative.dev/release: "v0.7.0"
  name: serverlessservices.networking.internal.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .spec.mode
    name: Mode
    type: string
  - JSONPath: .status.serviceName
    name: ServiceName
    type: string
  - JSONPath: .status.privateServiceName
    name: PrivateServiceName
    type: string
  - JSONPath: .status.conditions[?(@.type=='Ready')].status
    name: Ready
    type: string
  - JSONPath: .status.conditions[?(@.type=='Ready')].reason
    name: Reason
    type: string
  group: networking.internal.knative.dev
  names:
    categories:
    - all
    - knative-internal
    - networking
    kind: ServerlessService
    plural: serverlessservices
    shortNames:
    - sks
    singular: serverlessservice
  scope: Namespaced
  subresources:
    status: {}
  version: v1alpha1
---

apiVersion: v1
kind: Service
metadata:
  labels:
    app: activator
    serving.knative.dev/release: "v0.7.0"
  name: activator-service
  namespace: knative-serving
spec:
  ports:
  - name: http
    port: 80
    protocol: TCP
    targetPort: 8012
  - name: http2
    port: 81
    protocol: TCP
    targetPort: 8013
  - name: metrics
    port: 9090
    protocol: TCP
    targetPort: 9090
  selector:
    app: activator
  type: ClusterIP
---

apiVersion: v1
kind: Service
metadata:
  labels:
    app: controller
    serving.knative.dev/release: "v0.7.0"
  name: controller
  namespace: knative-serving
spec:
  ports:
  - name: metrics
    port: 9090
    protocol: TCP
    targetPort: 9090
  selector:
    app: controller
---

apiVersion: v1
kind: Service
metadata:
  labels:
    role: webhook
    serving.knative.dev/release: "v0.7.0"
  name: webhook
  namespace: knative-serving
spec:
  ports:
  - port: 443
    targetPort: 8443
  selector:
    role: webhook
---

apiVersion: caching.internal.knative.dev/v1alpha1
kind: Image
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: queue-proxy
  namespace: knative-serving
spec:
  image: gcr.io/knative-releases/github.com/knative/serving/cmd/queue@sha256:e007c0a78c541600466f88954deee65c517246a23345bfba45a7f212d09b8f3b
---

apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: activator
  namespace: knative-serving
spec:
  selector:
    matchLabels:
      app: activator
      role: activator
  template:
    metadata:
      annotations:
        cluster-autoscaler.kubernetes.io/safe-to-evict: "false"
        sidecar.istio.io/inject: "true"
      labels:
        app: activator
        role: activator
        serving.knative.dev/release: "v0.7.0"
    spec:
      containers:
      - args:
        - -logtostderr=false
        - -stderrthreshold=FATAL
        env:
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: SYSTEM_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: CONFIG_LOGGING_NAME
          value: config-logging
        - name: CONFIG_OBSERVABILITY_NAME
          value: config-observability
        - name: METRICS_DOMAIN
          value: knative.dev/serving
        image: gcr.io/knative-releases/github.com/knative/serving/cmd/activator@sha256:57fe5f1a8b1d12f29fe9e3a904b00c7219e5ce5825d94f33339db929e92257db
        livenessProbe:
          httpGet:
            httpHeaders:
            - name: k-kubelet-probe
              value: activator
            path: /healthz
            port: 8012
        name: activator
        ports:
        - containerPort: 8012
          name: http1-port
        - containerPort: 8013
          name: h2c-port
        - containerPort: 9090
          name: metrics-port
        readinessProbe:
          httpGet:
            httpHeaders:
            - name: k-kubelet-probe
              value: activator
            path: /healthz
            port: 8012
        resources:
          limits:
            cpu: 200m
            memory: 600Mi
          requests:
            cpu: 20m
            memory: 60Mi
        securityContext:
          allowPrivilegeEscalation: false
        volumeMounts:
        - mountPath: /etc/config-logging
          name: config-logging
        - mountPath: /etc/config-observability
          name: config-observability
      serviceAccountName: controller
      volumes:
      - configMap:
          name: config-logging
        name: config-logging
      - configMap:
          name: config-observability
        name: config-observability
---

apiVersion: v1
kind: Service
metadata:
  labels:
    app: autoscaler
    serving.knative.dev/release: "v0.7.0"
  name: autoscaler
  namespace: knative-serving
spec:
  ports:
  - name: http
    port: 8080
    protocol: TCP
    targetPort: 8080
  - name: metrics
    port: 9090
    protocol: TCP
    targetPort: 9090
  - name: custom-metrics
    port: 443
    protocol: TCP
    targetPort: 8443
  selector:
    app: autoscaler
---

apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: autoscaler
  namespace: knative-serving
spec:
  replicas: 1
  selector:
    matchLabels:
      app: autoscaler
  template:
    metadata:
      annotations:
        cluster-autoscaler.kubernetes.io/safe-to-evict: "false"
        sidecar.istio.io/inject: "true"
      labels:
        app: autoscaler
        serving.knative.dev/release: "v0.7.0"
    spec:
      containers:
      - args:
        - --secure-port=8443
        - --cert-dir=/tmp
        env:
        - name: SYSTEM_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: CONFIG_LOGGING_NAME
          value: config-logging
        - name: CONFIG_OBSERVABILITY_NAME
          value: config-observability
        - name: METRICS_DOMAIN
          value: knative.dev/serving
        image: gcr.io/knative-releases/github.com/knative/serving/cmd/autoscaler@sha256:2c2370df2751741348e1cc456f31425cb2455c377ddb45d3f6c17e743fd63d78
        livenessProbe:
          httpGet:
            httpHeaders:
            - name: k-kubelet-probe
              value: autoscaler
            path: /healthz
            port: 8080
        name: autoscaler
        ports:
        - containerPort: 8080
          name: websocket
        - containerPort: 9090
          name: metrics
        - containerPort: 8443
          name: custom-metrics
        readinessProbe:
          httpGet:
            httpHeaders:
            - name: k-kubelet-probe
              value: autoscaler
            path: /healthz
            port: 8080
        resources:
          limits:
            cpu: 300m
            memory: 400Mi
          requests:
            cpu: 30m
            memory: 40Mi
        securityContext:
          allowPrivilegeEscalation: false
        volumeMounts:
        - mountPath: /etc/config-autoscaler
          name: config-autoscaler
        - mountPath: /etc/config-logging
          name: config-logging
        - mountPath: /etc/config-observability
          name: config-observability
      serviceAccountName: controller
      volumes:
      - configMap:
          name: config-autoscaler
        name: config-autoscaler
      - configMap:
          name: config-logging
        name: config-logging
      - configMap:
          name: config-observability
        name: config-observability
---

apiVersion: v1
data:
  _example: |
    ################################
    #                              #
    #    EXAMPLE CONFIGURATION     #
    #                              #
    ################################

    # This block is not actually functional configuration,
    # but serves to illustrate the available configuration
    # options and document them in a way that is accessible
    # to users that ` + "`" + `kubectl edit` + "`" + ` this config map.
    #
    # These sample configuration options may be copied out of
    # this example block and unindented to be in the data block
    # to actually change the configuration.

    # The Revision ContainerConcurrency field specifies the maximum number
    # of requests the Container can handle at once. Container concurrency
    # target percentage is how much of that maximum to use in a stable
    # state. E.g. if a Revision specifies ContainerConcurrency of 10, then
    # the Autoscaler will try to maintain 7 concurrent connections per pod
    # on average. A value of 70 is chosen because the Autoscaler panics
    # when concurrency exceeds 2x the desired set point. So we will panic
    # before we reach the limit.
    # For legacy and backwards compatibility reasons, this value also accepts
    # fractional values in (0, 1] interval (i.e. 0.7 ⇒ 70%).
    # Thus minimal percentage value must be greater than 1.0, or it will be
    # treated as a fraction.
    # TODO(#2016): Set to 70%.
    container-concurrency-target-percentage: "100"

    # The container concurrency target default is what the Autoscaler will
    # try to maintain when the Revision specifies unlimited concurrency.
    # Even when specifying unlimited concurrency, the autoscaler will
    # horizontally scale the application based on this target concurrency.
    #
    # A value of 100 is chosen because it's enough to allow vertical pod
    # autoscaling to tune resource requests. E.g. maintaining 1 concurrent
    # "hello world" request doesn't consume enough resources to allow VPA
    # to achieve efficient resource usage (VPA CPU minimum is 300m).
    container-concurrency-target-default: "100"

    # When operating in a stable mode, the autoscaler operates on the
    # average concurrency over the stable window.
    stable-window: "60s"

    # When observed average concurrency during the panic window reaches
    # panic-threshold-percentage the target concurrency, the autoscaler
    # enters panic mode. When operating in panic mode, the autoscaler
    # scales on the average concurrency over the panic window which is
    # panic-window-percentage of the stable-window.
    panic-window-percentage: "10.0"

    # Absolute panic window duration.
    # Deprecated in favor of panic-window-percentage.
    # Existing revisions will continue to scale based on panic-window
    # but new revisions will default to panic-window-percentage.
    panic-window: "6s"

    # The percentage of the container concurrency target at which to
    # enter panic mode when reached within the panic window.
    panic-threshold-percentage: "200.0"

    # Max scale up rate limits the rate at which the autoscaler will
    # increase pod count. It is the maximum ratio of desired pods versus
    # observed pods.
    max-scale-up-rate: "10"

    # Scale to zero feature flag
    enable-scale-to-zero: "true"

    # Tick interval is the time between autoscaling calculations.
    tick-interval: "2s"

    # Dynamic parameters (take effect when config map is updated):

    # Scale to zero grace period is the time an inactive revision is left
    # running before it is scaled to zero (min: 30s).
    scale-to-zero-grace-period: "30s"
kind: ConfigMap
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: config-autoscaler
  namespace: knative-serving
---

apiVersion: v1
data:
  _example: |
    ################################
    #                              #
    #    EXAMPLE CONFIGURATION     #
    #                              #
    ################################

    # This block is not actually functional configuration,
    # but serves to illustrate the available configuration
    # options and document them in a way that is accessible
    # to users that ` + "`" + `kubectl edit` + "`" + ` this config map.
    #
    # These sample configuration options may be copied out of
    # this block and unindented to actually change the configuration.

    # IssuerRef is a reference to the issuer for this certificate.
    # IssuerRef should be either ` + "`" + `ClusterIssuer` + "`" + ` or ` + "`" + `Issuer` + "`" + `.
    # Please refer ` + "`" + `IssuerRef` + "`" + ` in https://github.com/jetstack/cert-manager/blob/master/pkg/apis/certmanager/v1alpha1/types_certificate.go
    # for more details about IssuerRef configuration.
    issuerRef: |
      kind: ClusterIssuer
      name: letsencrypt-issuer

    # solverConfig defines the configuration for the ACME certificate provider.
    # The solverConfig should be either dns01 or http01.
    # Please refer ` + "`" + `SolverConfig` + "`" + ` in https://github.com/jetstack/cert-manager/blob/master/pkg/apis/certmanager/v1alpha1/types_certificate.go
    # for more details about ACME configuration.
    solverConfig: |
      dns01:
        provider: cloud-dns-provider
kind: ConfigMap
metadata:
  labels:
    networking.knative.dev/certificate-provider: cert-manager
    serving.knative.dev/release: "v0.7.0"
  name: config-certmanager
  namespace: knative-serving
---

apiVersion: v1
data:
  _example: |
    ################################
    #                              #
    #    EXAMPLE CONFIGURATION     #
    #                              #
    ################################

    # This block is not actually functional configuration,
    # but serves to illustrate the available configuration
    # options and document them in a way that is accessible
    # to users that ` + "`" + `kubectl edit` + "`" + ` this config map.
    #
    # These sample configuration options may be copied out of
    # this example block and unindented to be in the data block
    # to actually change the configuration.

    # revision-timeout-seconds contains the default number of
    # seconds to use for the revision's per-request timeout, if
    # none is specified.
    revision-timeout-seconds: "300"  # 5 minutes

    # max-revision-timeout-seconds contains the maximum number of
    # seconds that can be used for revision-timeout-seconds.
    # This value must be greater than or equal to revision-timeout-seconds.
    # If omitted, the system default is used (600 seconds).
    max-revision-timeout-seconds: "600"  # 10 minutes

    # revision-cpu-request contains the cpu allocation to assign
    # to revisions by default.  If omitted, no value is specified
    # and the system default is used.
    revision-cpu-request: "400m"  # 0.4 of a CPU (aka 400 milli-CPU)

    # revision-memory-request contains the memory allocation to assign
    # to revisions by default.  If omitted, no value is specified
    # and the system default is used.
    revision-memory-request: "100M"  # 100 megabytes of memory

    # revision-cpu-limit contains the cpu allocation to limit
    # revisions to by default.  If omitted, no value is specified
    # and the system default is used.
    revision-cpu-limit: "1000m"  # 1 CPU (aka 1000 milli-CPU)

    # revision-memory-limit contains the memory allocation to limit
    # revisions to by default.  If omitted, no value is specified
    # and the system default is used.
    revision-memory-limit: "200M"  # 200 megabytes of memory

    # container-name-template contains a template for the default
    # container name, if none is specified.  This field supports
    # Go templating and is supplied with the ObjectMeta of the
    # enclosing Service or Configuration, so values such as
    # {{.Name}} are also valid.
    container-name-template: "user-container"
kind: ConfigMap
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: config-defaults
  namespace: knative-serving
---

apiVersion: v1
data:
  _example: |
    ################################
    #                              #
    #    EXAMPLE CONFIGURATION     #
    #                              #
    ################################

    # This block is not actually functional configuration,
    # but serves to illustrate the available configuration
    # options and document them in a way that is accessible
    # to users that ` + "`" + `kubectl edit` + "`" + ` this config map.
    #
    # These sample configuration options may be copied out of
    # this example block and unindented to be in the data block
    # to actually change the configuration.

    # List of repositories for which tag to digest resolving should be skipped
    registriesSkippingTagResolving: "ko.local,dev.local"
  queueSidecarImage: gcr.io/knative-releases/github.com/knative/serving/cmd/queue@sha256:e007c0a78c541600466f88954deee65c517246a23345bfba45a7f212d09b8f3b
kind: ConfigMap
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: config-deployment
  namespace: knative-serving
---

apiVersion: v1
data:
  _example: |
    ################################
    #                              #
    #    EXAMPLE CONFIGURATION     #
    #                              #
    ################################

    # This block is not actually functional configuration,
    # but serves to illustrate the available configuration
    # options and document them in a way that is accessible
    # to users that ` + "`" + `kubectl edit` + "`" + ` this config map.
    #
    # These sample configuration options may be copied out of
    # this example block and unindented to be in the data block
    # to actually change the configuration.

    # Default value for domain.
    # Although it will match all routes, it is the least-specific rule so it
    # will only be used if no other domain matches.
    example.com: |

    # These are example settings of domain.
    # example.org will be used for routes having app=nonprofit.
    example.org: |
      selector:
        app: nonprofit

    # Routes having domain suffix of 'svc.cluster.local' will not be exposed
    # through Ingress. You can define your own label selector to assign that
    # domain suffix to your Route here, or you can set the label
    #    "serving.knative.dev/visibility=cluster-local"
    # to achieve the same effect.  This shows how to make routes having
    # the label app=secret only exposed to the local cluster.
    svc.cluster.local: |
      selector:
        app: secret
kind: ConfigMap
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: config-domain
  namespace: knative-serving
---

apiVersion: v1
data:
  _example: |
    ################################
    #                              #
    #    EXAMPLE CONFIGURATION     #
    #                              #
    ################################

    # This block is not actually functional configuration,
    # but serves to illustrate the available configuration
    # options and document them in a way that is accessible
    # to users that ` + "`" + `kubectl edit` + "`" + ` this config map.
    #
    # These sample configuration options may be copied out of
    # this example block and unindented to be in the data block
    # to actually change the configuration.

    # Delay after revision creation before considering it for GC
    stale-revision-create-delay: "24h"

    # Duration since a route has been pointed at a revision before it should be GC'd
    # This minus lastpinned-debounce be longer than the controller resync period (10 hours)
    stale-revision-timeout: "15h"

    # Minimum number of generations of revisions to keep before considering for GC
    stale-revision-minimum-generations: "1"

    # To avoid constant updates, we allow an existing annotation to be stale by this
    # amount before we update the timestamp
    stale-revision-lastpinned-debounce: "5h"
kind: ConfigMap
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: config-gc
  namespace: knative-serving
---

apiVersion: v1
data:
  _example: |
    ################################
    #                              #
    #    EXAMPLE CONFIGURATION     #
    #                              #
    ################################

    # This block is not actually functional configuration,
    # but serves to illustrate the available configuration
    # options and document them in a way that is accessible
    # to users that ` + "`" + `kubectl edit` + "`" + ` this config map.
    #
    # These sample configuration options may be copied out of
    # this example block and unindented to be in the data block
    # to actually change the configuration.

    # Common configuration for all Knative codebase
    zap-logger-config: |
      {
        "level": "info",
        "development": false,
        "outputPaths": ["stdout"],
        "errorOutputPaths": ["stderr"],
        "encoding": "json",
        "encoderConfig": {
          "timeKey": "ts",
          "levelKey": "level",
          "nameKey": "logger",
          "callerKey": "caller",
          "messageKey": "msg",
          "stacktraceKey": "stacktrace",
          "lineEnding": "",
          "levelEncoder": "",
          "timeEncoder": "iso8601",
          "durationEncoder": "",
          "callerEncoder": ""
        }
      }

    # Log level overrides
    # For all components except the autoscaler and queue proxy,
    # changes are be picked up immediately.
    # For autoscaler and queue proxy, changes require recreation of the pods.
    loglevel.controller: "info"
    loglevel.autoscaler: "info"
    loglevel.queueproxy: "info"
    loglevel.webhook: "info"
    loglevel.activator: "info"
kind: ConfigMap
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: config-logging
  namespace: knative-serving
---

apiVersion: v1
data:
  _example: |
    ################################
    #                              #
    #    EXAMPLE CONFIGURATION     #
    #                              #
    ################################

    # This block is not actually functional configuration,
    # but serves to illustrate the available configuration
    # options and document them in a way that is accessible
    # to users that ` + "`" + `kubectl edit` + "`" + ` this config map.
    #
    # These sample configuration options may be copied out of
    # this example block and unindented to be in the data block
    # to actually change the configuration.

    # istio.sidecar.includeOutboundIPRanges specifies the IP ranges that Istio sidecar
    # will intercept.
    #
    # Replace this with the IP ranges of your cluster (see below for some examples).
    # Separate multiple entries with a comma.
    # Example: "10.4.0.0/14,10.7.240.0/20"
    #
    # If set to "*" Istio will intercept all traffic within
    # the cluster as well as traffic that is going outside the cluster.
    # Traffic going outside the cluster will be blocked unless
    # necessary egress rules are created.
    #
    # If omitted or set to "", value of global.proxy.includeIPRanges
    # provided at Istio deployment time is used. In default Knative serving
    # deployment, global.proxy.includeIPRanges value is set to "*".
    #
    # If an invalid value is passed, "" is used instead.
    #
    # If valid set of IP address ranges are put into this value,
    # Istio will no longer intercept traffic going to IP addresses
    # outside the provided ranges and there is no need to specify
    # egress rules.
    #
    # To determine the IP ranges of your cluster:
    #   IBM Cloud Private: cat cluster/config.yaml | grep service_cluster_ip_range
    #   IBM Cloud Kubernetes Service: "172.30.0.0/16,172.20.0.0/16,10.10.10.0/24"
    #   Google Container Engine (GKE): gcloud container clusters describe XXXXXXX --zone=XXXXXX | grep -e clusterIpv4Cidr -e servicesIpv4Cidr
    #   Azure Kubernetes Service (AKS): "10.0.0.0/16"
    #   Azure Container Service (ACS; deprecated): "10.244.0.0/16,10.240.0.0/16"
    #   Azure Container Service Engine (ACS-Engine; OSS): Configurable, but defaults to "10.0.0.0/16"
    #   Minikube: "10.0.0.1/24"
    #
    # For more information, visit
    # https://istio.io/docs/tasks/traffic-management/egress/
    #
    istio.sidecar.includeOutboundIPRanges: "*"

    # clusteringress.class specifies the default cluster ingress class
    # to use when not dictated by Route annotation.
    #
    # If not specified, will use the Istio ingress.
    #
    # Note that changing the ClusterIngress class of an existing Route
    # will result in undefined behavior.  Therefore it is best to only
    # update this value during the setup of Knative, to avoid getting
    # undefined behavior.
    clusteringress.class: "istio.ingress.networking.knative.dev"

    # domainTemplate specifies the golang text template string to use
    # when constructing the Knative service's DNS name. The default
    # value is "{{.Name}}.{{.Namespace}}.{{.Domain}}". And those three
    # values (Name, Namespace, Domain) are the only variables defined.
    #
    # Changing this value might be necessary when the extra levels in
    # the domain name generated is problematic for wildcard certificates
    # that only support a single level of domain name added to the
    # certificate's domain. In those cases you might consider using a value
    # of "{{.Name}}-{{.Namespace}}.{{.Domain}}", or removing the Namespace
    # entirely from the template. When choosing a new value be thoughtful
    # of the potential for conflicts - for example, when users choose to use
    # characters such as ` + "`" + `-` + "`" + ` in their service, or namespace, names.
    # {{.Annotations}} can be used for any customization in the go template if needed.
    # We strongly recommend keeping namespace part of the template to avoid domain name clashes
    # Example '{{.Name}}-{{.Namespace}}.{{ index .Annotations "sub"}}.{{.Domain}}'
    # and you have an annotation {"sub":"foo"}, then the generated template would be {Name}-{Namespace}.foo.{Domain}
    domainTemplate: "{{.Name}}.{{.Namespace}}.{{.Domain}}"

    # tagTemplate specifies the golang text template string to use
    # when constructing the DNS name for "tags" within the traffic blocks
    # of Routes and Configuration.  This is used in conjunction with the
    # domainTemplate above to determine the full URL for the tag.
    tagTemplate: "{{.Name}}-{{.Tag}}"

    # Controls whether TLS certificates are automatically provisioned and
    # installed in the Knative ingress to terminate external TLS connection.
    # 1. Enabled: enabling auto-TLS feature.
    # 2. Disabled: disabling auto-TLS feature.
    autoTLS: "Disabled"

    # Controls the behavior of the HTTP endpoint for the Knative ingress.
    # It requires autoTLS to be enabled.
    # 1. Enabled: The Knative ingress will be able to serve HTTP connection.
    # 2. Disabled: The Knative ingress ter will reject HTTP traffic.
    # 3. Redirected: The Knative ingress will send a 302 redirect for all
    # http connections, asking the clients to use HTTPS
    httpProtocol: "Enabled"
kind: ConfigMap
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: config-network
  namespace: knative-serving
---

apiVersion: v1
data:
  _example: |
    ################################
    #                              #
    #    EXAMPLE CONFIGURATION     #
    #                              #
    ################################

    # This block is not actually functional configuration,
    # but serves to illustrate the available configuration
    # options and document them in a way that is accessible
    # to users that ` + "`" + `kubectl edit` + "`" + ` this config map.
    #
    # These sample configuration options may be copied out of
    # this example block and unindented to be in the data block
    # to actually change the configuration.

    # logging.enable-var-log-collection defaults to false.
    # The fluentd daemon set will be set up to collect /var/log if
    # this flag is true.
    logging.enable-var-log-collection: false

    # logging.revision-url-template provides a template to use for producing the
    # logging URL that is injected into the status of each Revision.
    # This value is what you might use the the Knative monitoring bundle, and provides
    # access to Kibana after setting up kubectl proxy.
    logging.revision-url-template: |
      http://localhost:8001/api/v1/namespaces/knative-monitoring/services/kibana-logging/proxy/app/kibana#/discover?_a=(query:(match:(kubernetes.labels.knative-dev%2FrevisionUID:(query:'${REVISION_UID}',type:phrase))))

    # If non-empty, this enables queue proxy writing request logs to stdout.
    # The value determines the shape of the request logs and it must be a valid go text/template.
    # It is important to keep this as a single line. Multiple lines are parsed as separate entities
    # by most collection agents and will split the request logs into multiple records.
    #
    # The following fields and functions are available to the template:
    #
    # Request: An http.Request (see https://golang.org/pkg/net/http/#Request)
    # representing an HTTP request received by the server.
    #
    # Response:
    # struct {
    #   Code    int       // HTTP status code (see https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml)
    #   Size    int       // An int representing the size of the response.
    #   Latency float64   // A float64 representing the latency of the response in seconds.
    # }
    #
    # Revision:
    # struct {
    #   Name          string  // Knative revision name
    #   Namespace     string  // Knative revision namespace
    #   Service       string  // Knative service name
    #   Configuration string  // Knative configuration name
    #   PodName       string  // Name of the pod hosting the revision
    #   PodIP         string  // IP of the pod hosting the revision
    # }
    #
    logging.request-log-template: '{"httpRequest": {"requestMethod": "{{.Request.Method}}", "requestUrl": "{{js .Request.RequestURI}}", "requestSize": "{{.Request.ContentLength}}", "status": {{.Response.Code}}, "responseSize": "{{.Response.Size}}", "userAgent": "{{js .Request.UserAgent}}", "remoteIp": "{{js .Request.RemoteAddr}}", "serverIp": "{{.Revision.PodIP}}", "referer": "{{js .Request.Referer}}", "latency": "{{.Response.Latency}}s", "protocol": "{{.Request.Proto}}"}, "traceId": "{{index .Request.Header "X-B3-Traceid"}}"}'

    # metrics.backend-destination field specifies the system metrics destination.
    # It supports either prometheus (the default) or stackdriver.
    # Note: Using stackdriver will incur additional charges
    metrics.backend-destination: prometheus

    # metrics.request-metrics-backend-destination specifies the request metrics
    # destination. If non-empty, it enables queue proxy to send request metrics.
    # Currently supported values: prometheus, stackdriver.
    metrics.request-metrics-backend-destination: prometheus

    # metrics.stackdriver-project-id field specifies the stackdriver project ID. This
    # field is optional. When running on GCE, application default credentials will be
    # used if this field is not provided.
    metrics.stackdriver-project-id: "<your stackdriver project id>"

    # metrics.allow-stackdriver-custom-metrics indicates whether it is allowed to send metrics to
    # Stackdriver using "global" resource type and custom metric type if the
    # metrics are not supported by "knative_revision" resource type. Setting this
    # flag to "true" could cause extra Stackdriver charge.
    # If metrics.backend-destination is not Stackdriver, this is ignored.
    metrics.allow-stackdriver-custom-metrics: "false"
kind: ConfigMap
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: config-observability
  namespace: knative-serving
---

apiVersion: v1
data:
  _example: |
    ################################
    #                              #
    #    EXAMPLE CONFIGURATION     #
    #                              #
    ################################

    # This block is not actually functional configuration,
    # but serves to illustrate the available configuration
    # options and document them in a way that is accessible
    # to users that ` + "`" + `kubectl edit` + "`" + ` this config map.
    #
    # These sample configuration options may be copied out of
    # this example block and unindented to be in the data block
    # to actually change the configuration.
    #
    # If true we enable adding spans within our applications.
    enable: "false"

    # URL to zipkin collector where traces are sent.
    zipkin-endpoint: "http://zipkin.istio-system.svc.cluster.local:9411/api/v2/spans"

    # Enable zipkin debug mode. This allows all spans to be sent to the server
    # bypassing sampling.
    debug: "false"

    # Percentage (0-1) of requests to trace
    sample-rate: "0.1"
kind: ConfigMap
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: config-tracing
  namespace: knative-serving
---

apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: controller
  namespace: knative-serving
spec:
  replicas: 1
  selector:
    matchLabels:
      app: controller
  template:
    metadata:
      annotations:
        sidecar.istio.io/inject: "false"
      labels:
        app: controller
        serving.knative.dev/release: "v0.7.0"
    spec:
      containers:
      - env:
        - name: SYSTEM_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: CONFIG_LOGGING_NAME
          value: config-logging
        - name: CONFIG_OBSERVABILITY_NAME
          value: config-observability
        - name: METRICS_DOMAIN
          value: knative.dev/serving
        image: gcr.io/knative-releases/github.com/knative/serving/cmd/controller@sha256:016c95f2d94be89683d1ddb7ea959667fd2d899087a4145a31d26b5d6f0bb38f
        name: controller
        ports:
        - containerPort: 9090
          name: metrics
        resources:
          limits:
            cpu: 1000m
            memory: 1000Mi
          requests:
            cpu: 100m
            memory: 100Mi
        securityContext:
          allowPrivilegeEscalation: false
        volumeMounts:
        - mountPath: /etc/config-logging
          name: config-logging
      serviceAccountName: controller
      volumes:
      - configMap:
          name: config-logging
        name: config-logging
---

apiVersion: apiregistration.k8s.io/v1beta1
kind: APIService
metadata:
  labels:
    autoscaling.knative.dev/metric-provider: custom-metrics
    serving.knative.dev/release: "v0.7.0"
  name: v1beta1.custom.metrics.k8s.io
spec:
  group: custom.metrics.k8s.io
  groupPriorityMinimum: 100
  insecureSkipTLSVerify: true
  service:
    name: autoscaler
    namespace: knative-serving
  version: v1beta1
  versionPriority: 100
---

apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    networking.knative.dev/certificate-provider: cert-manager
    serving.knative.dev/release: "v0.7.0"
  name: networking-certmanager
  namespace: knative-serving
spec:
  replicas: 1
  selector:
    matchLabels:
      app: networking-certmanager
  template:
    metadata:
      annotations:
        sidecar.istio.io/inject: "false"
      labels:
        app: networking-certmanager
    spec:
      containers:
      - env:
        - name: SYSTEM_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: CONFIG_LOGGING_NAME
          value: config-logging
        - name: CONFIG_OBSERVABILITY_NAME
          value: config-observability
        - name: METRICS_DOMAIN
          value: knative.dev/serving
        image: gcr.io/knative-releases/github.com/knative/serving/cmd/networking/certmanager@sha256:c757629165393f778d5c0e8b611c9c4857b24f0c748d985d3a080d0161a85248
        name: networking-certmanager
        ports:
        - containerPort: 9090
          name: metrics
        resources:
          limits:
            cpu: 1000m
            memory: 1000Mi
          requests:
            cpu: 100m
            memory: 100Mi
        securityContext:
          allowPrivilegeEscalation: false
        volumeMounts:
        - mountPath: /etc/config-logging
          name: config-logging
      serviceAccountName: controller
      volumes:
      - configMap:
          name: config-logging
        name: config-logging
---

apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: webhook
  namespace: knative-serving
spec:
  replicas: 1
  selector:
    matchLabels:
      app: webhook
      role: webhook
  template:
    metadata:
      annotations:
        cluster-autoscaler.kubernetes.io/safe-to-evict: "false"
        sidecar.istio.io/inject: "false"
      labels:
        app: webhook
        role: webhook
        serving.knative.dev/release: "v0.7.0"
    spec:
      containers:
      - env:
        - name: SYSTEM_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: CONFIG_LOGGING_NAME
          value: config-logging
        image: gcr.io/knative-releases/github.com/knative/serving/cmd/webhook@sha256:d9918d40492e0b20b48576ff6182e2ab896e50dfd2313cb471419be98f821b9c
        name: webhook
        resources:
          limits:
            cpu: 200m
            memory: 200Mi
          requests:
            cpu: 20m
            memory: 20Mi
        securityContext:
          allowPrivilegeEscalation: false
        volumeMounts:
        - mountPath: /etc/config-logging
          name: config-logging
      serviceAccountName: controller
      volumes:
      - configMap:
          name: config-logging
        name: config-logging
---
apiVersion: v1
kind: Namespace
metadata:
  name: knative-build
---

apiVersion: policy/v1beta1
kind: PodSecurityPolicy
metadata:
  name: knative-build
spec:
  allowPrivilegeEscalation: false
  fsGroup:
    ranges:
    - max: 65535
      min: 1
    rule: MustRunAs
  hostIPC: false
  hostNetwork: false
  hostPID: false
  privileged: false
  runAsUser:
    rule: RunAsAny
  seLinux:
    rule: RunAsAny
  supplementalGroups:
    ranges:
    - max: 65535
      min: 1
    rule: MustRunAs
  volumes:
  - configMap
  - secret
---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: knative-build-admin
rules:
- apiGroups:
  - ""
  resources:
  - pods
  - namespaces
  - secrets
  - events
  - serviceaccounts
  - configmaps
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
- apiGroups:
  - apps
  resources:
  - deployments
  - deployments/finalizers
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
- apiGroups:
  - admissionregistration.k8s.io
  resources:
  - mutatingwebhookconfigurations
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
- apiGroups:
  - build.knative.dev
  resources:
  - builds
  - buildtemplates
  - clusterbuildtemplates
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
- apiGroups:
  - build.knative.dev
  resources:
  - builds/status
  - buildtemplates/status
  - clusterbuildtemplates/status
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
- apiGroups:
  - caching.internal.knative.dev
  resources:
  - images
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - deletecollection
  - patch
  - watch
- apiGroups:
  - policy
  resourceNames:
  - knative-build
  resources:
  - podsecuritypolicies
  verbs:
  - use
---

apiVersion: v1
kind: ServiceAccount
metadata:
  name: build-controller
  namespace: knative-build
---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: build-controller-admin
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: knative-build-admin
subjects:
- kind: ServiceAccount
  name: build-controller
  namespace: knative-build
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    knative.dev/crd-install: "true"
  name: builds.build.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .status.conditions[?(@.type=="Succeeded")].status
    name: Succeeded
    type: string
  - JSONPath: .status.conditions[?(@.type=="Succeeded")].reason
    name: Reason
    type: string
  - JSONPath: .status.startTime
    name: StartTime
    type: date
  - JSONPath: .status.completionTime
    name: CompletionTime
    type: date
  group: build.knative.dev
  names:
    categories:
    - all
    - knative
    kind: Build
    plural: builds
  scope: Namespaced
  subresources:
    status: {}
  version: v1alpha1
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    knative.dev/crd-install: "true"
  name: buildtemplates.build.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .metadata.creationTimestamp
    name: Age
    type: date
  group: build.knative.dev
  names:
    categories:
    - all
    - knative
    kind: BuildTemplate
    plural: buildtemplates
  scope: Namespaced
  subresources:
    status: {}
  version: v1alpha1
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    knative.dev/crd-install: "true"
  name: clusterbuildtemplates.build.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .metadata.creationTimestamp
    name: Age
    type: date
  group: build.knative.dev
  names:
    categories:
    - all
    - knative
    kind: ClusterBuildTemplate
    plural: clusterbuildtemplates
  scope: Cluster
  subresources:
    status: {}
  version: v1alpha1
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    knative.dev/crd-install: "true"
  name: images.caching.internal.knative.dev
spec:
  group: caching.internal.knative.dev
  names:
    categories:
    - all
    - knative-internal
    - caching
    kind: Image
    plural: images
    shortNames:
    - img
    singular: image
  scope: Namespaced
  subresources:
    status: {}
  version: v1alpha1
---

apiVersion: v1
kind: Service
metadata:
  labels:
    app: build-controller
  name: build-controller
  namespace: knative-build
spec:
  ports:
  - name: metrics
    port: 9090
    protocol: TCP
    targetPort: 9090
  selector:
    app: build-controller
---

apiVersion: v1
kind: Service
metadata:
  labels:
    role: build-webhook
  name: build-webhook
  namespace: knative-build
spec:
  ports:
  - port: 443
    targetPort: 8443
  selector:
    role: build-webhook
---

apiVersion: caching.internal.knative.dev/v1alpha1
kind: Image
metadata:
  name: creds-init
  namespace: knative-build
spec:
  image: gcr.io/knative-releases/github.com/knative/build/cmd/creds-init@sha256:1a984c032a2606f8491f4a19a85209dcc1ae2cfd494c3dafe8a74269379ad2c8---

apiVersion: caching.internal.knative.dev/v1alpha1
kind: Image
metadata:
  name: git-init
  namespace: knative-build
spec:
  image: gcr.io/knative-releases/github.com/knative/build/cmd/git-init@sha256:06505d8c621e9337d0dd1bc13ed4545a33e857fbb6374740cc6337d2ba55796d---

apiVersion: caching.internal.knative.dev/v1alpha1
kind: Image
metadata:
  name: gcs-fetcher
  namespace: knative-build
spec:
  image: gcr.io/cloud-builders/gcs-fetcher---

apiVersion: caching.internal.knative.dev/v1alpha1
kind: Image
metadata:
  name: nop
  namespace: knative-build
spec:
  image: gcr.io/knative-releases/github.com/knative/build/cmd/nop@sha256:8aca9c97ede9a550ac3536d00c5d7acaae5e3a4fe514f4329ec261d935eddabb
---

apiVersion: v1
data:
  loglevel.controller: info
  loglevel.creds-init: info
  loglevel.git-init: info
  loglevel.webhook: info
  zap-logger-config: |
    {
      "level": "info",
      "development": false,
      "sampling": {
        "initial": 100,
        "thereafter": 100
      },
      "outputPaths": ["stdout"],
      "errorOutputPaths": ["stderr"],
      "encoding": "json",
      "encoderConfig": {
        "timeKey": "",
        "levelKey": "level",
        "nameKey": "logger",
        "callerKey": "caller",
        "messageKey": "msg",
        "stacktraceKey": "stacktrace",
        "lineEnding": "",
        "levelEncoder": "",
        "timeEncoder": "",
        "durationEncoder": "",
        "callerEncoder": ""
      }
    }
kind: ConfigMap
metadata:
  name: config-logging
  namespace: knative-build
---

apiVersion: v1
data:
  _example: |
    ################################
    #                              #
    #    EXAMPLE CONFIGURATION     #
    #                              #
    ################################

    # This block is not actually functional configuration,
    # but serves to illustrate the available configuration
    # options and document them in a way that is accessible
    # to users that ` + "`" + `kubectl edit` + "`" + ` this config map.
    #
    # These sample configuration options may be copied out of
    # this example block and unindented to be in the data block
    # to actually change the configuration.

    # metrics.backend-destination field specifies the system metrics destination.
    # It supports either prometheus (the default) or stackdriver.
    # Note: Using stackdriver will incur additional charges
    metrics.backend-destination: prometheus

    # metrics.stackdriver-project-id field specifies the stackdriver project ID. This
    # field is optional. When running on GCE, application default credentials will be
    # used if this field is not provided.
    metrics.stackdriver-project-id: "<your stackdriver project id>"

    # metrics.allow-stackdriver-custom-metrics indicates whether it is allowed to send metrics to
    # Stackdriver using "global" resource type and custom metric type if the
    # metrics are not supported by "knative_revision" resource type. Setting this
    # flag to "true" could cause extra Stackdriver charge.
    # If metrics.backend-destination is not Stackdriver, this is ignored.
    metrics.allow-stackdriver-custom-metrics: "false"
kind: ConfigMap
metadata:
  name: config-observability
  namespace: knative-build
---

apiVersion: apps/v1
kind: Deployment
metadata:
  name: build-controller
  namespace: knative-build
spec:
  replicas: 1
  selector:
    matchLabels:
      app: build-controller
  template:
    metadata:
      labels:
        app: build-controller
    spec:
      containers:
      - args:
        - -logtostderr
        - -stderrthreshold
        - INFO
        - -creds-image
        - gcr.io/knative-releases/github.com/knative/build/cmd/creds-init@sha256:1a984c032a2606f8491f4a19a85209dcc1ae2cfd494c3dafe8a74269379ad2c8
        - -git-image
        - gcr.io/knative-releases/github.com/knative/build/cmd/git-init@sha256:06505d8c621e9337d0dd1bc13ed4545a33e857fbb6374740cc6337d2ba55796d
        - -nop-image
        - gcr.io/knative-releases/github.com/knative/build/cmd/nop@sha256:8aca9c97ede9a550ac3536d00c5d7acaae5e3a4fe514f4329ec261d935eddabb
        env:
        - name: SYSTEM_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: CONFIG_LOGGING_NAME
          value: config-logging
        - name: CONFIG_OBSERVABILITY_NAME
          value: config-observability
        - name: METRICS_DOMAIN
          value: knative.dev/build
        image: gcr.io/knative-releases/github.com/knative/build/cmd/controller@sha256:5adb5ba0647a7b1af1d90848bf72a75fa84efeb89e1d688465a2105c1cce1dc2
        name: build-controller
        ports:
        - containerPort: 9090
          name: metrics
        resources:
          limits:
            cpu: 1000m
            memory: 1000Mi
          requests:
            cpu: 100m
            memory: 100Mi
        volumeMounts:
        - mountPath: /etc/config-logging
          name: config-logging
      serviceAccountName: build-controller
      volumes:
      - configMap:
          name: config-logging
        name: config-logging
---

apiVersion: apps/v1
kind: Deployment
metadata:
  name: build-webhook
  namespace: knative-build
spec:
  replicas: 1
  selector:
    matchLabels:
      app: build-webhook
  template:
    metadata:
      labels:
        app: build-webhook
        role: build-webhook
    spec:
      containers:
      - args:
        - -logtostderr
        - -stderrthreshold
        - INFO
        env:
        - name: SYSTEM_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        image: gcr.io/knative-releases/github.com/knative/build/cmd/webhook@sha256:35b1b5f72642e9c1ee71809fec309a019111beebf805f9ddddf154a97ad23975
        name: build-webhook
        resources:
          limits:
            memory: 1000Mi
        volumeMounts:
        - mountPath: /etc/config-logging
          name: config-logging
      serviceAccountName: build-controller
      volumes:
      - configMap:
          name: config-logging
        name: config-logging
---

# eventing.yaml
apiVersion: v1
kind: Namespace
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: knative-eventing
---

aggregationRule:
  clusterRoleSelectors:
  - matchLabels:
      duck.knative.dev/addressable: "true"
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: addressable-resolver
rules: []---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    duck.knative.dev/addressable: "true"
    eventing.knative.dev/release: v0.7.0
  name: service-addressable-resolver
rules:
- apiGroups:
  - ""
  resources:
  - services
  verbs:
  - get
  - list
  - watch---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    duck.knative.dev/addressable: "true"
    eventing.knative.dev/release: v0.7.0
  name: serving-addressable-resolver
rules:
- apiGroups:
  - serving.knative.dev
  resources:
  - routes
  - routes/status
  - services
  - services/status
  verbs:
  - get
  - list
  - watch---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    duck.knative.dev/addressable: "true"
    eventing.knative.dev/release: v0.7.0
  name: channel-addressable-resolver
rules:
- apiGroups:
  - eventing.knative.dev
  resources:
  - channels
  - channels/status
  verbs:
  - get
  - list
  - watch---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    duck.knative.dev/addressable: "true"
    eventing.knative.dev/release: v0.7.0
  name: broker-addressable-resolver
rules:
- apiGroups:
  - eventing.knative.dev
  resources:
  - brokers
  - brokers/status
  verbs:
  - get
  - list
  - watch
---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: eventing-broker-filter
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - eventing.knative.dev
  resources:
  - triggers
  - triggers/status
  verbs:
  - get
  - list
  - watch---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: eventing-broker-ingress
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: eventing-config-reader
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
---

aggregationRule:
  clusterRoleSelectors:
  - matchLabels:
      duck.knative.dev/channelable: "true"
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: channelable-manipulator
rules: []
---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: knative-eventing-controller
rules:
- apiGroups:
  - ""
  resources:
  - namespaces
  - secrets
  - configmaps
  - services
  - events
  - serviceaccounts
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
- apiGroups:
  - apps
  resources:
  - deployments
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
- apiGroups:
  - rbac.authorization.k8s.io
  resources:
  - rolebindings
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
- apiGroups:
  - eventing.knative.dev
  resources:
  - brokers
  - brokers/status
  - channels
  - channels/status
  - clusterchannelprovisioners
  - clusterchannelprovisioners/status
  - subscriptions
  - subscriptions/status
  - triggers
  - triggers/status
  - eventtypes
  - eventtypes/status
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
- apiGroups:
  - eventing.knative.dev
  resources:
  - brokers/finalizers
  - triggers/finalizers
  verbs:
  - update
- apiGroups:
  - messaging.knative.dev
  resources:
  - sequences
  - sequences/status
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
- apiGroups:
  - sources.eventing.knative.dev
  resources:
  - cronjobsources
  - cronjobsources/status
  - cronjobsources/finalizers
  - containersources
  - containersources/status
  - containersources/finalizers
  - apiserversources
  - apiserversources/status
  - apiserversources/finalizers
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions
  verbs:
  - get
  - list
  - watch
---

apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: eventing-controller
  namespace: knative-eventing---

apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: eventing-webhook
  namespace: knative-eventing
---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: knative-eventing-webhook
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - get
  - create
- apiGroups:
  - apps
  resources:
  - deployments
  verbs:
  - get
- apiGroups:
  - apps
  resources:
  - deployments/finalizers
  verbs:
  - update
- apiGroups:
  - admissionregistration.k8s.io
  resources:
  - mutatingwebhookconfigurations
  verbs:
  - get
  - list
  - create
  - update
  - delete
  - patch
  - watch
- apiGroups:
  - eventing.knative.dev
  resources:
  - brokers
  - brokers/status
  - channels
  - channels/status
  - clusterchannelprovisioners
  - clusterchannelprovisioners/status
  - subscriptions
  - subscriptions/status
  - triggers
  - triggers/status
  - eventtypes
  - eventtypes/status
  verbs:
  - get
  - list
  - watch
---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: eventing-controller
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: knative-eventing-controller
subjects:
- kind: ServiceAccount
  name: eventing-controller
  namespace: knative-eventing---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: eventing-controller-resolver
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: addressable-resolver
subjects:
- kind: ServiceAccount
  name: eventing-controller
  namespace: knative-eventing---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: eventing-controller-manipulator
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: channelable-manipulator
subjects:
- kind: ServiceAccount
  name: eventing-controller
  namespace: knative-eventing---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: eventing-webhook
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: knative-eventing-webhook
subjects:
- kind: ServiceAccount
  name: eventing-webhook
  namespace: knative-eventing
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  creationTimestamp: null
  labels:
    eventing.knative.dev/release: v0.7.0
    eventing.knative.dev/source: "true"
    knative.dev/crd-install: "true"
  name: apiserversources.sources.eventing.knative.dev
spec:
  group: sources.eventing.knative.dev
  names:
    categories:
    - all
    - knative
    - eventing
    - sources
    kind: ApiServerSource
    plural: apiserversources
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        apiVersion:
          type: string
        kind:
          type: string
        metadata:
          type: object
        spec:
          properties:
            mode:
              description: 'Mode controls the content of the event payload. One of:
                ''Ref'' (only references of resources), ''Resource'' (full resource).'
              type: string
            resources:
              items:
                properties:
                  apiVersion:
                    description: API version of the objects to watch.
                    type: string
                  controller:
                    description: 'If true, emits the managing controller ref. Only
                      supported for mode=Ref. More info: https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/'
                    type: boolean
                  kind:
                    description: Kind of the objects to watch.
                    type: string
              type: array
            serviceAccountName:
              description: 'name of the ServiceAccount to use to run the receive adapter.
                More info: https://kubernetes.io/docs/reference/access-authn-authz/service-accounts-admin/.'
              type: string
            sink:
              description: A reference to the object that should receive events.
              type: object
          required:
          - resources
          - sink
          type: object
        status:
          properties:
            conditions:
              items:
                properties:
                  lastTransitionTime:
                    type: string
                  message:
                    type: string
                  reason:
                    type: string
                  severity:
                    type: string
                  status:
                    type: string
                  type:
                    type: string
                required:
                - type
                - status
                type: object
              type: array
            sinkUri:
              type: string
          type: object
  version: v1alpha1
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
    knative.dev/crd-install: "true"
  name: brokers.eventing.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .status.conditions[?(@.type=="Ready")].status
    name: Ready
    type: string
  - JSONPath: .status.conditions[?(@.type=="Ready")].reason
    name: Reason
    type: string
  - JSONPath: .status.address.hostname
    name: Hostname
    type: string
  - JSONPath: .metadata.creationTimestamp
    name: Age
    type: date
  group: eventing.knative.dev
  names:
    categories:
    - all
    - knative
    - eventing
    kind: Broker
    plural: brokers
    singular: broker
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          properties:
            channelTemplate:
              properties:
                arguments:
                  type: object
                provisioner:
                  properties:
                    apiVersion:
                      minLength: 1
                      type: string
                    kind:
                      minLength: 1
                      type: string
                    name:
                      minLength: 1
                      type: string
                  required:
                  - apiVersion
                  - kind
                  - name
                  type: object
              required:
              - provisioner
              type: object
  version: v1alpha1
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
    knative.dev/crd-install: "true"
  name: channels.eventing.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .status.conditions[?(@.type=="Ready")].status
    name: Ready
    type: string
  - JSONPath: .status.conditions[?(@.type=="Ready")].reason
    name: Reason
    type: string
  - JSONPath: .metadata.creationTimestamp
    name: Age
    type: date
  group: eventing.knative.dev
  names:
    categories:
    - all
    - knative
    - eventing
    kind: Channel
    plural: channels
    shortNames:
    - chan
    singular: channel
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          properties:
            arguments:
              type: object
            provisioner:
              properties:
                apiVersion:
                  minLength: 1
                  type: string
                kind:
                  minLength: 1
                  type: string
                name:
                  minLength: 1
                  type: string
              required:
              - apiVersion
              - kind
              - name
              type: object
            subscribable:
              properties:
                subscribers:
                  items:
                    properties:
                      generation:
                        type: integer
                      ref:
                        properties:
                          apiVersion:
                            type: string
                          kind:
                            type: string
                          name:
                            minLength: 1
                            type: string
                          namespace:
                            minLength: 1
                            type: string
                          uid:
                            minLength: 1
                            type: string
                        required:
                        - namespace
                        - name
                        - uid
                        type: object
                      replyURI:
                        minLength: 1
                        type: string
                      subscriberURI:
                        minLength: 1
                        type: string
                      uid:
                        minLength: 1
                        type: string
                    required:
                    - uid
                  type: array
              type: object
  version: v1alpha1
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
    knative.dev/crd-install: "true"
  name: clusterchannelprovisioners.eventing.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .status.conditions[?(@.type=="Ready")].status
    name: Ready
    type: string
  - JSONPath: .status.conditions[?(@.type=="Ready")].reason
    name: Reason
    type: string
  - JSONPath: .metadata.creationTimestamp
    name: Age
    type: date
  group: eventing.knative.dev
  names:
    categories:
    - all
    - knative
    - eventing
    - provisioner
    kind: ClusterChannelProvisioner
    plural: clusterchannelprovisioners
    shortNames:
    - ccp
    singular: clusterchannelprovisioner
  scope: Cluster
  subresources:
    status: {}
  version: v1alpha1
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
    eventing.knative.dev/source: "true"
    knative.dev/crd-install: "true"
  name: containersources.sources.eventing.knative.dev
spec:
  group: sources.eventing.knative.dev
  names:
    categories:
    - all
    - knative
    - eventing
    - sources
    kind: ContainerSource
    plural: containersources
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        apiVersion:
          type: string
        kind:
          type: string
        metadata:
          type: object
        spec:
          properties:
            args:
              items:
                type: string
              type: array
            env:
              items:
                type: object
              type: array
            image:
              minLength: 1
              type: string
            serviceAccountName:
              type: string
            sink:
              type: object
            template:
              type: object
          type: object
        status:
          properties:
            conditions:
              items:
                properties:
                  lastTransitionTime:
                    type: string
                  message:
                    type: string
                  reason:
                    type: string
                  severity:
                    type: string
                  status:
                    type: string
                  type:
                    type: string
                required:
                - type
                - status
                type: object
              type: array
            sinkUri:
              type: string
          type: object
  version: v1alpha1
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
    eventing.knative.dev/source: "true"
    knative.dev/crd-install: "true"
  name: cronjobsources.sources.eventing.knative.dev
spec:
  group: sources.eventing.knative.dev
  names:
    categories:
    - all
    - knative
    - eventing
    - sources
    kind: CronJobSource
    plural: cronjobsources
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        apiVersion:
          type: string
        kind:
          type: string
        metadata:
          type: object
        spec:
          properties:
            data:
              type: string
            resources:
              properties:
                limits:
                  properties:
                    cpu:
                      type: string
                    memory:
                      type: string
                  type: object
                requests:
                  properties:
                    cpu:
                      type: string
                    memory:
                      type: string
                  type: object
              type: object
            schedule:
              type: string
            serviceAccountName:
              type: string
            sink:
              type: object
          required:
          - schedule
          type: object
        status:
          properties:
            conditions:
              items:
                properties:
                  lastTransitionTime:
                    type: string
                  message:
                    type: string
                  reason:
                    type: string
                  severity:
                    type: string
                  status:
                    type: string
                  type:
                    type: string
                required:
                - type
                - status
                type: object
              type: array
            sinkUri:
              type: string
          type: object
  version: v1alpha1
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
    knative.dev/crd-install: "true"
  name: eventtypes.eventing.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .spec.type
    name: Type
    type: string
  - JSONPath: .spec.source
    name: Source
    type: string
  - JSONPath: .spec.schema
    name: Schema
    type: string
  - JSONPath: .spec.broker
    name: Broker
    type: string
  - JSONPath: .spec.description
    name: Description
    type: string
  - JSONPath: .status.conditions[?(@.type=="Ready")].status
    name: Ready
    type: string
  - JSONPath: .status.conditions[?(@.type=="Ready")].reason
    name: Reason
    type: string
  group: eventing.knative.dev
  names:
    categories:
    - all
    - knative
    - eventing
    kind: EventType
    plural: eventtypes
    singular: eventtype
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          properties:
            broker:
              minLength: 1
              type: string
            description:
              type: string
            schema:
              type: string
            source:
              minLength: 1
              type: string
            type:
              minLength: 1
              type: string
          required:
          - type
          - source
          - broker
  version: v1alpha1
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
    knative.dev/crd-install: "true"
  name: sequences.messaging.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .status.conditions[?(@.type=="Ready")].status
    name: Ready
    type: string
  - JSONPath: .status.conditions[?(@.type=="Ready")].reason
    name: Reason
    type: string
  - JSONPath: .status.address.hostname
    name: Hostname
    type: string
  - JSONPath: .metadata.creationTimestamp
    name: Age
    type: date
  group: messaging.knative.dev
  names:
    categories:
    - all
    - knative
    - eventing
    - messaging
    kind: Sequence
    plural: sequences
    singular: sequence
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          properties:
            channelTemplate:
              properties:
                apiVersion:
                  minLength: 1
                  type: string
                kind:
                  minLength: 1
                  type: string
                spec:
                  type: object
              required:
              - apiVersion
              - kind
              type: object
            steps:
              items:
                properties:
                  dnsName:
                    minLength: 1
                    type: string
                  ref:
                    properties:
                      apiVersion:
                        minLength: 1
                        type: string
                      kind:
                        minLength: 1
                        type: string
                      name:
                        minLength: 1
                        type: string
                      namespace:
                        maxLength: 0
                        type: string
                    required:
                    - apiVersion
                    - kind
                    - name
                    type: object
                  uri:
                    minLength: 1
                    type: string
                type: object
              type: array
          required:
          - steps
          - channelTemplate
  version: v1alpha1
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
    knative.dev/crd-install: "true"
  name: subscriptions.eventing.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .status.conditions[?(@.type=="Ready")].status
    name: Ready
    type: string
  - JSONPath: .status.conditions[?(@.type=="Ready")].reason
    name: Reason
    type: string
  - JSONPath: .metadata.creationTimestamp
    name: Age
    type: date
  group: eventing.knative.dev
  names:
    categories:
    - all
    - knative
    - eventing
    kind: Subscription
    plural: subscriptions
    shortNames:
    - sub
    singular: subscription
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          properties:
            channel:
              properties:
                apiVersion:
                  minLength: 1
                  type: string
                kind:
                  type: string
                name:
                  minLength: 1
                  type: string
              required:
              - apiVersion
              - kind
              - name
              type: object
            reply:
              properties:
                channel:
                  properties:
                    apiVersion:
                      minLength: 1
                      type: string
                    kind:
                      type: string
                    name:
                      minLength: 1
                      type: string
                  required:
                  - apiVersion
                  - kind
                  - name
                  type: object
              type: object
            subscriber:
              properties:
                dnsName:
                  minLength: 1
                  type: string
                ref:
                  properties:
                    apiVersion:
                      minLength: 1
                      type: string
                    kind:
                      minLength: 1
                      type: string
                    name:
                      minLength: 1
                      type: string
                  required:
                  - apiVersion
                  - kind
                  - name
                  type: object
                uri:
                  minLength: 1
                  type: string
              type: object
          required:
          - channel
  version: v1alpha1
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
    knative.dev/crd-install: "true"
  name: triggers.eventing.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .status.conditions[?(@.type=="Ready")].status
    name: Ready
    type: string
  - JSONPath: .status.conditions[?(@.type=="Ready")].reason
    name: Reason
    type: string
  - JSONPath: .spec.broker
    name: Broker
    type: string
  - JSONPath: .status.subscriberURI
    name: Subscriber_URI
    type: string
  - JSONPath: .metadata.creationTimestamp
    name: Age
    type: date
  group: eventing.knative.dev
  names:
    categories:
    - all
    - knative
    - eventing
    kind: Trigger
    plural: triggers
    singular: trigger
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          properties:
            broker:
              type: string
            filter:
              properties:
                sourceAndType:
                  properties:
                    source:
                      type: string
                    type:
                      type: string
                  type: object
              type: object
            subscriber:
              properties:
                dnsName:
                  minLength: 1
                  type: string
                ref:
                  properties:
                    apiVersion:
                      minLength: 1
                      type: string
                    kind:
                      minLength: 1
                      type: string
                    name:
                      minLength: 1
                      type: string
                  required:
                  - apiVersion
                  - kind
                  - name
                  type: object
                uri:
                  minLength: 1
                  type: string
              type: object
          required:
          - subscriber
  version: v1alpha1
---

apiVersion: v1
data:
  default-channel-config: |
    clusterdefault:
      apiversion: eventing.knative.dev/v1alpha1
      kind: ClusterChannelProvisioner
      name: in-memory
    namespacedefaults:
      some-namespace:
        apiversion: eventing.knative.dev/v1alpha1
        kind: ClusterChannelProvisioner
        name: some-other-provisioner
kind: ConfigMap
metadata:
  name: default-channel-webhook
  namespace: knative-eventing
---

apiVersion: v1
kind: Service
metadata:
  labels:
    app: sources-controller
    eventing.knative.dev/release: v0.7.0
  name: sources-controller
  namespace: knative-eventing
spec:
  ports:
  - name: metrics
    port: 9090
    protocol: TCP
    targetPort: 9090
  selector:
    app: sources-controller
---

apiVersion: v1
kind: Service
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
    role: eventing-webhook
  name: eventing-webhook
  namespace: knative-eventing
spec:
  ports:
  - port: 443
    targetPort: 8443
  selector:
    role: eventing-webhook
---

apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: eventing-controller
  namespace: knative-eventing
spec:
  replicas: 1
  selector:
    matchLabels:
      app: eventing-controller
  template:
    metadata:
      annotations:
        sidecar.istio.io/inject: "false"
      labels:
        app: eventing-controller
        eventing.knative.dev/release: v0.7.0
    spec:
      containers:
      - env:
        - name: SYSTEM_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: CONFIG_LOGGING_NAME
          value: config-logging
        - name: CONFIG_OBSERVABILITY_NAME
          value: config-observability
        - name: METRICS_DOMAIN
          value: knative.dev/eventing
        - name: BROKER_INGRESS_IMAGE
          value: gcr.io/knative-releases/github.com/knative/eventing/cmd/broker/ingress@sha256:33d41dc38208bf9752b38e5c90149da4d5af74a67f317b7da8cb3c458fbd0fff
        - name: BROKER_INGRESS_SERVICE_ACCOUNT
          value: eventing-broker-ingress
        - name: BROKER_FILTER_IMAGE
          value: gcr.io/knative-releases/github.com/knative/eventing/cmd/broker/filter@sha256:5a4eb60a605e189516a36a01c7fd39001d5766b2e6bb80c69744e15515282360
        - name: BROKER_FILTER_SERVICE_ACCOUNT
          value: eventing-broker-filter
        image: gcr.io/knative-releases/github.com/knative/eventing/cmd/controller@sha256:57f273774efb017bbf06729af802514db2f3ab070b51730dba9330903aa34163
        name: eventing-controller
        ports:
        - containerPort: 9090
          name: metrics
        terminationMessagePolicy: FallbackToLogsOnError
        volumeMounts:
        - mountPath: /etc/config-logging
          name: config-logging
      serviceAccountName: eventing-controller
      volumes:
      - configMap:
          name: config-logging
        name: config-logging
---

apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: sources-controller
  namespace: knative-eventing
spec:
  replicas: 1
  selector:
    matchLabels:
      app: sources-controller
  template:
    metadata:
      annotations:
        sidecar.istio.io/inject: "false"
      labels:
        app: sources-controller
        eventing.knative.dev/release: v0.7.0
    spec:
      containers:
      - env:
        - name: SYSTEM_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: CONFIG_LOGGING_NAME
          value: config-logging
        - name: CONFIG_OBSERVABILITY_NAME
          value: config-observability
        - name: METRICS_DOMAIN
          value: knative.dev/sources
        - name: CRONJOB_RA_IMAGE
          value: gcr.io/knative-releases/github.com/knative/eventing/cmd/cronjob_receive_adapter@sha256:634fbf0348f9f10d09c8110c505173aed91ce747bb2b87605e6e1bb10dce270b
        - name: APISERVER_RA_IMAGE
          value: gcr.io/knative-releases/github.com/knative/eventing/cmd/apiserver_receive_adapter@sha256:f18cdbc3c3077ece8505a4f4c49055e6c1c577e9fa42446f0f81193e48aa1d60
        image: gcr.io/knative-releases/github.com/knative/eventing/cmd/sources_controller@sha256:31ec2b4a1d1d9b81cd1eed6632b8a1c540b510b584e58c14ebc1c000330fe32c
        name: controller
        ports:
        - containerPort: 9090
          name: metrics
        resources:
          requests:
            cpu: 100m
            memory: 100Mi
        volumeMounts:
        - mountPath: /etc/config-logging
          name: config-logging
      serviceAccountName: eventing-controller
      volumes:
      - configMap:
          name: config-logging
        name: config-logging
---

apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: eventing-webhook
  namespace: knative-eventing
spec:
  replicas: 1
  selector:
    matchLabels:
      app: eventing-webhook
      role: eventing-webhook
  template:
    metadata:
      annotations:
        sidecar.istio.io/inject: "false"
      labels:
        app: eventing-webhook
        role: eventing-webhook
    spec:
      containers:
      - env:
        - name: SYSTEM_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: CONFIG_LOGGING_NAME
          value: config-logging
        - name: WEBHOOK_NAME
          value: eventing-webhook
        image: gcr.io/knative-releases/github.com/knative/eventing/cmd/webhook@sha256:3b5de8074f00469c393910fd0fbac70cec10838a858c94ad755af1b6bd6712fd
        name: eventing-webhook
        terminationMessagePolicy: FallbackToLogsOnError
        volumeMounts:
        - mountPath: /etc/config-logging
          name: config-logging
      serviceAccountName: eventing-webhook
      volumes:
      - configMap:
          name: config-logging
        name: config-logging---

apiVersion: apps/v1
kind: Deployment
metadata:
  name: webhook
  namespace: knative-eventing
spec:
  replicas: 0
  selector:
    matchLabels:
      app: webhook
      role: webhook
  template:
    metadata:
      annotations:
        sidecar.istio.io/inject: "false"
      labels:
        app: webhook
        role: webhook
    spec:
      containers:
      - env:
        - name: SYSTEM_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: CONFIG_LOGGING_NAME
          value: config-logging
        - name: WEBHOOK_NAME
          value: webhook
        image: gcr.io/knative-releases/github.com/knative/eventing/cmd/webhook@sha256:3b5de8074f00469c393910fd0fbac70cec10838a858c94ad755af1b6bd6712fd
        name: webhook
        resources:
          limits:
            memory: 1000Mi
        terminationMessagePolicy: FallbackToLogsOnError
        volumeMounts:
        - mountPath: /etc/config-logging
          name: config-logging
      serviceAccountName: eventing-webhook
      volumes:
      - configMap:
          name: config-logging
        name: config-logging
---

apiVersion: v1
data:
  loglevel.controller: info
  loglevel.webhook: info
  zap-logger-config: |
    {
      "level": "info",
      "development": false,
      "outputPaths": ["stdout"],
      "errorOutputPaths": ["stderr"],
      "encoding": "json",
      "encoderConfig": {
        "timeKey": "ts",
        "levelKey": "level",
        "nameKey": "logger",
        "callerKey": "caller",
        "messageKey": "msg",
        "stacktraceKey": "stacktrace",
        "lineEnding": "",
        "levelEncoder": "",
        "timeEncoder": "iso8601",
        "durationEncoder": "",
        "callerEncoder": ""
      }
    }
kind: ConfigMap
metadata:
  name: config-logging
  namespace: knative-eventing
---

apiVersion: v1
data:
  _example: |
    ################################
    #                              #
    #    EXAMPLE CONFIGURATION     #
    #                              #
    ################################

    # This block is not actually functional configuration,
    # but serves to illustrate the available configuration
    # options and document them in a way that is accessible
    # to users that ` + "`" + `kubectl edit` + "`" + ` this config map.
    #
    # These sample configuration options may be copied out of
    # this example block and unindented to be in the data block
    # to actually change the configuration.

    # logging.enable-var-log-collection defaults to false.
    # A fluentd sidecar will be set up to collect var log if
    # this flag is true.
    logging.enable-var-log-collection: false

    # logging.fluentd-sidecar-image provides the fluentd sidecar image
    # to inject as a sidecar to collect logs from /var/log.
    # Must be presented if logging.enable-var-log-collection is true.
    logging.fluentd-sidecar-image: k8s.gcr.io/fluentd-elasticsearch:v2.0.4

    # logging.fluentd-sidecar-output-config provides the configuration
    # for the fluentd sidecar, which will be placed into a configmap and
    # mounted into the fluentd sidecar image.
    logging.fluentd-sidecar-output-config: |
      # Parse json log before sending to Elastic Search
      <filter **>
        @type parser
        key_name log
        <parse>
          @type multi_format
          <pattern>
            format json
            time_key fluentd-time # fluentd-time is reserved for structured logs
            time_format %Y-%m-%dT%H:%M:%S.%NZ
          </pattern>
          <pattern>
            format none
            message_key log
          </pattern>
        </parse>
      </filter>
      # Send to Elastic Search
      <match **>
        @id elasticsearch
        @type elasticsearch
        @log_level info
        include_tag_key true
        # Elasticsearch service is in monitoring namespace.
        host elasticsearch-logging.knative-monitoring
        port 9200
        logstash_format true
        <buffer>
          @type file
          path /var/log/fluentd-buffers/kubernetes.system.buffer
          flush_mode interval
          retry_type exponential_backoff
          flush_thread_count 2
          flush_interval 5s
          retry_forever
          retry_max_interval 30
          chunk_limit_size 2M
          queue_limit_length 8
          overflow_action block
        </buffer>
      </match>

    # logging.revision-url-template provides a template to use for producing the
    # logging URL that is injected into the status of each Revision.
    # This value is what you might use the the Knative monitoring bundle, and provides
    # access to Kibana after setting up kubectl proxy.
    logging.revision-url-template: |
      http://localhost:8001/api/v1/namespaces/knative-monitoring/services/kibana-logging/proxy/app/kibana#/discover?_a=(query:(match:(kubernetes.labels.knative-dev%2FrevisionUID:(query:'${REVISION_UID}',type:phrase))))

    # If non-empty, this enables queue proxy writing request logs to stdout.
    # The value determines the shape of the request logs and it must be a valid go text/template.
    # It is important to keep this as a single line. Multiple lines are parsed as separate entities
    # by most collection agents and will split the request logs into multiple records.
    #
    # The following fields and functions are available to the template:
    #
    # Request: An http.Request (see https://golang.org/pkg/net/http/#Request)
    # representing an HTTP request received by the server.
    #
    # Response:
    # struct {
    #   Code    int       // HTTP status code (see https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml)
    #   Size    int       // An int representing the size of the response.
    #   Latency float64   // A float64 representing the latency of the response in seconds.
    # }
    #
    # Revision:
    # struct {
    #   Name          string  // Knative revision name
    #   Namespace     string  // Knative revision namespace
    #   Service       string  // Knative service name
    #   Configuration string  // Knative configuration name
    #   PodName       string  // Name of the pod hosting the revision
    #   PodIP         string  // IP of the pod hosting the revision
    # }
    #
    logging.request-log-template: '{"httpRequest": {"requestMethod": "{{.Request.Method}}", "requestUrl": "{{js .Request.RequestURI}}", "requestSize": "{{.Request.ContentLength}}", "status": {{.Response.Code}}, "responseSize": "{{.Response.Size}}", "userAgent": "{{js .Request.UserAgent}}", "remoteIp": "{{js .Request.RemoteAddr}}", "serverIp": "{{.Revision.PodIP}}", "referer": "{{js .Request.Referer}}", "latency": "{{.Response.Latency}}s", "protocol": "{{.Request.Proto}}"}, "traceId": "{{index .Request.Header "X-B3-Traceid"}}"}'

    # metrics.backend-destination field specifies the system metrics destination.
    # It supports either prometheus (the default) or stackdriver.
    # Note: Using stackdriver will incur additional charges
    metrics.backend-destination: prometheus

    # metrics.request-metrics-backend-destination specifies the request metrics
    # destination. If non-empty, it enables queue proxy to send request metrics.
    # Currently supported values: prometheus, stackdriver.
    metrics.request-metrics-backend-destination: prometheus

    # metrics.stackdriver-project-id field specifies the stackdriver project ID. This
    # field is optional. When running on GCE, application default credentials will be
    # used if this field is not provided.
    metrics.stackdriver-project-id: "<your stackdriver project id>"

    # metrics.allow-stackdriver-custom-metrics indicates whether it is allowed to send metrics to
    # Stackdriver using "global" resource type and custom metric type if the
    # metrics are not supported by "knative_revision" resource type. Setting this
    # flag to "true" could cause extra Stackdriver charge.
    # If metrics.backend-destination is not Stackdriver, this is ignored.
    metrics.allow-stackdriver-custom-metrics: "false"
kind: ConfigMap
metadata:
  name: config-observability
  namespace: knative-eventing
---

apiVersion: v1
data:
  _example: |
    ################################
    #                              #
    #    EXAMPLE CONFIGURATION     #
    #                              #
    ################################

    # This block is not actually functional configuration,
    # but serves to illustrate the available configuration
    # options and document them in a way that is accessible
    # to users that ` + "`" + `kubectl edit` + "`" + ` this config map.
    #
    # These sample configuration options may be copied out of
    # this example block and unindented to be in the data block
    # to actually change the configuration.
    #
    # If true we enable adding spans within our applications.
    enable: "false"

    # URL to zipkin collector where traces are sent.
    zipkin-endpoint: "http://zipkin.istio-system.svc.cluster.local:9411/api/v2/spans"

    # Enable zipkin debug mode. This allows all spans to be sent to the server
    # bypassing sampling.
    debug: "false"

    # Percentage (0-1) of requests to trace
    sample-rate: "0.1"
kind: ConfigMap
metadata:
  name: config-tracing
  namespace: knative-eventing
---

# in-memory-channel-crd.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    duck.knative.dev/channelable: "true"
    eventing.knative.dev/release: v0.7.0
  name: imc-channelable-manipulator
rules:
- apiGroups:
  - messaging.knative.dev
  resources:
  - inmemorychannels
  - inmemorychannels/status
  verbs:
  - create
  - get
  - list
  - watch
  - update
  - patch
---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: imc-controller
rules:
- apiGroups:
  - messaging.knative.dev
  resources:
  - inmemorychannels
  - inmemorychannels/status
  verbs:
  - get
  - list
  - watch
  - update
- apiGroups:
  - messaging.knative.dev
  resources:
  - inmemorychannels/finalizers
  verbs:
  - update
- apiGroups:
  - ""
  resources:
  - services
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
- apiGroups:
  - ""
  resources:
  - endpoints
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - apps
  resources:
  - deployments
  - deployments/status
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - create
  - patch
---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: imc-dispatcher
rules:
- apiGroups:
  - messaging.knative.dev
  resources:
  - inmemorychannels
  - inmemorychannels/status
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - messaging.knative.dev
  resources:
  - inmemorychannels/status
  verbs:
  - update
---

apiVersion: v1
kind: Service
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
    messaging.knative.dev/channel: in-memory-channel
    messaging.knative.dev/role: dispatcher
  name: imc-dispatcher
  namespace: knative-eventing
spec:
  ports:
  - port: 80
    protocol: TCP
    targetPort: 8080
  selector:
    messaging.knative.dev/channel: in-memory-channel
    messaging.knative.dev/role: dispatcher
---

apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: imc-controller
  namespace: knative-eventing---

apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: imc-dispatcher
  namespace: knative-eventing
---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: imc-controller
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: imc-controller
subjects:
- kind: ServiceAccount
  name: imc-controller
  namespace: knative-eventing---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: imc-dispatcher
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: imc-dispatcher
subjects:
- kind: ServiceAccount
  name: imc-dispatcher
  namespace: knative-eventing
---

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
    knative.dev/crd-install: "true"
    messaging.knative.dev/subscribable: "true"
  name: inmemorychannels.messaging.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .status.conditions[?(@.type=="Ready")].status
    name: Ready
    type: string
  - JSONPath: .status.conditions[?(@.type=="Ready")].reason
    name: Reason
    type: string
  - JSONPath: .status.address.hostname
    name: Hostname
    type: string
  - JSONPath: .metadata.creationTimestamp
    name: Age
    type: date
  group: messaging.knative.dev
  names:
    categories:
    - all
    - knative
    - messaging
    - channel
    kind: InMemoryChannel
    plural: inmemorychannels
    shortNames:
    - imc
    singular: inmemorychannel
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        spec:
          properties:
            subscribable:
              properties:
                subscribers:
                  items:
                    properties:
                      ref:
                        properties:
                          apiVersion:
                            type: string
                          kind:
                            type: string
                          name:
                            minLength: 1
                            type: string
                          namespace:
                            minLength: 1
                            type: string
                          uid:
                            minLength: 1
                            type: string
                        required:
                        - namespace
                        - name
                        - uid
                        type: object
                      replyURI:
                        minLength: 1
                        type: string
                      subscriberURI:
                        minLength: 1
                        type: string
                      uid:
                        minLength: 1
                        type: string
                    required:
                    - uid
                  type: array
              type: object
  version: v1alpha1
---

apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: imc-controller
  namespace: knative-eventing
spec:
  replicas: 1
  selector:
    matchLabels:
      messaging.knative.dev/channel: in-memory-channel
      messaging.knative.dev/role: controller
  template:
    metadata:
      labels:
        messaging.knative.dev/channel: in-memory-channel
        messaging.knative.dev/role: controller
    spec:
      containers:
      - env:
        - name: CONFIG_LOGGING_NAME
          value: config-logging
        - name: CONFIG_OBSERVABILITY_NAME
          value: config-observability
        - name: METRICS_DOMAIN
          value: knative.dev/inmemorychannel-controller
        - name: SYSTEM_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        image: gcr.io/knative-releases/github.com/knative/eventing/cmd/in_memory/channel_controller@sha256:292b2dddf074ce355f5793f3d4893ad0863152e0783f32c05e7ae50328b1e2e6
        name: controller
        volumeMounts:
        - mountPath: /etc/config-logging
          name: config-logging
      serviceAccountName: imc-controller
      volumes:
      - configMap:
          name: config-logging
        name: config-logging
---

apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: imc-dispatcher
  namespace: knative-eventing
spec:
  replicas: 1
  selector:
    matchLabels:
      messaging.knative.dev/channel: in-memory-channel
      messaging.knative.dev/role: dispatcher
  template:
    metadata:
      labels:
        messaging.knative.dev/channel: in-memory-channel
        messaging.knative.dev/role: dispatcher
    spec:
      containers:
      - env:
        - name: CONFIG_LOGGING_NAME
          value: config-logging
        - name: CONFIG_OBSERVABILITY_NAME
          value: config-observability
        - name: METRICS_DOMAIN
          value: knative.dev/inmemorychannel-dispatcher
        - name: SYSTEM_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        image: gcr.io/knative-releases/github.com/knative/eventing/cmd/in_memory/channel_dispatcher@sha256:0c95cafd668283cb045fa4941922ff7365f8e6caef623a9c5a68452be1404b5e
        name: dispatcher
        volumeMounts:
        - mountPath: /etc/config-logging
          name: config-logging
      serviceAccountName: imc-dispatcher
      volumes:
      - configMap:
          name: config-logging
        name: config-logging
---

# in-memory-channel-provisioner.yaml
apiVersion: eventing.knative.dev/v1alpha1
kind: ClusterChannelProvisioner
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: in-memory
spec: {}---

apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: in-memory-channel-controller
  namespace: knative-eventing---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: in-memory-channel-controller
rules:
- apiGroups:
  - eventing.knative.dev
  resources:
  - channels
  - channels/status
  - clusterchannelprovisioners
  - clusterchannelprovisioners/status
  verbs:
  - get
  - list
  - watch
  - update
- apiGroups:
  - eventing.knative.dev
  resources:
  - channels/finalizers
  - clusterchannelprovisioners/finalizers
  verbs:
  - update
- apiGroups:
  - ""
  resources:
  - services
  verbs:
  - get
  - list
  - watch
  - create
- apiGroups:
  - ""
  resources:
  - services
  verbs:
  - update
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - create
  - patch
  - update---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: in-memory-channel-controller
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: in-memory-channel-controller
subjects:
- kind: ServiceAccount
  name: in-memory-channel-controller
  namespace: knative-eventing---

apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: in-memory-channel-controller
  namespace: knative-eventing
spec:
  replicas: 1
  selector:
    matchLabels:
      clusterChannelProvisioner: in-memory-channel
      role: controller
  template:
    metadata:
      labels:
        clusterChannelProvisioner: in-memory-channel
        role: controller
    spec:
      containers:
      - env:
        - name: SYSTEM_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        image: gcr.io/knative-releases/github.com/knative/eventing/cmd/in_memory/controller@sha256:dfd7b2852c9bc2e391b04193a50a4f635db0e7a4bbd79a20c61199e6880394fe
        name: controller
      serviceAccountName: in-memory-channel-controller---

apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: in-memory-channel-dispatcher
  namespace: knative-eventing---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: in-memory-channel-dispatcher
rules:
- apiGroups:
  - eventing.knative.dev
  resources:
  - channels
  - channels/status
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: in-memory-channel-dispatcher
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: in-memory-channel-dispatcher
subjects:
- kind: ServiceAccount
  name: in-memory-channel-dispatcher
  namespace: knative-eventing---

apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: in-memory-channel-dispatcher
  namespace: knative-eventing
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch---

apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: in-memory-channel-dispatcher
  namespace: knative-eventing
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: in-memory-channel-dispatcher
subjects:
- kind: ServiceAccount
  name: in-memory-channel-dispatcher
  namespace: knative-eventing---

apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    eventing.knative.dev/release: v0.7.0
  name: in-memory-channel-dispatcher
  namespace: knative-eventing
spec:
  replicas: 1
  selector:
    matchLabels:
      clusterChannelProvisioner: in-memory-channel
      role: dispatcher
  template:
    metadata:
      labels:
        clusterChannelProvisioner: in-memory-channel
        role: dispatcher
    spec:
      containers:
      - env:
        - name: SYSTEM_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        image: gcr.io/knative-releases/github.com/knative/eventing/cmd/in_memory/dispatcher@sha256:345888cb32ce69a45165e02fc87ebfde903fd1d9dfe059892289d66b779e6bee
        name: dispatcher
      serviceAccountName: in-memory-channel-dispatcher
---
apiVersion: v1
kind: Namespace
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: knative-monitoring
---

apiVersion: v1
kind: Service
metadata:
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    app: elasticsearch-logging
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: Elasticsearch
  name: elasticsearch-logging
  namespace: knative-monitoring
spec:
  ports:
  - port: 9200
    protocol: TCP
    targetPort: db
  selector:
    app: elasticsearch-logging---

apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    app: elasticsearch-logging
    kubernetes.io/cluster-service: "true"
  name: elasticsearch-logging
  namespace: knative-monitoring---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    app: elasticsearch-logging
    kubernetes.io/cluster-service: "true"
  name: elasticsearch-logging
rules:
- apiGroups:
  - ""
  resources:
  - services
  - namespaces
  - endpoints
  verbs:
  - get---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    app: elasticsearch-logging
    kubernetes.io/cluster-service: "true"
  name: elasticsearch-logging
  namespace: knative-monitoring
roleRef:
  apiGroup: ""
  kind: ClusterRole
  name: elasticsearch-logging
subjects:
- apiGroup: ""
  kind: ServiceAccount
  name: elasticsearch-logging
  namespace: knative-monitoring---

apiVersion: apps/v1
kind: StatefulSet
metadata:
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    app: elasticsearch-logging
    kubernetes.io/cluster-service: "true"
    version: v5.6.4
  name: elasticsearch-logging
  namespace: knative-monitoring
spec:
  replicas: 2
  selector:
    matchLabels:
      app: elasticsearch-logging
      version: v5.6.4
  serviceName: elasticsearch-logging
  template:
    metadata:
      labels:
        app: elasticsearch-logging
        kubernetes.io/cluster-service: "true"
        version: v5.6.4
    spec:
      containers:
      - env:
        - name: NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        image: k8s.gcr.io/elasticsearch:v5.6.4
        name: elasticsearch-logging
        ports:
        - containerPort: 9200
          name: db
          protocol: TCP
        - containerPort: 9300
          name: transport
          protocol: TCP
        resources:
          limits:
            cpu: 1000m
          requests:
            cpu: 100m
        volumeMounts:
        - mountPath: /data
          name: elasticsearch-logging
      initContainers:
      - command:
        - /sbin/sysctl
        - -w
        - vm.max_map_count=262144
        image: alpine:3.6
        name: elasticsearch-logging-init
        securityContext:
          privileged: true
      serviceAccountName: elasticsearch-logging
      volumes:
      - emptyDir: {}
        name: elasticsearch-logging
---

apiVersion: v1
kind: Service
metadata:
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    app: kibana-logging
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: Kibana
  name: kibana-logging
  namespace: knative-monitoring
spec:
  ports:
  - port: 5601
    protocol: TCP
    targetPort: ui
  selector:
    app: kibana-logging
  type: NodePort---

apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    app: kibana-logging
    kubernetes.io/cluster-service: "true"
  name: kibana-logging
  namespace: knative-monitoring
spec:
  replicas: 1
  selector:
    matchLabels:
      app: kibana-logging
  template:
    metadata:
      labels:
        app: kibana-logging
    spec:
      containers:
      - env:
        - name: ELASTICSEARCH_URL
          value: http://elasticsearch-logging:9200
        - name: SERVER_BASEPATH
          value: /api/v1/namespaces/knative-monitoring/services/kibana-logging/proxy
        - name: XPACK_MONITORING_ENABLED
          value: "false"
        - name: XPACK_SECURITY_ENABLED
          value: "false"
        image: docker.elastic.co/kibana/kibana:5.6.4
        name: kibana-logging
        ports:
        - containerPort: 5601
          name: ui
          protocol: TCP
        resources:
          limits:
            cpu: 1000m
          requests:
            cpu: 100m
---

apiVersion: v1
data:
  100.system.conf: |-
    <system>
      root_dir /tmp/fluentd-buffers/
    </system>
  200.containers.input.conf: |-
    # Capture logs from container's stdout/stderr -> Docker -> .log in JSON format
    <source>
      @id containers-stdout-stderr
      @type tail
      path /var/log/containers/*user-container-*.log,/var/log/containers/*build-step-*.log,/var/log/containers/controller-*controller-*.log,/var/log/containers/webhook-*webhook-*.log,/var/log/containers/*autoscaler-*autoscaler-*.log,/var/log/containers/*queue-proxy-*.log,/var/log/containers/activator-*activator-*.log
      pos_file /var/log/containers-stdout-stderr.pos
      time_format %Y-%m-%dT%H:%M:%S.%NZ
      tag raw.kubernetes.*
      format json
      read_from_head true
    </source>
    # Capture logs from Knative containers' /var/log
    <source>
      @id containers-var-log
      @type tail
      # **/*/**/* allows path expansion to go through one symlink (the one created by the init container)
      path /var/lib/kubelet/pods/*/volumes/kubernetes.io~empty-dir/knative-internal/**/*/**/*
      path_key stream
      pos_file /var/log/containers-var-log.pos
      tag raw.kubernetes.*
      message_key log
      read_from_head true
      <parse>
        @type multi_format
        <pattern>
          format json
          time_key fluentd-time # fluentd-time is reserved for structured logs
          time_format %Y-%m-%dT%H:%M:%S.%NZ
        </pattern>
        <pattern>
          format none
          message_key log
        </pattern>
      </parse>
    </source>
    # Combine multi line logs which form an exception stack trace into a single log entry
    <match raw.kubernetes.**>
      @id raw.kubernetes
      @type detect_exceptions
      remove_tag_prefix raw
      message log
      stream stream
      multiline_flush_interval 5
      max_bytes 500000
      max_lines 1000
    </match>
    # Make stream path correct from the container's point of view
    <filter kubernetes.var.lib.kubelet.pods.*.volumes.kubernetes.io~empty-dir.knative-internal.*.**>
      @type record_transformer
      enable_ruby true
      <record>
        stream /var/log/${record["stream"].scan(/\/knative-internal\/[^\/]+\/(.*)/).last.last}
      </record>
    </filter>
    # Add Kubernetes metadata to logs from /var/log/containers
    <filter kubernetes.var.log.containers.**>
      @type kubernetes_metadata
    </filter>
    # Add Kubernetes metadata to logs from /var/lib/kubelet/pods/*/volumes/kubernetes.io~empty-dir/knative-internal/**/*/**/*
    <filter kubernetes.var.lib.kubelet.pods.**>
      @type kubernetes_metadata
      tag_to_kubernetes_name_regexp (?<docker_id>[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12})\.volumes.kubernetes\.io~empty-dir\.knative-internal\.(?<namespace>[^_]+)_(?<pod_name>[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*)_(?<container_name>user-container)\..*?$
    </filter>
  300.forward.input.conf: |-
    # Takes the messages sent over TCP, e.g. request logs from Istio
    <source>
      @type forward
      port 24224
    </source>
  900.output.conf: |-
    # Send to Elastic Search
    <match **>
      @id elasticsearch
      @type elasticsearch
      @log_level info
      host elasticsearch-logging
      port 9200
      logstash_format true
      <buffer>
        @type file
        path /var/log/fluentd-buffers/kubernetes.system.buffer
        flush_mode interval
        retry_type exponential_backoff
        flush_thread_count 2
        flush_interval 5s
        retry_forever
        retry_max_interval 30
        chunk_limit_size 2M
        queue_limit_length 8
        overflow_action block
      </buffer>
    </match>
kind: ConfigMap
metadata:
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    serving.knative.dev/release: "v0.7.0"
  name: fluentd-ds-config
  namespace: knative-monitoring
---

apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    app: fluentd-ds
    kubernetes.io/cluster-service: "true"
    serving.knative.dev/release: "v0.7.0"
  name: fluentd-ds
  namespace: knative-monitoring---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    app: fluentd-ds
    kubernetes.io/cluster-service: "true"
    serving.knative.dev/release: "v0.7.0"
  name: fluentd-ds
rules:
- apiGroups:
  - ""
  resources:
  - namespaces
  - pods
  verbs:
  - get
  - watch
  - list---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    app: fluentd-ds
    kubernetes.io/cluster-service: "true"
    serving.knative.dev/release: "v0.7.0"
  name: fluentd-ds
roleRef:
  apiGroup: ""
  kind: ClusterRole
  name: fluentd-ds
subjects:
- apiGroup: ""
  kind: ServiceAccount
  name: fluentd-ds
  namespace: knative-monitoring---

apiVersion: v1
kind: Service
metadata:
  labels:
    app: fluentd-ds
    serving.knative.dev/release: "v0.7.0"
  name: fluentd-ds
  namespace: knative-monitoring
spec:
  ports:
  - name: fluentd-tcp
    port: 24224
    protocol: TCP
    targetPort: 24224
  - name: fluentd-udp
    port: 24224
    protocol: UDP
    targetPort: 24224
  selector:
    app: fluentd-ds---

apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    app: fluentd-ds
    kubernetes.io/cluster-service: "true"
    serving.knative.dev/release: "v0.7.0"
    version: v2.0.4
  name: fluentd-ds
  namespace: knative-monitoring
spec:
  selector:
    matchLabels:
      app: fluentd-ds
      version: v2.0.4
  template:
    metadata:
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ""
      labels:
        app: fluentd-ds
        kubernetes.io/cluster-service: "true"
        serving.knative.dev/release: "v0.7.0"
        version: v2.0.4
    spec:
      containers:
      - env:
        - name: FLUENTD_ARGS
          value: --no-supervisor -q
        image: k8s.gcr.io/fluentd-elasticsearch:v2.0.4
        name: fluentd-ds
        resources:
          limits:
            memory: 500Mi
          requests:
            cpu: 100m
            memory: 200Mi
        volumeMounts:
        - mountPath: /var/log/containers
          name: varlogcontainers
          readOnly: true
        - mountPath: /var/log/pods
          name: varlogpods
          readOnly: true
        - mountPath: /var/lib/docker/containers
          name: varlibdockercontainers
          readOnly: true
        - mountPath: /var/lib/kubelet/pods
          name: varlibkubeletpods
          readOnly: true
        - mountPath: /host/lib
          name: libsystemddir
          readOnly: true
        - mountPath: /etc/fluent/config.d
          name: config-volume
      nodeSelector:
        beta.kubernetes.io/fluentd-ds-ready: "true"
      serviceAccountName: fluentd-ds
      terminationGracePeriodSeconds: 30
      volumes:
      - hostPath:
          path: /var/log/containers
        name: varlogcontainers
      - hostPath:
          path: /var/log/pods
        name: varlogpods
      - hostPath:
          path: /var/lib/docker/containers
        name: varlibdockercontainers
      - hostPath:
          path: /var/lib/kubelet/pods
        name: varlibkubeletpods
      - hostPath:
          path: /usr/lib64
        name: libsystemddir
      - configMap:
          name: fluentd-ds-config
        name: config-volume
---

apiVersion: v1
kind: ServiceAccount
metadata:
  name: kube-state-metrics
  namespace: knative-monitoring---

apiVersion: rbac.authorization.k8s.io/v1beta1
kind: Role
metadata:
  name: kube-state-metrics-resizer
  namespace: knative-monitoring
rules:
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - get
- apiGroups:
  - extensions
  resourceNames:
  - kube-state-metrics
  resources:
  - deployments
  verbs:
  - get
  - update---

apiVersion: rbac.authorization.k8s.io/v1beta1
kind: RoleBinding
metadata:
  name: kube-state-metrics
  namespace: knative-monitoring
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: kube-state-metrics-resizer
subjects:
- kind: ServiceAccount
  name: kube-state-metrics
  namespace: knative-monitoring---

apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: kube-state-metrics
  namespace: knative-monitoring
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  - secrets
  - nodes
  - pods
  - services
  - resourcequotas
  - replicationcontrollers
  - limitranges
  - persistentvolumeclaims
  - persistentvolumes
  - namespaces
  - endpoints
  verbs:
  - list
  - watch
- apiGroups:
  - extensions
  resources:
  - daemonsets
  - deployments
  - replicasets
  verbs:
  - list
  - watch
- apiGroups:
  - apps
  resources:
  - statefulsets
  verbs:
  - list
  - watch
- apiGroups:
  - batch
  resources:
  - cronjobs
  - jobs
  verbs:
  - list
  - watch
- apiGroups:
  - autoscaling
  resources:
  - horizontalpodautoscalers
  verbs:
  - list
  - watch
- apiGroups:
  - authentication.k8s.io
  resources:
  - tokenreviews
  verbs:
  - create
- apiGroups:
  - authorization.k8s.io
  resources:
  - subjectaccessreviews
  verbs:
  - create---

apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: kube-state-metrics
  namespace: knative-monitoring
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: kube-state-metrics
subjects:
- kind: ServiceAccount
  name: kube-state-metrics
  namespace: knative-monitoring---

apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: kube-state-metrics
  namespace: knative-monitoring
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: kube-state-metrics
    spec:
      containers:
      - args:
        - --secure-listen-address=:8443
        - --upstream=http://127.0.0.1:8081/
        image: quay.io/coreos/kube-rbac-proxy:v0.3.0
        name: kube-rbac-proxy-main
        ports:
        - containerPort: 8443
          name: https-main
        resources:
          limits:
            cpu: 20m
            memory: 40Mi
          requests:
            cpu: 10m
            memory: 20Mi
      - args:
        - --secure-listen-address=:9443
        - --upstream=http://127.0.0.1:8082/
        image: quay.io/coreos/kube-rbac-proxy:v0.3.0
        name: kube-rbac-proxy-self
        ports:
        - containerPort: 9443
          name: https-self
        resources:
          limits:
            cpu: 20m
            memory: 40Mi
          requests:
            cpu: 10m
            memory: 20Mi
      - args:
        - --host=127.0.0.1
        - --port=8081
        - --telemetry-host=127.0.0.1
        - --telemetry-port=8082
        image: quay.io/coreos/kube-state-metrics:v1.3.0
        name: kube-state-metrics
      - command:
        - /pod_nanny
        - --container=kube-state-metrics
        - --cpu=100m
        - --extra-cpu=1m
        - --memory=100Mi
        - --extra-memory=2Mi
        - --threshold=5
        - --deployment=kube-state-metrics
        env:
        - name: MY_POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: MY_POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        image: k8s.gcr.io/addon-resizer:1.7
        name: addon-resizer
        resources:
          limits:
            cpu: 100m
            memory: 30Mi
          requests:
            cpu: 100m
            memory: 30Mi
      securityContext:
        runAsNonRoot: true
        runAsUser: 65534
      serviceAccountName: kube-state-metrics---

apiVersion: v1
kind: Service
metadata:
  labels:
    app: kube-state-metrics
  name: kube-state-metrics
  namespace: knative-monitoring
spec:
  clusterIP: None
  ports:
  - name: https-main
    port: 8443
    protocol: TCP
    targetPort: https-main
  - name: https-self
    port: 9443
    protocol: TCP
    targetPort: https-self
  selector:
    app: kube-state-metrics
---

apiVersion: v1
data:
  kubernetes-deployment-dashboard.json: |
    {
      "__inputs": [
        {
          "description": "",
          "label": "prometheus",
          "name": "prometheus",
          "pluginId": "prometheus",
          "pluginName": "Prometheus",
          "type": "datasource"
        }
      ],
      "annotations": {
        "list": []
      },
      "editable": false,
      "gnetId": null,
      "graphTooltip": 0,
      "id": null,
      "links": [],
      "rows": [
        {
          "collapse": false,
          "editable": false,
          "height": "200px",
          "panels": [
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 8,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "cores",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 4,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": true
              },
              "targets": [
                {
                  "expr": "sum(rate(container_cpu_usage_seconds_total{namespace=\"$deployment_namespace\",pod_name=~\"$deployment_name.*\"}[3m]))",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "title": "CPU",
              "type": "singlestat",
              "valueFontSize": "110%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 9,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "GB",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "80%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 4,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": true
              },
              "targets": [
                {
                  "expr": "sum(container_memory_usage_bytes{namespace=\"$deployment_namespace\",pod_name=~\"$deployment_name.*\"}) / 1024^3",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "title": "Memory",
              "type": "singlestat",
              "valueFontSize": "110%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "Bps",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": false
              },
              "id": 7,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 4,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": true
              },
              "targets": [
                {
                  "expr": "sum(rate(container_network_transmit_bytes_total{namespace=\"$deployment_namespace\",pod_name=~\"$deployment_name.*\"}[3m])) + sum(rate(container_network_receive_bytes_total{namespace=\"$deployment_namespace\",pod_name=~\"$deployment_name.*\"}[3m]))",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "title": "Network",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            }
          ],
          "showTitle": false,
          "title": "Dashboard Row",
          "titleSize": "h6"
        },
        {
          "collapse": false,
          "editable": false,
          "height": "100px",
          "panels": [
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": false
              },
              "id": 5,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "max(kube_deployment_spec_replicas{deployment=\"$deployment_name\",namespace=\"$deployment_namespace\"}) without (instance, pod)",
                  "intervalFactor": 2,
                  "metric": "kube_deployment_spec_replicas",
                  "refId": "A",
                  "step": 600
                }
              ],
              "title": "Desired Replicas",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 6,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "min(kube_deployment_status_replicas_available{deployment=\"$deployment_name\",namespace=\"$deployment_namespace\"}) without (instance, pod)",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "title": "Available Replicas",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 3,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "max(kube_deployment_status_observed_generation{deployment=\"$deployment_name\",namespace=\"$deployment_namespace\"}) without (instance, pod)",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "title": "Observed Generation",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 2,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "max(kube_deployment_metadata_generation{deployment=\"$deployment_name\",namespace=\"$deployment_namespace\"}) without (instance, pod)",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "title": "Metadata Generation",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            }
          ],
          "showTitle": false,
          "title": "Dashboard Row",
          "titleSize": "h6"
        },
        {
          "collapse": false,
          "editable": false,
          "height": "350px",
          "panels": [
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 1,
              "isNew": true,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 2,
              "links": [],
              "nullPointMode": "connected",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [],
              "spaceLength": 10,
              "span": 12,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "max(kube_deployment_status_replicas{deployment=\"$deployment_name\",namespace=\"$deployment_namespace\"}) without (instance, pod)",
                  "intervalFactor": 2,
                  "legendFormat": "current replicas",
                  "refId": "A",
                  "step": 30
                },
                {
                  "expr": "min(kube_deployment_status_replicas_available{deployment=\"$deployment_name\",namespace=\"$deployment_namespace\"}) without (instance, pod)",
                  "intervalFactor": 2,
                  "legendFormat": "available",
                  "refId": "B",
                  "step": 30
                },
                {
                  "expr": "max(kube_deployment_status_replicas_unavailable{deployment=\"$deployment_name\",namespace=\"$deployment_namespace\"}) without (instance, pod)",
                  "intervalFactor": 2,
                  "legendFormat": "unavailable",
                  "refId": "C",
                  "step": 30
                },
                {
                  "expr": "min(kube_deployment_status_replicas_updated{deployment=\"$deployment_name\",namespace=\"$deployment_namespace\"}) without (instance, pod)",
                  "intervalFactor": 2,
                  "legendFormat": "updated",
                  "refId": "D",
                  "step": 30
                },
                {
                  "expr": "max(kube_deployment_spec_replicas{deployment=\"$deployment_name\",namespace=\"$deployment_namespace\"}) without (instance, pod)",
                  "intervalFactor": 2,
                  "legendFormat": "desired",
                  "refId": "E",
                  "step": 30
                }
              ],
              "title": "Replicas",
              "tooltip": {
                "msResolution": true,
                "shared": true,
                "sort": 0,
                "value_type": "cumulative"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "none",
                  "label": "",
                  "logBase": 1,
                  "show": true
                },
                {
                  "format": "short",
                  "label": "",
                  "logBase": 1,
                  "show": false
                }
              ]
            }
          ],
          "showTitle": false,
          "title": "Dashboard Row",
          "titleSize": "h6"
        }
      ],
      "schemaVersion": 14,
      "sharedCrosshair": false,
      "style": "dark",
      "tags": [],
      "templating": {
        "list": [
          {
            "allValue": ".*",
            "current": {},
            "datasource": "prometheus",
            "hide": 0,
            "includeAll": false,
            "label": "Namespace",
            "multi": false,
            "name": "deployment_namespace",
            "options": [],
            "query": "label_values(kube_deployment_metadata_generation, namespace)",
            "refresh": 1,
            "regex": "",
            "sort": 0,
            "tagValuesQuery": null,
            "tags": [],
            "tagsQuery": "",
            "type": "query",
            "useTags": false
          },
          {
            "allValue": null,
            "current": {},
            "datasource": "prometheus",
            "hide": 0,
            "includeAll": false,
            "label": "Deployment",
            "multi": false,
            "name": "deployment_name",
            "options": [],
            "query": "label_values(kube_deployment_metadata_generation{namespace=\"$deployment_namespace\"}, deployment)",
            "refresh": 1,
            "regex": "",
            "sort": 0,
            "tagValuesQuery": "",
            "tags": [],
            "tagsQuery": "deployment",
            "type": "query",
            "useTags": false
          }
        ]
      },
      "time": {
        "from": "now-6h",
        "to": "now"
      },
      "timepicker": {
        "refresh_intervals": [
          "5s",
          "10s",
          "30s",
          "1m",
          "5m",
          "15m",
          "30m",
          "1h",
          "2h",
          "1d"
        ],
        "time_options": [
          "5m",
          "15m",
          "1h",
          "6h",
          "12h",
          "24h",
          "2d",
          "7d",
          "30d"
        ]
      },
      "timezone": "browser",
      "title": "Deployment",
      "version": 1
    }
kind: ConfigMap
metadata:
  name: grafana-dashboard-definition-kubernetes-deployment
  namespace: knative-monitoring
---

apiVersion: v1
data:
  kubernetes-capacity-planning-dashboard.json: |
    {
      "__inputs": [
        {
          "description": "",
          "label": "prometheus",
          "name": "prometheus",
          "pluginId": "prometheus",
          "pluginName": "Prometheus",
          "type": "datasource"
        }
      ],
      "annotations": {
        "list": []
      },
      "editable": false,
      "gnetId": null,
      "graphTooltip": 0,
      "id": null,
      "links": [],
      "refresh": false,
      "rows": [
        {
          "collapse": false,
          "editable": false,
          "height": "250px",
          "panels": [
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 3,
              "isNew": false,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 2,
              "links": [],
              "nullPointMode": "connected",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [],
              "spaceLength": 10,
              "span": 6,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "sum(rate(node_cpu{mode=\"idle\"}[2m])) * 100",
                  "hide": false,
                  "intervalFactor": 10,
                  "legendFormat": "",
                  "refId": "A",
                  "step": 50
                }
              ],
              "title": "Idle CPU",
              "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 0,
                "value_type": "cumulative"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "percent",
                  "label": "cpu usage",
                  "logBase": 1,
                  "min": 0,
                  "show": true
                },
                {
                  "format": "short",
                  "logBase": 1,
                  "show": true
                }
              ]
            },
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 9,
              "isNew": false,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 2,
              "links": [],
              "nullPointMode": "connected",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [],
              "spaceLength": 10,
              "span": 6,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "sum(node_load1)",
                  "intervalFactor": 4,
                  "legendFormat": "load 1m",
                  "refId": "A",
                  "step": 20,
                  "target": ""
                },
                {
                  "expr": "sum(node_load5)",
                  "intervalFactor": 4,
                  "legendFormat": "load 5m",
                  "refId": "B",
                  "step": 20,
                  "target": ""
                },
                {
                  "expr": "sum(node_load15)",
                  "intervalFactor": 4,
                  "legendFormat": "load 15m",
                  "refId": "C",
                  "step": 20,
                  "target": ""
                }
              ],
              "title": "System Load",
              "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 0,
                "value_type": "cumulative"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "percentunit",
                  "logBase": 1,
                  "show": true
                },
                {
                  "format": "short",
                  "logBase": 1,
                  "show": true
                }
              ]
            }
          ],
          "showTitle": false,
          "title": "New Row",
          "titleSize": "h6"
        },
        {
          "collapse": false,
          "editable": false,
          "height": "250px",
          "panels": [
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 4,
              "isNew": false,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 2,
              "links": [],
              "nullPointMode": "connected",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [
                {
                  "alias": "node_memory_SwapFree{instance=\"172.17.0.1:9100\",job=\"prometheus\"}",
                  "yaxis": 2
                }
              ],
              "spaceLength": 10,
              "span": 9,
              "stack": true,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "sum(node_memory_MemTotal) - sum(node_memory_MemFree) - sum(node_memory_Buffers) - sum(node_memory_Cached)",
                  "intervalFactor": 2,
                  "legendFormat": "memory usage",
                  "metric": "memo",
                  "refId": "A",
                  "step": 10,
                  "target": ""
                },
                {
                  "expr": "sum(node_memory_Buffers)",
                  "interval": "",
                  "intervalFactor": 2,
                  "legendFormat": "memory buffers",
                  "metric": "memo",
                  "refId": "B",
                  "step": 10,
                  "target": ""
                },
                {
                  "expr": "sum(node_memory_Cached)",
                  "interval": "",
                  "intervalFactor": 2,
                  "legendFormat": "memory cached",
                  "metric": "memo",
                  "refId": "C",
                  "step": 10,
                  "target": ""
                },
                {
                  "expr": "sum(node_memory_MemFree)",
                  "interval": "",
                  "intervalFactor": 2,
                  "legendFormat": "memory free",
                  "metric": "memo",
                  "refId": "D",
                  "step": 10,
                  "target": ""
                }
              ],
              "title": "Memory Usage",
              "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 0,
                "value_type": "individual"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "bytes",
                  "logBase": 1,
                  "min": "0",
                  "show": true
                },
                {
                  "format": "short",
                  "logBase": 1,
                  "show": true
                }
              ]
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "percent",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": true,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "hideTimeOverride": false,
              "id": 5,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "((sum(node_memory_MemTotal) - sum(node_memory_MemFree) - sum(node_memory_Buffers) - sum(node_memory_Cached)) / sum(node_memory_MemTotal)) * 100",
                  "intervalFactor": 2,
                  "metric": "",
                  "refId": "A",
                  "step": 60,
                  "target": ""
                }
              ],
              "thresholds": "80, 90",
              "title": "Memory Usage",
              "transparent": false,
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            }
          ],
          "showTitle": false,
          "title": "New Row",
          "titleSize": "h6"
        },
        {
          "collapse": false,
          "editable": false,
          "height": "246px",
          "panels": [
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 6,
              "isNew": false,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 2,
              "links": [],
              "nullPointMode": "connected",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [
                {
                  "alias": "read",
                  "yaxis": 1
                },
                {
                  "alias": "{instance=\"172.17.0.1:9100\"}",
                  "yaxis": 2
                },
                {
                  "alias": "io time",
                  "yaxis": 2
                }
              ],
              "spaceLength": 10,
              "span": 9,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "sum(rate(node_disk_bytes_read[5m]))",
                  "hide": false,
                  "intervalFactor": 4,
                  "legendFormat": "read",
                  "refId": "A",
                  "step": 20,
                  "target": ""
                },
                {
                  "expr": "sum(rate(node_disk_bytes_written[5m]))",
                  "intervalFactor": 4,
                  "legendFormat": "written",
                  "refId": "B",
                  "step": 20
                },
                {
                  "expr": "sum(rate(node_disk_io_time_ms[5m]))",
                  "intervalFactor": 4,
                  "legendFormat": "io time",
                  "refId": "C",
                  "step": 20
                }
              ],
              "title": "Disk I/O",
              "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 0,
                "value_type": "cumulative"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "bytes",
                  "logBase": 1,
                  "show": true
                },
                {
                  "format": "ms",
                  "logBase": 1,
                  "show": true
                }
              ]
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "percentunit",
              "gauge": {
                "maxValue": 1,
                "minValue": 0,
                "show": true,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "hideTimeOverride": false,
              "id": 12,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "(sum(node_filesystem_size{device!=\"rootfs\"}) - sum(node_filesystem_free{device!=\"rootfs\"})) / sum(node_filesystem_size{device!=\"rootfs\"})",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 60,
                  "target": ""
                }
              ],
              "thresholds": "0.75, 0.9",
              "title": "Disk Space Usage",
              "transparent": false,
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "current"
            }
          ],
          "showTitle": false,
          "title": "New Row",
          "titleSize": "h6"
        },
        {
          "collapse": false,
          "editable": false,
          "height": "250px",
          "panels": [
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 8,
              "isNew": false,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 2,
              "links": [],
              "nullPointMode": "connected",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [
                {
                  "alias": "transmitted",
                  "yaxis": 2
                }
              ],
              "spaceLength": 10,
              "span": 6,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "sum(rate(node_network_receive_bytes{device!~\"lo\"}[5m]))",
                  "hide": false,
                  "intervalFactor": 2,
                  "legendFormat": "",
                  "refId": "A",
                  "step": 10,
                  "target": ""
                }
              ],
              "title": "Network Received",
              "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 0,
                "value_type": "cumulative"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "bytes",
                  "logBase": 1,
                  "show": true
                },
                {
                  "format": "bytes",
                  "logBase": 1,
                  "show": true
                }
              ]
            },
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 10,
              "isNew": false,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 2,
              "links": [],
              "nullPointMode": "connected",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [
                {
                  "alias": "transmitted",
                  "yaxis": 2
                }
              ],
              "spaceLength": 10,
              "span": 6,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "sum(rate(node_network_transmit_bytes{device!~\"lo\"}[5m]))",
                  "hide": false,
                  "intervalFactor": 2,
                  "legendFormat": "",
                  "refId": "B",
                  "step": 10,
                  "target": ""
                }
              ],
              "title": "Network Transmitted",
              "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 0,
                "value_type": "cumulative"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "bytes",
                  "logBase": 1,
                  "show": true
                },
                {
                  "format": "bytes",
                  "logBase": 1,
                  "show": true
                }
              ]
            }
          ],
          "showTitle": false,
          "title": "New Row",
          "titleSize": "h6"
        },
        {
          "collapse": false,
          "editable": false,
          "height": "276px",
          "panels": [
            {
              "aliasColors": {},
              "bars": false,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 11,
              "isNew": true,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 2,
              "links": [],
              "nullPointMode": "connected",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [],
              "spaceLength": 11,
              "span": 9,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "sum(kube_pod_info)",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "legendFormat": "Current number of Pods",
                  "refId": "A",
                  "step": 10
                },
                {
                  "expr": "sum(kube_node_status_capacity_pods)",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "legendFormat": "Maximum capacity of pods",
                  "refId": "B",
                  "step": 10
                }
              ],
              "title": "Cluster Pod Utilization",
              "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 0,
                "value_type": "individual"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "short",
                  "logBase": 1,
                  "show": true
                },
                {
                  "format": "short",
                  "logBase": 1,
                  "show": true
                }
              ]
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "percent",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": true,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "hideTimeOverride": false,
              "id": 7,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "100 - (sum(kube_node_status_capacity_pods) - sum(kube_pod_info)) / sum(kube_node_status_capacity_pods) * 100",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "legendFormat": "",
                  "refId": "A",
                  "step": 60,
                  "target": ""
                }
              ],
              "thresholds": "80, 90",
              "title": "Pod Utilization",
              "transparent": false,
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "current"
            }
          ],
          "showTitle": false,
          "title": "New Row",
          "titleSize": "h6"
        }
      ],
      "schemaVersion": 14,
      "sharedCrosshair": false,
      "style": "dark",
      "tags": [],
      "templating": {
        "list": []
      },
      "time": {
        "from": "now-1h",
        "to": "now"
      },
      "timepicker": {
        "refresh_intervals": [
          "5s",
          "10s",
          "30s",
          "1m",
          "5m",
          "15m",
          "30m",
          "1h",
          "2h",
          "1d"
        ],
        "time_options": [
          "5m",
          "15m",
          "1h",
          "6h",
          "12h",
          "24h",
          "2d",
          "7d",
          "30d"
        ]
      },
      "timezone": "browser",
      "title": "Kubernetes Capacity Planning",
      "version": 4
    }
kind: ConfigMap
metadata:
  name: grafana-dashboard-definition-kubernetes-capacity-planning
  namespace: knative-monitoring
---

apiVersion: v1
data:
  kubernetes-cluster-health-dashboard.json: |
    {
      "__inputs": [
        {
          "description": "",
          "label": "prometheus",
          "name": "prometheus",
          "pluginId": "prometheus",
          "pluginName": "Prometheus",
          "type": "datasource"
        }
      ],
      "annotations": {
        "list": []
      },
      "editable": false,
      "gnetId": null,
      "graphTooltip": 0,
      "id": null,
      "links": [],
      "refresh": "10s",
      "rows": [
        {
          "collapse": false,
          "editable": false,
          "height": "254px",
          "panels": [
            {
              "colorBackground": false,
              "colorValue": true,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "hideTimeOverride": false,
              "id": 1,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "sum(up{job=~\"apiserver|kube-scheduler|kube-controller-manager\"} == 0)",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "legendFormat": "",
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "1, 3",
              "title": "Control Plane Components Down",
              "transparent": false,
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "Everything UP and healthy",
                  "value": "null"
                },
                {
                  "op": "=",
                  "text": "",
                  "value": ""
                }
              ],
              "valueName": "avg"
            },
            {
              "colorBackground": false,
              "colorValue": true,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "hideTimeOverride": false,
              "id": 2,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "sum(ALERTS{alertstate=\"firing\",alertname!=\"DeadMansSwitch\"})",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "legendFormat": "",
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "1, 3",
              "title": "Alerts Firing",
              "transparent": false,
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "0",
                  "value": "null"
                }
              ],
              "valueName": "current"
            },
            {
              "colorBackground": false,
              "colorValue": true,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "hideTimeOverride": false,
              "id": 3,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "sum(ALERTS{alertstate=\"pending\",alertname!=\"DeadMansSwitch\"})",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "legendFormat": "",
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "3, 5",
              "title": "Alerts Pending",
              "transparent": false,
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "0",
                  "value": "null"
                }
              ],
              "valueName": "current"
            },
            {
              "colorBackground": false,
              "colorValue": true,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "hideTimeOverride": false,
              "id": 4,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "count(increase(kube_pod_container_status_restarts[1h]) > 5)",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "legendFormat": "",
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "1, 3",
              "title": "Crashlooping Pods",
              "transparent": false,
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "0",
                  "value": "null"
                }
              ],
              "valueName": "current"
            }
          ],
          "showTitle": false,
          "title": "Row",
          "titleSize": "h6"
        },
        {
          "collapse": false,
          "editable": false,
          "height": "250px",
          "panels": [
            {
              "colorBackground": false,
              "colorValue": true,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "hideTimeOverride": false,
              "id": 5,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "sum(kube_node_status_condition{condition=\"Ready\",status!=\"true\"})",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "legendFormat": "",
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "1, 3",
              "title": "Node Not Ready",
              "transparent": false,
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "current"
            },
            {
              "colorBackground": false,
              "colorValue": true,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "hideTimeOverride": false,
              "id": 6,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "sum(kube_node_status_condition{condition=\"DiskPressure\",status=\"true\"})",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "legendFormat": "",
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "1, 3",
              "title": "Node Disk Pressure",
              "transparent": false,
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "current"
            },
            {
              "colorBackground": false,
              "colorValue": true,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "hideTimeOverride": false,
              "id": 7,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "sum(kube_node_status_condition{condition=\"MemoryPressure\",status=\"true\"})",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "legendFormat": "",
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "1, 3",
              "title": "Node Memory Pressure",
              "transparent": false,
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "current"
            },
            {
              "colorBackground": false,
              "colorValue": true,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "hideTimeOverride": false,
              "id": 8,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "sum(kube_node_spec_unschedulable)",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "legendFormat": "",
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "1, 3",
              "title": "Nodes Unschedulable",
              "transparent": false,
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "current"
            }
          ],
          "showTitle": false,
          "title": "Row",
          "titleSize": "h6"
        }
      ],
      "schemaVersion": 14,
      "sharedCrosshair": false,
      "style": "dark",
      "tags": [],
      "templating": {
        "list": []
      },
      "time": {
        "from": "now-6h",
        "to": "now"
      },
      "timepicker": {
        "refresh_intervals": [
          "5s",
          "10s",
          "30s",
          "1m",
          "5m",
          "15m",
          "30m",
          "1h",
          "2h",
          "1d"
        ],
        "time_options": [
          "5m",
          "15m",
          "1h",
          "6h",
          "12h",
          "24h",
          "2d",
          "7d",
          "30d"
        ]
      },
      "timezone": "browser",
      "title": "Kubernetes Cluster Health",
      "version": 9
    }
kind: ConfigMap
metadata:
  name: grafana-dashboard-definition-kubernetes-cluster-health
  namespace: knative-monitoring
---

apiVersion: v1
data:
  kubernetes-cluster-status-dashboard.json: |
    {
      "__inputs": [
        {
          "description": "",
          "label": "prometheus",
          "name": "prometheus",
          "pluginId": "prometheus",
          "pluginName": "Prometheus",
          "type": "datasource"
        }
      ],
      "annotations": {
        "list": []
      },
      "editable": false,
      "graphTooltip": 0,
      "hideControls": false,
      "links": [],
      "rows": [
        {
          "collapse": false,
          "editable": false,
          "height": "129px",
          "panels": [
            {
              "colorBackground": false,
              "colorValue": true,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 5,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 6,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "sum(up{job=~\"apiserver|kube-scheduler|kube-controller-manager\"} == 0)",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "1, 3",
              "title": "Control Plane UP",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "UP",
                  "value": "null"
                }
              ],
              "valueName": "total"
            },
            {
              "colorBackground": false,
              "colorValue": true,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 6,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 6,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "sum(ALERTS{alertstate=\"firing\",alertname!=\"DeadMansSwitch\"})",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "3, 5",
              "title": "Alerts Firing",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "0",
                  "value": "null"
                }
              ],
              "valueName": "current"
            }
          ],
          "showTitle": true,
          "title": "Cluster Health",
          "titleSize": "h6"
        },
        {
          "collapse": false,
          "editable": false,
          "height": "168px",
          "panels": [
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "percent",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": true,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 1,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "(sum(up{job=\"apiserver\"} == 1) / count(up{job=\"apiserver\"})) * 100",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "50, 80",
              "title": "API Servers UP",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "current"
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "percent",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": true,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 2,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "(sum(up{job=\"kube-controller-manager\"} == 1) / count(up{job=\"kube-controller-manager\"})) * 100",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "50, 80",
              "title": "Controller Managers UP",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "current"
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "percent",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": true,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 3,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "(sum(up{job=\"kube-scheduler\"} == 1) / count(up{job=\"kube-scheduler\"})) * 100",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "50, 80",
              "title": "Schedulers UP",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "current"
            },
            {
              "colorBackground": false,
              "colorValue": true,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 4,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "count(increase(kube_pod_container_status_restarts{namespace=~\"kube-system|tectonic-system\"}[1h]) > 5)",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "1, 3",
              "title": "Crashlooping Control Plane Pods",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "0",
                  "value": "null"
                }
              ],
              "valueName": "current"
            }
          ],
          "showTitle": true,
          "title": "Control Plane Status",
          "titleSize": "h6"
        },
        {
          "collapse": false,
          "editable": false,
          "height": "158px",
          "panels": [
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "percent",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": true,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 8,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "sum(100 - (avg by (instance) (rate(node_cpu{job=\"node-exporter\",mode=\"idle\"}[5m])) * 100)) / count(node_cpu{job=\"node-exporter\",mode=\"idle\"})",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "80, 90",
              "title": "CPU Utilization",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "percent",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": true,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 7,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "((sum(node_memory_MemTotal) - sum(node_memory_MemFree) - sum(node_memory_Buffers) - sum(node_memory_Cached)) / sum(node_memory_MemTotal)) * 100",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "80, 90",
              "title": "Memory Utilization",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "percent",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": true,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 9,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "(sum(node_filesystem_size{device!=\"rootfs\"}) - sum(node_filesystem_free{device!=\"rootfs\"})) / sum(node_filesystem_size{device!=\"rootfs\"})",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "80, 90",
              "title": "Filesystem Utilization",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "percent",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": true,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 10,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "100 - (sum(kube_node_status_capacity_pods) - sum(kube_pod_info)) / sum(kube_node_status_capacity_pods) * 100",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "80, 90",
              "title": "Pod Utilization",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            }
          ],
          "showTitle": true,
          "title": "Capacity Planning",
          "titleSize": "h6"
        }
      ],
      "schemaVersion": 14,
      "sharedCrosshair": false,
      "style": "dark",
      "tags": [],
      "templating": {
        "list": []
      },
      "time": {
        "from": "now-6h",
        "to": "now"
      },
      "timepicker": {
        "refresh_intervals": [
          "5s",
          "10s",
          "30s",
          "1m",
          "5m",
          "15m",
          "30m",
          "1h",
          "2h",
          "1d"
        ],
        "time_options": [
          "5m",
          "15m",
          "1h",
          "6h",
          "12h",
          "24h",
          "2d",
          "7d",
          "30d"
        ]
      },
      "timezone": "browser",
      "title": "Kubernetes Cluster Status",
      "version": 3
    }
kind: ConfigMap
metadata:
  name: grafana-dashboard-definition-kubernetes-cluster-status
  namespace: knative-monitoring
---

apiVersion: v1
data:
  kubernetes-control-plane-status-dashboard.json: |
    {
      "__inputs": [
        {
          "description": "",
          "label": "prometheus",
          "name": "prometheus",
          "pluginId": "prometheus",
          "pluginName": "Prometheus",
          "type": "datasource"
        }
      ],
      "annotations": {
        "list": []
      },
      "editable": false,
      "graphTooltip": 0,
      "hideControls": false,
      "links": [],
      "rows": [
        {
          "collapse": false,
          "editable": false,
          "height": "250px",
          "panels": [
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "percent",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": true,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "hideTimeOverride": false,
              "id": 1,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "(sum(up{job=\"apiserver\"} == 1) / sum(up{job=\"apiserver\"})) * 100",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "50, 80",
              "title": "API Servers UP",
              "transparent": false,
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "percent",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": true,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "hideTimeOverride": false,
              "id": 2,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "(sum(up{job=\"kube-controller-manager\"} == 1) / sum(up{job=\"kube-controller-manager\"})) * 100",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "50, 80",
              "title": "Controller Managers UP",
              "transparent": false,
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "percent",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": true,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "hideTimeOverride": false,
              "id": 3,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "(sum(up{job=\"kube-scheduler\"} == 1) / sum(up{job=\"kube-scheduler\"})) * 100",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "50, 80",
              "title": "Schedulers UP",
              "transparent": false,
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "percent",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": true,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "hideTimeOverride": false,
              "id": 4,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "max(sum by(instance) (rate(apiserver_request_count{code=~\"5..\"}[5m])) / sum by(instance) (rate(apiserver_request_count[5m]))) * 100",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "legendFormat": "",
                  "refId": "A",
                  "step": 600
                }
              ],
              "thresholds": "5, 10",
              "title": "API Server Request Error Rate",
              "transparent": false,
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "0",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            }
          ],
          "showTitle": false,
          "title": "Dashboard Row",
          "titleSize": "h6"
        },
        {
          "collapse": false,
          "editable": false,
          "height": "250px",
          "panels": [
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 7,
              "isNew": false,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 1,
              "links": [],
              "nullPointMode": "null",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [],
              "spaceLength": 10,
              "span": 12,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "sum by(verb) (rate(apiserver_latency_seconds:quantile[5m]) >= 0)",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "legendFormat": "",
                  "refId": "A",
                  "step": 30
                }
              ],
              "title": "API Server Request Latency",
              "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 0,
                "value_type": "individual"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "short",
                  "logBase": 1,
                  "show": true
                },
                {
                  "format": "short",
                  "logBase": 1,
                  "show": true
                }
              ]
            }
          ],
          "showTitle": false,
          "title": "Dashboard Row",
          "titleSize": "h6"
        },
        {
          "collapse": false,
          "editable": false,
          "height": "250px",
          "panels": [
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 5,
              "isNew": false,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 1,
              "links": [],
              "nullPointMode": "null",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [],
              "spaceLength": 10,
              "span": 6,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "cluster:scheduler_e2e_scheduling_latency_seconds:quantile",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 60
                }
              ],
              "title": "End to End Scheduling Latency",
              "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 0,
                "value_type": "individual"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "short",
                  "logBase": 1,
                  "show": true
                },
                {
                  "format": "dtdurations",
                  "logBase": 1,
                  "show": true
                }
              ]
            },
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 6,
              "isNew": false,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 1,
              "links": [],
              "nullPointMode": "null",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [],
              "spaceLength": 10,
              "span": 6,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "sum by(instance) (rate(apiserver_request_count{code!~\"2..\"}[5m]))",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "legendFormat": "Error Rate",
                  "refId": "A",
                  "step": 60
                },
                {
                  "expr": "sum by(instance) (rate(apiserver_request_count[5m]))",
                  "format": "time_series",
                  "intervalFactor": 2,
                  "legendFormat": "Request Rate",
                  "refId": "B",
                  "step": 60
                }
              ],
              "title": "API Server Request Rates",
              "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 0,
                "value_type": "individual"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "short",
                  "logBase": 1,
                  "show": true
                },
                {
                  "format": "short",
                  "logBase": 1,
                  "show": true
                }
              ]
            }
          ],
          "showTitle": false,
          "title": "Dashboard Row",
          "titleSize": "h6"
        }
      ],
      "schemaVersion": 14,
      "sharedCrosshair": false,
      "style": "dark",
      "tags": [],
      "templating": {
        "list": []
      },
      "time": {
        "from": "now-6h",
        "to": "now"
      },
      "timepicker": {
        "refresh_intervals": [
          "5s",
          "10s",
          "30s",
          "1m",
          "5m",
          "15m",
          "30m",
          "1h",
          "2h",
          "1d"
        ],
        "time_options": [
          "5m",
          "15m",
          "1h",
          "6h",
          "12h",
          "24h",
          "2d",
          "7d",
          "30d"
        ]
      },
      "timezone": "browser",
      "title": "Kubernetes Control Plane Status",
      "version": 3
    }
kind: ConfigMap
metadata:
  name: grafana-dashboard-definition-kubernetes-control-plane-status
  namespace: knative-monitoring
---

apiVersion: v1
data:
  kubernetes-resource-requests-dashboard.json: |
    {
      "__inputs": [
        {
          "description": "",
          "label": "prometheus",
          "name": "prometheus",
          "pluginId": "prometheus",
          "pluginName": "Prometheus",
          "type": "datasource"
        }
      ],
      "annotations": {
        "list": []
      },
      "editable": false,
      "graphTooltip": 0,
      "hideControls": false,
      "links": [],
      "refresh": false,
      "rows": [
        {
          "collapse": false,
          "editable": false,
          "height": "300px",
          "panels": [
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "description": "This represents the total [CPU resource requests](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#meaning-of-cpu) in the cluster.\nFor comparison the total [allocatable CPU cores](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/node-allocatable.md) is also shown.",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 1,
              "isNew": false,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 1,
              "links": [],
              "nullPointMode": "null",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [],
              "spaceLength": 10,
              "span": 9,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "min(sum(kube_node_status_allocatable_cpu_cores) by (instance))",
                  "hide": false,
                  "intervalFactor": 2,
                  "legendFormat": "Allocatable CPU Cores",
                  "refId": "A",
                  "step": 20
                },
                {
                  "expr": "max(sum(kube_pod_container_resource_requests_cpu_cores) by (instance))",
                  "hide": false,
                  "intervalFactor": 2,
                  "legendFormat": "Requested CPU Cores",
                  "refId": "B",
                  "step": 20
                }
              ],
              "title": "CPU Cores",
              "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 0,
                "value_type": "individual"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "short",
                  "label": "CPU Cores",
                  "logBase": 1,
                  "show": true
                },
                {
                  "format": "short",
                  "logBase": 1,
                  "show": true
                }
              ]
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "percent",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": true,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "hideTimeOverride": false,
              "id": 2,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": true
              },
              "targets": [
                {
                  "expr": "max(sum(kube_pod_container_resource_requests_cpu_cores) by (instance)) / min(sum(kube_node_status_allocatable_cpu_cores) by (instance)) * 100",
                  "intervalFactor": 2,
                  "legendFormat": "",
                  "refId": "A",
                  "step": 240
                }
              ],
              "thresholds": "80, 90",
              "title": "CPU Cores",
              "transparent": false,
              "type": "singlestat",
              "valueFontSize": "110%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            }
          ],
          "showTitle": false,
          "title": "CPU Cores",
          "titleSize": "h6"
        },
        {
          "collapse": false,
          "editable": false,
          "height": "300px",
          "panels": [
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "description": "This represents the total [memory resource requests](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#meaning-of-memory) in the cluster.\nFor comparison the total [allocatable memory](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/node-allocatable.md) is also shown.",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 3,
              "isNew": false,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 1,
              "links": [],
              "nullPointMode": "null",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [],
              "spaceLength": 10,
              "span": 9,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "min(sum(kube_node_status_allocatable_memory_bytes) by (instance))",
                  "hide": false,
                  "intervalFactor": 2,
                  "legendFormat": "Allocatable Memory",
                  "refId": "A",
                  "step": 20
                },
                {
                  "expr": "max(sum(kube_pod_container_resource_requests_memory_bytes) by (instance))",
                  "hide": false,
                  "intervalFactor": 2,
                  "legendFormat": "Requested Memory",
                  "refId": "B",
                  "step": 20
                }
              ],
              "title": "Memory",
              "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 0,
                "value_type": "individual"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "bytes",
                  "label": "Memory",
                  "logBase": 1,
                  "show": true
                },
                {
                  "format": "short",
                  "logBase": 1,
                  "show": true
                }
              ]
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "percent",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": true,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "hideTimeOverride": false,
              "id": 4,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": true
              },
              "targets": [
                {
                  "expr": "max(sum(kube_pod_container_resource_requests_memory_bytes) by (instance)) / min(sum(kube_node_status_allocatable_memory_bytes) by (instance)) * 100",
                  "intervalFactor": 2,
                  "legendFormat": "",
                  "refId": "A",
                  "step": 240
                }
              ],
              "thresholds": "80, 90",
              "title": "Memory",
              "transparent": false,
              "type": "singlestat",
              "valueFontSize": "110%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            }
          ],
          "showTitle": false,
          "title": "Memory",
          "titleSize": "h6"
        }
      ],
      "schemaVersion": 14,
      "sharedCrosshair": false,
      "style": "dark",
      "tags": [],
      "templating": {
        "list": []
      },
      "time": {
        "from": "now-3h",
        "to": "now"
      },
      "timepicker": {
        "refresh_intervals": [
          "5s",
          "10s",
          "30s",
          "1m",
          "5m",
          "15m",
          "30m",
          "1h",
          "2h",
          "1d"
        ],
        "time_options": [
          "5m",
          "15m",
          "1h",
          "6h",
          "12h",
          "24h",
          "2d",
          "7d",
          "30d"
        ]
      },
      "timezone": "browser",
      "title": "Kubernetes Resource Requests",
      "version": 2
    }
kind: ConfigMap
metadata:
  name: grafana-dashboard-definition-kubernetes-resource-requests
  namespace: knative-monitoring
---

apiVersion: v1
data:
  kubernetes-nodes-dashboard.json: |
    {
      "__inputs": [
        {
          "description": "",
          "label": "prometheus",
          "name": "prometheus",
          "pluginId": "prometheus",
          "pluginName": "Prometheus",
          "type": "datasource"
        }
      ],
      "annotations": {
        "list": []
      },
      "description": "Dashboard to get an overview of one server",
      "editable": false,
      "gnetId": 22,
      "graphTooltip": 0,
      "hideControls": false,
      "links": [],
      "refresh": false,
      "rows": [
        {
          "collapse": false,
          "editable": false,
          "height": "250px",
          "panels": [
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 3,
              "isNew": false,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 2,
              "links": [],
              "nullPointMode": "connected",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [],
              "spaceLength": 10,
              "span": 6,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "100 - (avg by (cpu) (irate(node_cpu{mode=\"idle\", instance=\"$server\"}[5m])) * 100)",
                  "hide": false,
                  "intervalFactor": 10,
                  "legendFormat": "{{cpu}}",
                  "refId": "A",
                  "step": 50
                }
              ],
              "title": "Idle CPU",
              "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 0,
                "value_type": "cumulative"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "percent",
                  "label": "cpu usage",
                  "logBase": 1,
                  "max": 100,
                  "min": 0,
                  "show": true
                },
                {
                  "format": "short",
                  "logBase": 1,
                  "show": true
                }
              ]
            },
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 9,
              "isNew": false,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 2,
              "links": [],
              "nullPointMode": "connected",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [],
              "spaceLength": 10,
              "span": 6,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "node_load1{instance=\"$server\"}",
                  "intervalFactor": 4,
                  "legendFormat": "load 1m",
                  "refId": "A",
                  "step": 20,
                  "target": ""
                },
                {
                  "expr": "node_load5{instance=\"$server\"}",
                  "intervalFactor": 4,
                  "legendFormat": "load 5m",
                  "refId": "B",
                  "step": 20,
                  "target": ""
                },
                {
                  "expr": "node_load15{instance=\"$server\"}",
                  "intervalFactor": 4,
                  "legendFormat": "load 15m",
                  "refId": "C",
                  "step": 20,
                  "target": ""
                }
              ],
              "title": "System Load",
              "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 0,
                "value_type": "cumulative"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "percentunit",
                  "logBase": 1,
                  "show": true
                },
                {
                  "format": "short",
                  "logBase": 1,
                  "show": true
                }
              ]
            }
          ],
          "showTitle": false,
          "title": "New Row",
          "titleSize": "h6"
        },
        {
          "collapse": false,
          "editable": false,
          "height": "250px",
          "panels": [
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 4,
              "isNew": false,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 2,
              "links": [],
              "nullPointMode": "connected",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [
                {
                  "alias": "node_memory_SwapFree{instance=\"172.17.0.1:9100\",job=\"prometheus\"}",
                  "yaxis": 2
                }
              ],
              "spaceLength": 10,
              "span": 9,
              "stack": true,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "node_memory_MemTotal{instance=\"$server\"} - node_memory_MemFree{instance=\"$server\"} - node_memory_Buffers{instance=\"$server\"} - node_memory_Cached{instance=\"$server\"}",
                  "hide": false,
                  "interval": "",
                  "intervalFactor": 2,
                  "legendFormat": "memory used",
                  "metric": "",
                  "refId": "C",
                  "step": 10
                },
                {
                  "expr": "node_memory_Buffers{instance=\"$server\"}",
                  "interval": "",
                  "intervalFactor": 2,
                  "legendFormat": "memory buffers",
                  "metric": "",
                  "refId": "E",
                  "step": 10
                },
                {
                  "expr": "node_memory_Cached{instance=\"$server\"}",
                  "intervalFactor": 2,
                  "legendFormat": "memory cached",
                  "metric": "",
                  "refId": "F",
                  "step": 10
                },
                {
                  "expr": "node_memory_MemFree{instance=\"$server\"}",
                  "intervalFactor": 2,
                  "legendFormat": "memory free",
                  "metric": "",
                  "refId": "D",
                  "step": 10
                }
              ],
              "title": "Memory Usage",
              "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 0,
                "value_type": "individual"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "bytes",
                  "logBase": 1,
                  "min": "0",
                  "show": true
                },
                {
                  "format": "short",
                  "logBase": 1,
                  "show": true
                }
              ]
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "percent",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": true,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "hideTimeOverride": false,
              "id": 5,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "((node_memory_MemTotal{instance=\"$server\"} - node_memory_MemFree{instance=\"$server\"}  - node_memory_Buffers{instance=\"$server\"} - node_memory_Cached{instance=\"$server\"}) / node_memory_MemTotal{instance=\"$server\"}) * 100",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 60,
                  "target": ""
                }
              ],
              "thresholds": "80, 90",
              "title": "Memory Usage",
              "transparent": false,
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            }
          ],
          "showTitle": false,
          "title": "New Row",
          "titleSize": "h6"
        },
        {
          "collapse": false,
          "editable": false,
          "height": "250px",
          "panels": [
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 6,
              "isNew": true,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 2,
              "links": [],
              "nullPointMode": "connected",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [
                {
                  "alias": "read",
                  "yaxis": 1
                },
                {
                  "alias": "{instance=\"172.17.0.1:9100\"}",
                  "yaxis": 2
                },
                {
                  "alias": "io time",
                  "yaxis": 2
                }
              ],
              "spaceLength": 10,
              "span": 9,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "sum by (instance) (rate(node_disk_bytes_read{instance=\"$server\"}[2m]))",
                  "hide": false,
                  "intervalFactor": 4,
                  "legendFormat": "read",
                  "refId": "A",
                  "step": 20,
                  "target": ""
                },
                {
                  "expr": "sum by (instance) (rate(node_disk_bytes_written{instance=\"$server\"}[2m]))",
                  "intervalFactor": 4,
                  "legendFormat": "written",
                  "refId": "B",
                  "step": 20
                },
                {
                  "expr": "sum by (instance) (rate(node_disk_io_time_ms{instance=\"$server\"}[2m]))",
                  "intervalFactor": 4,
                  "legendFormat": "io time",
                  "refId": "C",
                  "step": 20
                }
              ],
              "title": "Disk I/O",
              "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 0,
                "value_type": "cumulative"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "bytes",
                  "logBase": 1,
                  "show": true
                },
                {
                  "format": "ms",
                  "logBase": 1,
                  "show": true
                }
              ]
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(50, 172, 45, 0.97)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(245, 54, 54, 0.9)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "percentunit",
              "gauge": {
                "maxValue": 1,
                "minValue": 0,
                "show": true,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "hideTimeOverride": false,
              "id": 7,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "(sum(node_filesystem_size{device!=\"rootfs\",instance=\"$server\"}) - sum(node_filesystem_free{device!=\"rootfs\",instance=\"$server\"})) / sum(node_filesystem_size{device!=\"rootfs\",instance=\"$server\"})",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 60,
                  "target": ""
                }
              ],
              "thresholds": "0.75, 0.9",
              "title": "Disk Space Usage",
              "transparent": false,
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "current"
            }
          ],
          "showTitle": false,
          "title": "New Row",
          "titleSize": "h6"
        },
        {
          "collapse": false,
          "editable": false,
          "height": "250px",
          "panels": [
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 8,
              "isNew": false,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 2,
              "links": [],
              "nullPointMode": "connected",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [
                {
                  "alias": "transmitted",
                  "yaxis": 2
                }
              ],
              "spaceLength": 10,
              "span": 6,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "rate(node_network_receive_bytes{instance=\"$server\",device!~\"lo\"}[5m])",
                  "hide": false,
                  "intervalFactor": 2,
                  "legendFormat": "{{device}}",
                  "refId": "A",
                  "step": 10,
                  "target": ""
                }
              ],
              "title": "Network Received",
              "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 0,
                "value_type": "cumulative"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "bytes",
                  "logBase": 1,
                  "show": true
                },
                {
                  "format": "bytes",
                  "logBase": 1,
                  "show": true
                }
              ]
            },
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 10,
              "isNew": false,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 2,
              "links": [],
              "nullPointMode": "connected",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [
                {
                  "alias": "transmitted",
                  "yaxis": 2
                }
              ],
              "spaceLength": 10,
              "span": 6,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "rate(node_network_transmit_bytes{instance=\"$server\",device!~\"lo\"}[5m])",
                  "hide": false,
                  "intervalFactor": 2,
                  "legendFormat": "{{device}}",
                  "refId": "B",
                  "step": 10,
                  "target": ""
                }
              ],
              "title": "Network Transmitted",
              "tooltip": {
                "msResolution": false,
                "shared": true,
                "sort": 0,
                "value_type": "cumulative"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "bytes",
                  "logBase": 1,
                  "show": true
                },
                {
                  "format": "bytes",
                  "logBase": 1,
                  "show": true
                }
              ]
            }
          ],
          "showTitle": false,
          "title": "New Row",
          "titleSize": "h6"
        }
      ],
      "schemaVersion": 14,
      "sharedCrosshair": false,
      "style": "dark",
      "tags": [],
      "templating": {
        "list": [
          {
            "allValue": null,
            "current": {},
            "datasource": "prometheus",
            "hide": 0,
            "includeAll": false,
            "label": null,
            "multi": false,
            "name": "server",
            "options": [],
            "query": "label_values(node_boot_time, instance)",
            "refresh": 1,
            "regex": "",
            "sort": 0,
            "tagValuesQuery": "",
            "tags": [],
            "tagsQuery": "",
            "type": "query",
            "useTags": false
          }
        ]
      },
      "time": {
        "from": "now-1h",
        "to": "now"
      },
      "timepicker": {
        "refresh_intervals": [
          "5s",
          "10s",
          "30s",
          "1m",
          "5m",
          "15m",
          "30m",
          "1h",
          "2h",
          "1d"
        ],
        "time_options": [
          "5m",
          "15m",
          "1h",
          "6h",
          "12h",
          "24h",
          "2d",
          "7d",
          "30d"
        ]
      },
      "timezone": "browser",
      "title": "Nodes",
      "version": 2
    }
kind: ConfigMap
metadata:
  name: grafana-dashboard-definition-kubernetes-nodes
  namespace: knative-monitoring
---

apiVersion: v1
data:
  kubernetes-pods-dashboard.json: |
    {
      "__inputs": [
        {
          "description": "",
          "label": "prometheus",
          "name": "prometheus",
          "pluginId": "prometheus",
          "pluginName": "Prometheus",
          "type": "datasource"
        }
      ],
      "annotations": {
        "list": []
      },
      "editable": false,
      "graphTooltip": 1,
      "hideControls": false,
      "links": [],
      "refresh": false,
      "rows": [
        {
          "collapse": false,
          "editable": false,
          "height": "250px",
          "panels": [
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 1,
              "isNew": false,
              "legend": {
                "alignAsTable": true,
                "avg": true,
                "current": true,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": true,
                "show": true,
                "total": false,
                "values": true
              },
              "lines": true,
              "linewidth": 2,
              "links": [],
              "nullPointMode": "connected",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [],
              "spaceLength": 10,
              "span": 12,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "sum by(container_name) (container_memory_usage_bytes{pod_name=\"$pod\", container_name=~\"$container\", container_name!=\"POD\"})",
                  "interval": "10s",
                  "intervalFactor": 1,
                  "legendFormat": "Current: {{ container_name }}",
                  "metric": "container_memory_usage_bytes",
                  "refId": "A",
                  "step": 15
                },
                {
                  "expr": "kube_pod_container_resource_requests_memory_bytes{pod=\"$pod\", container=~\"$container\"}",
                  "interval": "10s",
                  "intervalFactor": 2,
                  "legendFormat": "Requested: {{ container }}",
                  "metric": "kube_pod_container_resource_requests_memory_bytes",
                  "refId": "B",
                  "step": 20
                },
                {
                  "expr": "kube_pod_container_resource_limits_memory_bytes{pod=\"$pod\", container=~\"$container\"}",
                  "interval": "10s",
                  "intervalFactor": 2,
                  "legendFormat": "Limit: {{ container }}",
                  "metric": "kube_pod_container_resource_limits_memory_bytes",
                  "refId": "C",
                  "step": 20
                }
              ],
              "title": "Memory Usage",
              "tooltip": {
                "msResolution": true,
                "shared": true,
                "sort": 0,
                "value_type": "cumulative"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "bytes",
                  "logBase": 1,
                  "show": true
                },
                {
                  "format": "short",
                  "logBase": 1,
                  "show": true
                }
              ]
            }
          ],
          "showTitle": false,
          "title": "Row",
          "titleSize": "h6"
        },
        {
          "collapse": false,
          "editable": false,
          "height": "250px",
          "panels": [
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 2,
              "isNew": false,
              "legend": {
                "alignAsTable": true,
                "avg": true,
                "current": true,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": true,
                "show": true,
                "total": false,
                "values": true
              },
              "lines": true,
              "linewidth": 2,
              "links": [],
              "nullPointMode": "connected",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [],
              "spaceLength": 10,
              "span": 12,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "sum by (container_name)(rate(container_cpu_usage_seconds_total{image!=\"\",container_name!=\"POD\",pod_name=\"$pod\"}[1m]))",
                  "intervalFactor": 2,
                  "legendFormat": "{{ container_name }}",
                  "refId": "A",
                  "step": 30
                },
                {
                  "expr": "kube_pod_container_resource_requests_cpu_cores{pod=\"$pod\", container=~\"$container\"}",
                  "interval": "10s",
                  "intervalFactor": 2,
                  "legendFormat": "Requested: {{ container }}",
                  "metric": "kube_pod_container_resource_requests_cpu_cores",
                  "refId": "B",
                  "step": 20
                },
                {
                  "expr": "kube_pod_container_resource_limits_cpu_cores{pod=\"$pod\", container=~\"$container\"}",
                  "interval": "10s",
                  "intervalFactor": 2,
                  "legendFormat": "Limit: {{ container }}",
                  "metric": "kube_pod_container_resource_limits_memory_bytes",
                  "refId": "C",
                  "step": 20
                }
              ],
              "title": "CPU Usage",
              "tooltip": {
                "msResolution": true,
                "shared": true,
                "sort": 0,
                "value_type": "cumulative"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "short",
                  "logBase": 1,
                  "show": true
                },
                {
                  "format": "short",
                  "logBase": 1,
                  "show": true
                }
              ]
            }
          ],
          "showTitle": false,
          "title": "Row",
          "titleSize": "h6"
        },
        {
          "collapse": false,
          "editable": false,
          "height": "250px",
          "panels": [
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 3,
              "isNew": false,
              "legend": {
                "alignAsTable": true,
                "avg": true,
                "current": true,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": true,
                "show": true,
                "total": false,
                "values": true
              },
              "lines": true,
              "linewidth": 2,
              "links": [],
              "nullPointMode": "connected",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [],
              "spaceLength": 10,
              "span": 12,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "sort_desc(sum by (pod_name) (rate(container_network_receive_bytes_total{pod_name=\"$pod\"}[1m])))",
                  "intervalFactor": 2,
                  "legendFormat": "{{ pod_name }}",
                  "refId": "A",
                  "step": 30
                }
              ],
              "title": "Network I/O",
              "tooltip": {
                "msResolution": true,
                "shared": true,
                "sort": 0,
                "value_type": "cumulative"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "bytes",
                  "logBase": 1,
                  "show": true
                },
                {
                  "format": "short",
                  "logBase": 1,
                  "show": true
                }
              ]
            }
          ],
          "showTitle": false,
          "title": "New Row",
          "titleSize": "h6"
        }
      ],
      "schemaVersion": 14,
      "sharedCrosshair": false,
      "style": "dark",
      "tags": [],
      "templating": {
        "list": [
          {
            "allValue": ".*",
            "current": {},
            "datasource": "prometheus",
            "hide": 0,
            "includeAll": true,
            "label": "Namespace",
            "multi": false,
            "name": "namespace",
            "options": [],
            "query": "label_values(kube_pod_info, namespace)",
            "refresh": 1,
            "regex": "",
            "sort": 0,
            "tagValuesQuery": "",
            "tags": [],
            "tagsQuery": "",
            "type": "query",
            "useTags": false
          },
          {
            "allValue": null,
            "current": {},
            "datasource": "prometheus",
            "hide": 0,
            "includeAll": false,
            "label": "Pod",
            "multi": false,
            "name": "pod",
            "options": [],
            "query": "label_values(kube_pod_info{namespace=~\"$namespace\"}, pod)",
            "refresh": 1,
            "regex": "",
            "sort": 0,
            "tagValuesQuery": "",
            "tags": [],
            "tagsQuery": "",
            "type": "query",
            "useTags": false
          },
          {
            "allValue": ".*",
            "current": {},
            "datasource": "prometheus",
            "hide": 0,
            "includeAll": true,
            "label": "Container",
            "multi": false,
            "name": "container",
            "options": [],
            "query": "label_values(kube_pod_container_info{namespace=\"$namespace\", pod=\"$pod\"}, container)",
            "refresh": 1,
            "regex": "",
            "sort": 0,
            "tagValuesQuery": "",
            "tags": [],
            "tagsQuery": "",
            "type": "query",
            "useTags": false
          }
        ]
      },
      "time": {
        "from": "now-6h",
        "to": "now"
      },
      "timepicker": {
        "refresh_intervals": [
          "5s",
          "10s",
          "30s",
          "1m",
          "5m",
          "15m",
          "30m",
          "1h",
          "2h",
          "1d"
        ],
        "time_options": [
          "5m",
          "15m",
          "1h",
          "6h",
          "12h",
          "24h",
          "2d",
          "7d",
          "30d"
        ]
      },
      "timezone": "browser",
      "title": "Pods",
      "version": 1
    }
kind: ConfigMap
metadata:
  name: grafana-dashboard-definition-kubernetes-pods
  namespace: knative-monitoring
---

apiVersion: v1
data:
  kubernetes-statefulset-dashboard.json: |
    {
      "__inputs": [
        {
          "description": "",
          "label": "prometheus",
          "name": "prometheus",
          "pluginId": "prometheus",
          "pluginName": "Prometheus",
          "type": "datasource"
        }
      ],
      "annotations": {
        "list": []
      },
      "editable": false,
      "graphTooltip": 1,
      "hideControls": false,
      "links": [],
      "rows": [
        {
          "collapse": false,
          "editable": false,
          "height": "200px",
          "panels": [
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 8,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "cores",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 4,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": true
              },
              "targets": [
                {
                  "expr": "sum(rate(container_cpu_usage_seconds_total{namespace=\"$statefulset_namespace\",pod_name=~\"$statefulset_name.*\"}[3m]))",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "title": "CPU",
              "type": "singlestat",
              "valueFontSize": "110%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 9,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "GB",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "80%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 4,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": true
              },
              "targets": [
                {
                  "expr": "sum(container_memory_usage_bytes{namespace=\"$statefulset_namespace\",pod_name=~\"$statefulset_name.*\"}) / 1024^3",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "title": "Memory",
              "type": "singlestat",
              "valueFontSize": "110%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "Bps",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": false
              },
              "id": 7,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfix": "",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 4,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": true
              },
              "targets": [
                {
                  "expr": "sum(rate(container_network_transmit_bytes_total{namespace=\"$statefulset_namespace\",pod_name=~\"$statefulset_name.*\"}[3m])) + sum(rate(container_network_receive_bytes_total{namespace=\"$statefulset_namespace\",pod_name=~\"$statefulset_name.*\"}[3m]))",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "title": "Network",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            }
          ],
          "showTitle": false,
          "title": "Dashboard Row",
          "titleSize": "h6"
        },
        {
          "collapse": false,
          "editable": false,
          "height": "100px",
          "panels": [
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": false
              },
              "id": 5,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "max(kube_statefulset_replicas{statefulset=\"$statefulset_name\",namespace=\"$statefulset_namespace\"}) without (instance, pod)",
                  "intervalFactor": 2,
                  "metric": "kube_statefulset_replicas",
                  "refId": "A",
                  "step": 600
                }
              ],
              "title": "Desired Replicas",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 6,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "min(kube_statefulset_status_replicas{statefulset=\"$statefulset_name\",namespace=\"$statefulset_namespace\"}) without (instance, pod)",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "title": "Available Replicas",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 3,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "max(kube_statefulset_status_observed_generation{statefulset=\"$statefulset_name\",namespace=\"$statefulset_namespace\"}) without (instance, pod)",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "title": "Observed Generation",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            },
            {
              "colorBackground": false,
              "colorValue": false,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "datasource": "prometheus",
              "editable": false,
              "format": "none",
              "gauge": {
                "maxValue": 100,
                "minValue": 0,
                "show": false,
                "thresholdLabels": false,
                "thresholdMarkers": true
              },
              "id": 2,
              "links": [],
              "mappingType": 1,
              "mappingTypes": [
                {
                  "name": "value to text",
                  "value": 1
                },
                {
                  "name": "range to text",
                  "value": 2
                }
              ],
              "maxDataPoints": 100,
              "nullPointMode": "connected",
              "postfixFontSize": "50%",
              "prefix": "",
              "prefixFontSize": "50%",
              "rangeMaps": [
                {
                  "from": "null",
                  "text": "N/A",
                  "to": "null"
                }
              ],
              "span": 3,
              "sparkline": {
                "fillColor": "rgba(31, 118, 189, 0.18)",
                "full": false,
                "lineColor": "rgb(31, 120, 193)",
                "show": false
              },
              "targets": [
                {
                  "expr": "max(kube_statefulset_metadata_generation{statefulset=\"$statefulset_name\",namespace=\"$statefulset_namespace\"}) without (instance, pod)",
                  "intervalFactor": 2,
                  "refId": "A",
                  "step": 600
                }
              ],
              "title": "Metadata Generation",
              "type": "singlestat",
              "valueFontSize": "80%",
              "valueMaps": [
                {
                  "op": "=",
                  "text": "N/A",
                  "value": "null"
                }
              ],
              "valueName": "avg"
            }
          ],
          "showTitle": false,
          "title": "Dashboard Row",
          "titleSize": "h6"
        },
        {
          "collapse": false,
          "editable": false,
          "height": "350px",
          "panels": [
            {
              "aliasColors": {},
              "bars": false,
              "dashLength": 10,
              "dashes": false,
              "datasource": "prometheus",
              "editable": false,
              "error": false,
              "fill": 1,
              "grid": {
                "threshold1Color": "rgba(216, 200, 27, 0.27)",
                "threshold2Color": "rgba(234, 112, 112, 0.22)"
              },
              "id": 1,
              "isNew": true,
              "legend": {
                "alignAsTable": false,
                "avg": false,
                "current": false,
                "hideEmpty": false,
                "hideZero": false,
                "max": false,
                "min": false,
                "rightSide": false,
                "show": true,
                "total": false
              },
              "lines": true,
              "linewidth": 2,
              "links": [],
              "nullPointMode": "connected",
              "percentage": false,
              "pointradius": 5,
              "points": false,
              "renderer": "flot",
              "seriesOverrides": [],
              "spaceLength": 10,
              "span": 12,
              "stack": false,
              "steppedLine": false,
              "targets": [
                {
                  "expr": "min(kube_statefulset_status_replicas{statefulset=\"$statefulset_name\",namespace=\"$statefulset_namespace\"}) without (instance, pod)",
                  "intervalFactor": 2,
                  "legendFormat": "available",
                  "refId": "B",
                  "step": 30
                },
                {
                  "expr": "max(kube_statefulset_replicas{statefulset=\"$statefulset_name\",namespace=\"$statefulset_namespace\"}) without (instance, pod)",
                  "intervalFactor": 2,
                  "legendFormat": "desired",
                  "refId": "E",
                  "step": 30
                }
              ],
              "title": "Replicas",
              "tooltip": {
                "msResolution": true,
                "shared": true,
                "sort": 0,
                "value_type": "cumulative"
              },
              "type": "graph",
              "xaxis": {
                "mode": "time",
                "show": true,
                "values": []
              },
              "yaxes": [
                {
                  "format": "none",
                  "label": "",
                  "logBase": 1,
                  "show": true
                },
                {
                  "format": "short",
                  "label": "",
                  "logBase": 1,
                  "show": false
                }
              ]
            }
          ],
          "showTitle": false,
          "title": "Dashboard Row",
          "titleSize": "h6"
        }
      ],
      "schemaVersion": 14,
      "sharedCrosshair": false,
      "style": "dark",
      "tags": [],
      "templating": {
        "list": [
          {
            "allValue": ".*",
            "current": {},
            "datasource": "prometheus",
            "hide": 0,
            "includeAll": false,
            "label": "Namespace",
            "multi": false,
            "name": "statefulset_namespace",
            "options": [],
            "query": "label_values(kube_statefulset_metadata_generation, namespace)",
            "refresh": 1,
            "regex": "",
            "sort": 0,
            "tagValuesQuery": null,
            "tags": [],
            "tagsQuery": "",
            "type": "query",
            "useTags": false
          },
          {
            "allValue": null,
            "current": {},
            "datasource": "prometheus",
            "hide": 0,
            "includeAll": false,
            "label": "StatefulSet",
            "multi": false,
            "name": "statefulset_name",
            "options": [],
            "query": "label_values(kube_statefulset_metadata_generation{namespace=\"$statefulset_namespace\"}, statefulset)",
            "refresh": 1,
            "regex": "",
            "sort": 0,
            "tagValuesQuery": "",
            "tags": [],
            "tagsQuery": "statefulset",
            "type": "query",
            "useTags": false
          }
        ]
      },
      "time": {
        "from": "now-6h",
        "to": "now"
      },
      "timepicker": {
        "refresh_intervals": [
          "5s",
          "10s",
          "30s",
          "1m",
          "5m",
          "15m",
          "30m",
          "1h",
          "2h",
          "1d"
        ],
        "time_options": [
          "5m",
          "15m",
          "1h",
          "6h",
          "12h",
          "24h",
          "2d",
          "7d",
          "30d"
        ]
      },
      "timezone": "browser",
      "title": "StatefulSet",
      "version": 1
    }
kind: ConfigMap
metadata:
  name: grafana-dashboard-definition-kubernetes-statefulset
  namespace: knative-monitoring
---

apiVersion: v1
kind: ServiceAccount
metadata:
  name: node-exporter
  namespace: knative-monitoring---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: node-exporter
  namespace: knative-monitoring
rules:
- apiGroups:
  - authentication.k8s.io
  resources:
  - tokenreviews
  verbs:
  - create
- apiGroups:
  - authorization.k8s.io
  resources:
  - subjectaccessreviews
  verbs:
  - create---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: node-exporter
  namespace: knative-monitoring
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: node-exporter
subjects:
- kind: ServiceAccount
  name: node-exporter
  namespace: knative-monitoring---

apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  name: node-exporter
  namespace: knative-monitoring
spec:
  template:
    metadata:
      labels:
        app: node-exporter
      name: node-exporter
      namespace: knative-monitoring
    spec:
      containers:
      - args:
        - --web.listen-address=127.0.0.1:9101
        - --path.procfs=/host/proc
        - --path.sysfs=/host/sys
        image: quay.io/prometheus/node-exporter:v0.15.2
        name: node-exporter
        resources:
          limits:
            cpu: 200m
            memory: 50Mi
          requests:
            cpu: 100m
            memory: 30Mi
        volumeMounts:
        - mountPath: /host/proc
          name: proc
          readOnly: true
        - mountPath: /host/sys
          name: sys
          readOnly: true
      - args:
        - --secure-listen-address=:9100
        - --upstream=http://127.0.0.1:9101/
        image: quay.io/coreos/kube-rbac-proxy:v0.3.0
        name: kube-rbac-proxy
        ports:
        - containerPort: 9100
          hostPort: 9100
          name: https
        resources:
          limits:
            cpu: 20m
            memory: 40Mi
          requests:
            cpu: 10m
            memory: 20Mi
      hostNetwork: true
      hostPID: true
      securityContext:
        runAsNonRoot: true
        runAsUser: 65534
      serviceAccountName: node-exporter
      tolerations:
      - effect: NoSchedule
        operator: Exists
      volumes:
      - hostPath:
          path: /proc
        name: proc
      - hostPath:
          path: /sys
        name: sys
  updateStrategy:
    rollingUpdate:
      maxUnavailable: 1
    type: RollingUpdate---

apiVersion: v1
kind: Service
metadata:
  labels:
    app: node-exporter
  name: node-exporter
  namespace: knative-monitoring
spec:
  clusterIP: None
  ports:
  - name: https
    port: 9100
    protocol: TCP
  selector:
    app: node-exporter
  type: ClusterIP
---

apiVersion: v1
data:
  knative-control-plane-efficiency-dashboard.json: |
    {
      "__inputs": [
        {
          "description": "",
          "label": "prometheus",
          "name": "prometheus",
          "pluginId": "prometheus",
          "pluginName": "Prometheus",
          "type": "datasource"
        }
      ],
      "annotations": {
        "list": []
      },
      "description": "Knative Serving - Control Plane Efficiency",
      "editable": false,
      "gnetId": null,
      "graphTooltip": 0,
      "id": null,
      "links": [],
      "panels": [
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "prometheus",
          "decimals": 2,
          "fill": 1,
          "gridPos": {
            "h": 9,
            "w": 12,
            "x": 0,
            "y": 0
          },
          "id": 2,
          "legend": {
            "alignAsTable": true,
            "avg": true,
            "current": true,
            "max": false,
            "min": false,
            "rightSide": true,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(rate(container_cpu_usage_seconds_total{namespace=\"knative-serving\"}[1m]))",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "knative-serving",
              "refId": "A"
            },
            {
              "expr": "sum(rate(container_cpu_usage_seconds_total{namespace=\"knative-build\"}[1m]))",
              "format": "time_series",
              "instant": false,
              "interval": "",
              "intervalFactor": 1,
              "legendFormat": "knative-build",
              "refId": "C"
            },
            {
              "expr": "sum(rate(container_cpu_usage_seconds_total{namespace=\"istio-system\"}[1m]))",
              "format": "time_series",
              "instant": false,
              "interval": "",
              "intervalFactor": 1,
              "legendFormat": "istio-system",
              "refId": "D"
            },
            {
              "expr": "sum(rate(container_cpu_usage_seconds_total{namespace=\"kube-system\"}[1m]))",
              "format": "time_series",
              "instant": false,
              "intervalFactor": 1,
              "legendFormat": "kube-system",
              "refId": "F"
            },
            {
              "expr": "sum(rate(container_cpu_usage_seconds_total{namespace=\"kube-public\"}[1m]))",
              "format": "time_series",
              "instant": false,
              "intervalFactor": 1,
              "legendFormat": "kube-public",
              "refId": "E"
            },
            {
              "expr": "sum(rate(container_cpu_usage_seconds_total{namespace=\"knative-monitoring\"}[1m]))",
              "format": "time_series",
              "instant": false,
              "intervalFactor": 1,
              "legendFormat": "knative-monitoring",
              "refId": "B"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Namespace CPU Usage",
          "tooltip": {
            "shared": true,
            "sort": 2,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "prometheus",
          "decimals": null,
          "fill": 1,
          "gridPos": {
            "h": 9,
            "w": 12,
            "x": 12,
            "y": 0
          },
          "id": 3,
          "legend": {
            "alignAsTable": true,
            "avg": true,
            "current": true,
            "max": false,
            "min": false,
            "rightSide": true,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(container_memory_usage_bytes{namespace=\"knative-serving\"})",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "knative-serving",
              "refId": "A"
            },
            {
              "expr": "sum(container_memory_usage_bytes{namespace=\"knative-build\"})",
              "format": "time_series",
              "instant": false,
              "interval": "",
              "intervalFactor": 1,
              "legendFormat": "knative-build",
              "refId": "C"
            },
            {
              "expr": "sum(container_memory_usage_bytes{namespace=\"istio-system\"})",
              "format": "time_series",
              "instant": false,
              "interval": "",
              "intervalFactor": 1,
              "legendFormat": "istio-system",
              "refId": "D"
            },
            {
              "expr": "sum(container_memory_usage_bytes{namespace=\"kube-system\"})",
              "format": "time_series",
              "instant": false,
              "intervalFactor": 1,
              "legendFormat": "kube-system",
              "refId": "F"
            },
            {
              "expr": "sum(container_memory_usage_bytes{namespace=\"kube-public\"})",
              "format": "time_series",
              "instant": false,
              "intervalFactor": 1,
              "legendFormat": "kube-public",
              "refId": "E"
            },
            {
              "expr": "sum(container_memory_usage_bytes{namespace=\"knative-monitoring\"})",
              "format": "time_series",
              "instant": false,
              "intervalFactor": 1,
              "legendFormat": "knative-monitoring",
              "refId": "B"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Namespace Memory Usage",
          "tooltip": {
            "shared": true,
            "sort": 2,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "decbytes",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": false
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "prometheus",
          "decimals": 2,
          "fill": 1,
          "gridPos": {
            "h": 9,
            "w": 12,
            "x": 0,
            "y": 9
          },
          "id": 4,
          "legend": {
            "alignAsTable": true,
            "avg": true,
            "current": true,
            "max": false,
            "min": false,
            "rightSide": true,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(rate(container_cpu_usage_seconds_total{namespace!~\"knative-serving|knative-monitoring|knative-build|istio-system|kube-system|kube-public|^$\"}[1m]))",
              "format": "time_series",
              "interval": "",
              "intervalFactor": 1,
              "legendFormat": "Data plane",
              "refId": "A"
            },
            {
              "expr": "sum(rate(container_cpu_usage_seconds_total{namespace=~\"knative-serving|knative-monitoring|knative-build|istio-system|kube-system|kube-public\"}[1m]))",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "Control plane",
              "refId": "B"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Control Plane vs Data Plane CPU Usage",
          "tooltip": {
            "shared": true,
            "sort": 2,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "decimals": null,
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": false
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "prometheus",
          "fill": 1,
          "gridPos": {
            "h": 9,
            "w": 12,
            "x": 12,
            "y": 9
          },
          "id": 5,
          "legend": {
            "alignAsTable": true,
            "avg": true,
            "current": true,
            "max": false,
            "min": false,
            "rightSide": true,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(container_memory_usage_bytes{namespace!~\"knative-serving|knative-monitoring|knative-build|istio-system|kube-system|kube-public|^$\"})",
              "format": "time_series",
              "interval": "",
              "intervalFactor": 1,
              "legendFormat": "Data plane",
              "refId": "A"
            },
            {
              "expr": "sum(container_memory_usage_bytes{namespace=~\"knative-serving|knative-monitoring|knative-build|istio-system|kube-system|kube-public\"})",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "Control plane",
              "refId": "B"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Control Plane vs Data Plane Memory Usage",
          "tooltip": {
            "shared": true,
            "sort": 2,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "decbytes",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        }
      ],
      "refresh": "5s",
      "schemaVersion": 16,
      "style": "dark",
      "tags": [],
      "templating": {
        "list": []
      },
      "time": {
        "from": "now-15m",
        "to": "now"
      },
      "timepicker": {
        "refresh_intervals": [
          "5s",
          "10s",
          "30s",
          "1m",
          "5m",
          "15m",
          "30m",
          "1h",
          "2h",
          "1d"
        ],
        "time_options": [
          "5m",
          "15m",
          "1h",
          "6h",
          "12h",
          "24h",
          "2d",
          "7d",
          "30d"
        ]
      },
      "timezone": "",
      "title": "Knative Serving - Control Plane Efficiency",
      "uid": "1oI1URnik",
      "version": 2
    }
kind: ConfigMap
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: grafana-dashboard-definition-knative-efficiency
  namespace: knative-monitoring
---

apiVersion: v1
data:
  knative-reconciler-dashboard.json: |
    {
      "__inputs": [
        {
          "description": "",
          "label": "prometheus",
          "name": "prometheus",
          "pluginId": "prometheus",
          "pluginName": "Prometheus",
          "type": "datasource"
        }
      ],
      "annotations": {
        "list": []
      },
      "description": "Knative Serving - Reconciler",
      "editable": false,
      "gnetId": null,
      "graphTooltip": 0,
      "id": null,
      "links": [],
      "panels": [
        {
          "collapsed": false,
          "gridPos": {
            "h": 1,
            "w": 24,
            "x": 0,
            "y": 0
          },
          "id": 7,
          "panels": [],
          "title": "Aggregate",
          "type": "row"
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "prometheus",
          "fill": 1,
          "gridPos": {
            "h": 9,
            "w": 12,
            "x": 0,
            "y": 1
          },
          "id": 10,
          "legend": {
            "alignAsTable": false,
            "avg": false,
            "current": false,
            "hideEmpty": true,
            "hideZero": false,
            "max": false,
            "min": false,
            "rightSide": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "stack": false,
          "steppedLine": true,
          "targets": [
            {
              "expr": "sum by (reconciler)(60 * rate(controller_reconcile_count[1m]))",
              "format": "time_series",
              "instant": false,
              "interval": "",
              "intervalFactor": 1,
              "legendFormat": "{{reconciler}}",
              "refId": "A"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Reconcile Count (per min)",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "collapsed": false,
          "gridPos": {
            "h": 1,
            "w": 24,
            "x": 0,
            "y": 10
          },
          "id": 5,
          "panels": [],
          "title": "Per Reconciler",
          "type": "row"
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "prometheus",
          "fill": 1,
          "gridPos": {
            "h": 9,
            "w": 12,
            "x": 0,
            "y": 11
          },
          "id": 2,
          "legend": {
            "alignAsTable": false,
            "avg": false,
            "current": false,
            "hideEmpty": true,
            "hideZero": false,
            "max": false,
            "min": false,
            "rightSide": false,
            "show": false,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "stack": false,
          "steppedLine": true,
          "targets": [
            {
              "expr": "sum(60 * rate(controller_reconcile_count{reconciler=\"$reconciler\"}[1m]))",
              "format": "time_series",
              "instant": false,
              "interval": "",
              "intervalFactor": 1,
              "legendFormat": "{{reconciler}}",
              "refId": "A"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "$reconciler Reconcile Count (per min)",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "prometheus",
          "fill": 1,
          "gridPos": {
            "h": 9,
            "w": 12,
            "x": 12,
            "y": 11
          },
          "id": 11,
          "legend": {
            "alignAsTable": false,
            "avg": false,
            "current": false,
            "hideEmpty": true,
            "hideZero": false,
            "max": false,
            "min": false,
            "rightSide": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "stack": false,
          "steppedLine": true,
          "targets": [
            {
              "expr": "histogram_quantile(0.99, sum(rate(controller_reconcile_latency_bucket{reconciler=\"$reconciler\"}[1m])) by (le))",
              "format": "time_series",
              "instant": false,
              "interval": "",
              "intervalFactor": 1,
              "legendFormat": "99th",
              "refId": "A"
            },
            {
              "expr": "histogram_quantile(0.90, sum(rate(controller_reconcile_latency_bucket{reconciler=\"$reconciler\"}[1m])) by (le))",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "90th",
              "refId": "B"
            },
            {
              "expr": "histogram_quantile(0.50, sum(rate(controller_reconcile_latency_bucket{reconciler=\"$reconciler\"}[1m])) by (le))",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "50th",
              "refId": "C"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "$reconciler Reconcile Latency Percentiles",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "collapsed": false,
          "gridPos": {
            "h": 1,
            "w": 24,
            "x": 0,
            "y": 20
          },
          "id": 9,
          "panels": [],
          "title": "Per Reconciler & Key",
          "type": "row"
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "prometheus",
          "fill": 1,
          "gridPos": {
            "h": 9,
            "w": 12,
            "x": 0,
            "y": 21
          },
          "id": 12,
          "legend": {
            "alignAsTable": false,
            "avg": false,
            "current": false,
            "hideEmpty": true,
            "hideZero": false,
            "max": false,
            "min": false,
            "rightSide": false,
            "show": false,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "stack": false,
          "steppedLine": true,
          "targets": [
            {
              "expr": "sum(60 * rate(controller_reconcile_count{reconciler=\"$reconciler\", key=\"$key\"}[1m]))",
              "format": "time_series",
              "instant": false,
              "interval": "",
              "intervalFactor": 1,
              "legendFormat": "{{reconciler}}",
              "refId": "A"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "$reconciler/$key Reconcile Count (per min)",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "prometheus",
          "fill": 1,
          "gridPos": {
            "h": 9,
            "w": 12,
            "x": 12,
            "y": 21
          },
          "id": 13,
          "legend": {
            "alignAsTable": false,
            "avg": false,
            "current": false,
            "hideEmpty": true,
            "hideZero": false,
            "max": false,
            "min": false,
            "rightSide": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "stack": false,
          "steppedLine": true,
          "targets": [
            {
              "expr": "histogram_quantile(0.99, sum(rate(controller_reconcile_latency_bucket{reconciler=\"$reconciler\", key=\"$key\"}[1m])) by (le))",
              "format": "time_series",
              "instant": false,
              "interval": "",
              "intervalFactor": 1,
              "legendFormat": "99th",
              "refId": "A"
            },
            {
              "expr": "histogram_quantile(0.90, sum(rate(controller_reconcile_latency_bucket{reconciler=\"$reconciler\", key=\"$key\"}[1m])) by (le))",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "90th",
              "refId": "B"
            },
            {
              "expr": "histogram_quantile(0.50, sum(rate(controller_reconcile_latency_bucket{reconciler=\"$reconciler\", key=\"$key\"}[1m])) by (le))",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "50th",
              "refId": "C"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "$reconciler/$key Reconcile Latency Percentiles",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        }
      ],
      "refresh": "5s",
      "schemaVersion": 16,
      "style": "dark",
      "tags": [],
      "templating": {
        "list": [
          {
            "allValue": null,
            "current": {},
            "datasource": "prometheus",
            "hide": 0,
            "includeAll": false,
            "label": "Reconciler",
            "multi": false,
            "name": "reconciler",
            "options": [],
            "query": "label_values(controller_reconcile_count, reconciler)",
            "refresh": 1,
            "regex": "",
            "sort": 0,
            "tagValuesQuery": "",
            "tags": [],
            "tagsQuery": "",
            "type": "query",
            "useTags": false
          },
          {
            "allValue": null,
            "current": {},
            "datasource": "prometheus",
            "hide": 0,
            "includeAll": false,
            "label": "Key",
            "multi": false,
            "name": "key",
            "options": [],
            "query": "label_values(controller_reconcile_count{reconciler=\"$reconciler\"}, key)",
            "refresh": 1,
            "regex": "",
            "sort": 1,
            "tagValuesQuery": "",
            "tags": [],
            "tagsQuery": "",
            "type": "query",
            "useTags": false
          }
        ]
      },
      "time": {
        "from": "now-15m",
        "to": "now"
      },
      "timepicker": {
        "refresh_intervals": [
          "5s",
          "10s",
          "30s",
          "1m",
          "5m",
          "15m",
          "30m",
          "1h",
          "2h",
          "1d"
        ],
        "time_options": [
          "5m",
          "15m",
          "1h",
          "6h",
          "12h",
          "24h",
          "2d",
          "7d",
          "30d"
        ]
      },
      "timezone": "",
      "title": "Knative Serving - Reconciler",
      "uid": "j0oFdEYiz",
      "version": 10
    }
kind: ConfigMap
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: grafana-dashboard-definition-knative-reconciler
  namespace: knative-monitoring
---

apiVersion: v1
data:
  scaling-dashboard.json: "{  \n   \"__inputs\":[  \n      {  \n         \"name\":\"prometheus\",\n
    \        \"label\":\"prometheus\",\n         \"description\":\"\",\n         \"type\":\"datasource\",\n
    \        \"pluginId\":\"prometheus\",\n         \"pluginName\":\"Prometheus\"\n
    \     }\n   ],\n   \"__requires\":[  \n      {  \n         \"type\":\"grafana\",\n
    \        \"id\":\"grafana\",\n         \"name\":\"Grafana\",\n         \"version\":\"5.0.3\"\n
    \     },\n      {  \n         \"id\":\"graph\",\n         \"name\":\"Graph\",\n
    \        \"type\":\"panel\",\n         \"version\":\"5.0.0\"\n      },\n      {
    \ \n         \"type\":\"datasource\",\n         \"id\":\"prometheus\",\n         \"name\":\"Prometheus\",\n
    \        \"version\":\"5.0.0\"\n      }\n   ],\n   \"annotations\":{  \n      \"list\":[
    \ \n         {  \n            \"builtIn\":1,\n            \"datasource\":\"--
    Grafana --\",\n            \"enable\":true,\n            \"hide\":true,\n            \"iconColor\":\"rgba(0,
    211, 255, 1)\",\n            \"name\":\"Annotations & Alerts\",\n            \"type\":\"dashboard\"\n
    \        }\n      ]\n   },\n   \"description\":\"Knative Serving - Scaling Debugging\",\n
    \  \"editable\":false,\n   \"gnetId\":null,\n   \"graphTooltip\":0,\n   \"id\":null,\n
    \  \"iteration\":1527886043818,\n   \"links\":[  \n\n   ],\n   \"panels\":[  \n
    \     {  \n\n         \"collapsed\":true,\n         \"gridPos\":{  \n            \"h\":1,\n
    \           \"w\":24,\n            \"x\":0,\n            \"y\":0\n         },\n
    \        \"id\":14,\n         \"panels\":[  \n            {  \n               \"aliasColors\":{
    \ \n\n               },\n               \"bars\":false,\n               \"dashLength\":10,\n
    \              \"dashes\":false,\n               \"datasource\":\"prometheus\",\n
    \              \"fill\":1,\n               \"gridPos\":{  \n                  \"h\":11,\n
    \                 \"w\":24,\n                  \"x\":0,\n                  \"y\":1\n
    \              },\n               \"id\":2,\n               \"legend\":{  \n                  \"avg\":false,\n
    \                 \"current\":false,\n                  \"max\":false,\n                  \"min\":false,\n
    \                 \"show\":true,\n                  \"total\":false,\n                  \"values\":false\n
    \              },\n               \"lines\":true,\n               \"linewidth\":1,\n
    \              \"links\":[  \n\n               ],\n               \"nullPointMode\":\"null\",\n
    \              \"percentage\":false,\n               \"pointradius\":5,\n               \"points\":false,\n
    \              \"renderer\":\"flot\",\n               \"seriesOverrides\":[  \n\n
    \              ],\n               \"spaceLength\":10,\n               \"stack\":false,\n
    \              \"steppedLine\":true,\n               \"targets\":[  \n                  {
    \ \n                     \"expr\":\"sum(autoscaler_actual_pods{namespace_name=\\\"$namespace\\\",
    configuration_name=\\\"$configuration\\\", revision_name=\\\"$revision\\\"})\",\n
    \                    \"format\":\"time_series\",\n                     \"interval\":\"1s\",\n
    \                    \"intervalFactor\":1,\n                     \"legendFormat\":\"Actual
    Pods\",\n                     \"refId\":\"A\"\n                  },\n                  {
    \ \n                     \"expr\":\"sum(autoscaler_requested_pods{namespace_name=\\\"$namespace\\\",
    configuration_name=\\\"$configuration\\\", revision_name=\\\"$revision\\\"})\",\n
    \                    \"format\":\"time_series\",\n                     \"interval\":\"1s\",\n
    \                    \"intervalFactor\":1,\n                     \"legendFormat\":\"Requested
    Pods\",\n                     \"refId\":\"C\"\n                  }\n               ],\n
    \              \"thresholds\":[  \n\n               ],\n               \"timeFrom\":null,\n
    \              \"timeShift\":null,\n               \"title\":\"Revision Pod Counts\",\n
    \              \"tooltip\":{  \n                  \"shared\":true,\n                  \"sort\":0,\n
    \                 \"value_type\":\"individual\"\n               },\n               \"type\":\"graph\",\n
    \              \"xaxis\":{  \n                  \"buckets\":null,\n                  \"mode\":\"time\",\n
    \                 \"name\":null,\n                  \"show\":true,\n                  \"values\":[
    \ \n\n                  ]\n               },\n               \"yaxes\":[  \n                  {
    \ \n                     \"format\":\"short\",\n                     \"label\":null,\n
    \                    \"logBase\":1,\n                     \"max\":null,\n                     \"min\":null,\n
    \                    \"show\":true\n                  },\n                  {
    \ \n                     \"decimals\":null,\n                     \"format\":\"short\",\n
    \                    \"label\":\"Concurrency\",\n                     \"logBase\":1,\n
    \                    \"max\":\"1\",\n                     \"min\":null,\n                     \"show\":false\n
    \                 }\n               ]\n            }\n         ],\n         \"title\":\"Revision
    Pod Counts\",\n         \"type\":\"row\"\n      },\n      {  \n         \"collapsed\":true,\n
    \        \"gridPos\":{  \n            \"h\":1,\n            \"w\":24,\n            \"x\":0,\n
    \           \"y\":1\n         },\n         \"id\":18,\n         \"panels\":[  \n
    \           {  \n               \"aliasColors\":{  \n\n               },\n               \"bars\":false,\n
    \              \"dashLength\":10,\n               \"dashes\":false,\n               \"datasource\":\"prometheus\",\n
    \              \"fill\":1,\n               \"gridPos\":{  \n                  \"h\":9,\n
    \                 \"w\":12,\n                  \"x\":0,\n                  \"y\":13\n
    \              },\n               \"id\":4,\n               \"legend\":{  \n                  \"avg\":false,\n
    \                 \"current\":false,\n                  \"max\":false,\n                  \"min\":false,\n
    \                 \"show\":true,\n                  \"total\":false,\n                  \"values\":false\n
    \              },\n               \"lines\":true,\n               \"linewidth\":1,\n
    \              \"links\":[  \n\n               ],\n               \"nullPointMode\":\"null\",\n
    \              \"percentage\":false,\n               \"pointradius\":5,\n               \"points\":false,\n
    \              \"renderer\":\"flot\",\n               \"seriesOverrides\":[  \n\n
    \              ],\n               \"spaceLength\":10,\n               \"stack\":false,\n
    \              \"steppedLine\":false,\n               \"targets\":[  \n                  {
    \ \n                     \"expr\":\"sum(kube_pod_container_resource_requests_cpu_cores{namespace=\\\"$namespace\\\",
    pod=~\\\"$revision-deployment-.*\\\"})\",\n                     \"format\":\"time_series\",\n
    \                    \"interval\":\"\",\n                     \"intervalFactor\":1,\n
    \                    \"legendFormat\":\"Cores requested\",\n                     \"refId\":\"A\"\n
    \                 },\n                  {  \n                     \"expr\":\"sum(rate(container_cpu_usage_seconds_total{namespace=\\\"$namespace\\\",
    pod_name=~\\\"$revision-deployment-.*\\\"}[1m]))\",\n                     \"format\":\"time_series\",\n
    \                    \"intervalFactor\":1,\n                     \"legendFormat\":\"Cores
    used\",\n                     \"refId\":\"B\"\n                  },\n                  {
    \ \n                     \"expr\":\"sum(kube_pod_container_resource_limits_cpu_cores{namespace=\\\"$namespace\\\",
    pod=~\\\"$revision-deployment-.*\\\"})\",\n                     \"format\":\"time_series\",\n
    \                    \"intervalFactor\":1,\n                     \"legendFormat\":\"Core
    limit\",\n                     \"refId\":\"C\"\n                  }\n               ],\n
    \              \"thresholds\":[  \n\n               ],\n               \"timeFrom\":null,\n
    \              \"timeShift\":null,\n               \"title\":\"Revision CPU Usage\",\n
    \              \"tooltip\":{  \n                  \"shared\":true,\n                  \"sort\":2,\n
    \                 \"value_type\":\"individual\"\n               },\n               \"type\":\"graph\",\n
    \              \"xaxis\":{  \n                  \"buckets\":null,\n                  \"mode\":\"time\",\n
    \                 \"name\":null,\n                  \"show\":true,\n                  \"values\":[
    \ \n\n                  ]\n               },\n               \"yaxes\":[  \n                  {
    \ \n                     \"decimals\":null,\n                     \"format\":\"short\",\n
    \                    \"label\":null,\n                     \"logBase\":1,\n                     \"max\":null,\n
    \                    \"min\":null,\n                     \"show\":true\n                  },\n
    \                 {  \n                     \"format\":\"short\",\n                     \"label\":null,\n
    \                    \"logBase\":1,\n                     \"max\":null,\n                     \"min\":null,\n
    \                    \"show\":false\n                  }\n               ]\n            },\n
    \           {  \n               \"aliasColors\":{  \n\n               },\n               \"bars\":false,\n
    \              \"dashLength\":10,\n               \"dashes\":false,\n               \"datasource\":\"prometheus\",\n
    \              \"fill\":1,\n               \"gridPos\":{  \n                  \"h\":9,\n
    \                 \"w\":12,\n                  \"x\":12,\n                  \"y\":13\n
    \              },\n               \"id\":6,\n               \"legend\":{  \n                  \"avg\":false,\n
    \                 \"current\":false,\n                  \"max\":false,\n                  \"min\":false,\n
    \                 \"show\":true,\n                  \"total\":false,\n                  \"values\":false\n
    \              },\n               \"lines\":true,\n               \"linewidth\":1,\n
    \              \"links\":[  \n\n               ],\n               \"nullPointMode\":\"null\",\n
    \              \"percentage\":false,\n               \"pointradius\":5,\n               \"points\":false,\n
    \              \"renderer\":\"flot\",\n               \"seriesOverrides\":[  \n\n
    \              ],\n               \"spaceLength\":10,\n               \"stack\":false,\n
    \              \"steppedLine\":false,\n               \"targets\":[  \n                  {
    \ \n                     \"expr\":\"sum(kube_pod_container_resource_requests_memory_bytes{namespace=\\\"$namespace\\\",
    pod=~\\\"$revision-deployment-.*\\\"})\",\n                     \"format\":\"time_series\",\n
    \                    \"interval\":\"\",\n                     \"intervalFactor\":1,\n
    \                    \"legendFormat\":\"Memory requested\",\n                     \"refId\":\"A\"\n
    \                 },\n                  {  \n                     \"expr\":\"sum(container_memory_usage_bytes{namespace=\\\"$namespace\\\",
    pod_name=~\\\"$revision-deployment-.*\\\"})\",\n                     \"format\":\"time_series\",\n
    \                    \"hide\":false,\n                     \"intervalFactor\":1,\n
    \                    \"legendFormat\":\"Memory used\",\n                     \"refId\":\"B\"\n
    \                 },\n                  {  \n                     \"expr\":\"sum(kube_pod_container_resource_limits_memory_bytes{namespace=\\\"$namespace\\\",
    pod=~\\\"$revision-deployment-.*\\\"})\",\n                     \"format\":\"time_series\",\n
    \                    \"intervalFactor\":1,\n                     \"refId\":\"C\"\n
    \                 }\n               ],\n               \"thresholds\":[  \n\n
    \              ],\n               \"timeFrom\":null,\n               \"timeShift\":null,\n
    \              \"title\":\"Pod Memory Usage\",\n               \"tooltip\":{  \n
    \                 \"shared\":true,\n                  \"sort\":2,\n                  \"value_type\":\"individual\"\n
    \              },\n               \"type\":\"graph\",\n               \"xaxis\":{
    \ \n                  \"buckets\":null,\n                  \"mode\":\"time\",\n
    \                 \"name\":null,\n                  \"show\":true,\n                  \"values\":[
    \ \n\n                  ]\n               },\n               \"yaxes\":[  \n                  {
    \ \n                     \"format\":\"decbytes\",\n                     \"label\":null,\n
    \                    \"logBase\":1,\n                     \"max\":null,\n                     \"min\":null,\n
    \                    \"show\":true\n                  },\n                  {
    \ \n                     \"format\":\"short\",\n                     \"label\":null,\n
    \                    \"logBase\":1,\n                     \"max\":null,\n                     \"min\":null,\n
    \                    \"show\":false\n                  }\n               ]\n            }\n
    \        ],\n         \"title\":\"Resource Usages\",\n         \"type\":\"row\"\n
    \     },\n      {  \n         \"collapsed\":true,\n         \"gridPos\":{  \n
    \           \"h\":1,\n            \"w\":24,\n            \"x\":0,\n            \"y\":2\n
    \        },\n         \"id\":16,\n         \"panels\":[  \n            {  \n               \"aliasColors\":{
    \ \n\n               },\n               \"bars\":false,\n               \"dashLength\":10,\n
    \              \"dashes\":false,\n               \"datasource\":\"prometheus\",\n
    \              \"fill\":1,\n               \"gridPos\":{  \n                  \"h\":10,\n
    \                 \"w\":24,\n                  \"x\":0,\n                  \"y\":3\n
    \              },\n               \"id\":10,\n               \"legend\":{  \n
    \                 \"avg\":false,\n                  \"current\":false,\n                  \"max\":false,\n
    \                 \"min\":false,\n                  \"show\":true,\n                  \"total\":false,\n
    \                 \"values\":false\n               },\n               \"lines\":true,\n
    \              \"linewidth\":1,\n               \"links\":[  \n\n               ],\n
    \              \"nullPointMode\":\"null\",\n               \"percentage\":false,\n
    \              \"pointradius\":5,\n               \"points\":false,\n               \"renderer\":\"flot\",\n
    \              \"seriesOverrides\":[  \n\n               ],\n               \"spaceLength\":10,\n
    \              \"stack\":false,\n               \"steppedLine\":true,\n               \"targets\":[
    \ \n                  {  \n                     \"expr\":\"sum(autoscaler_desired_pods{namespace_name=\\\"$namespace\\\",
    configuration_name=\\\"$configuration\\\", revision_name=\\\"$revision\\\"}) \",\n
    \                    \"format\":\"time_series\",\n                     \"intervalFactor\":1,\n
    \                    \"legendFormat\":\"Desired Pods\",\n                     \"refId\":\"A\"\n
    \                 },\n                  {  \n                     \"expr\":\"sum(autoscaler_observed_pods{namespace_name=\\\"$namespace\\\",
    configuration_name=\\\"$configuration\\\", revision_name=\\\"$revision\\\"})\",\n
    \                    \"format\":\"time_series\",\n                     \"intervalFactor\":1,\n
    \                    \"legendFormat\":\"Observed Pods\",\n                     \"refId\":\"B\"\n
    \                 }\n               ],\n               \"thresholds\":[  \n\n
    \              ],\n               \"timeFrom\":null,\n               \"timeShift\":null,\n
    \              \"title\":\"Pod Counts\",\n               \"tooltip\":{  \n                  \"shared\":true,\n
    \                 \"sort\":0,\n                  \"value_type\":\"individual\"\n
    \              },\n               \"type\":\"graph\",\n               \"xaxis\":{
    \ \n                  \"buckets\":null,\n                  \"mode\":\"time\",\n
    \                 \"name\":null,\n                  \"show\":true,\n                  \"values\":[
    \ \n\n                  ]\n               },\n               \"yaxes\":[  \n                  {
    \ \n                     \"format\":\"short\",\n                     \"label\":null,\n
    \                    \"logBase\":1,\n                     \"max\":null,\n                     \"min\":null,\n
    \                    \"show\":true\n                  },\n                  {
    \ \n                     \"format\":\"short\",\n                     \"label\":null,\n
    \                    \"logBase\":1,\n                     \"max\":null,\n                     \"min\":null,\n
    \                    \"show\":true\n                  }\n               ]\n            },\n
    \           {  \n               \"aliasColors\":{  \n\n               },\n               \"bars\":false,\n
    \              \"dashLength\":10,\n               \"dashes\":false,\n               \"datasource\":\"prometheus\",\n
    \              \"fill\":1,\n               \"gridPos\":{  \n                  \"h\":9,\n
    \                 \"w\":24,\n                  \"x\":0,\n                  \"y\":13\n
    \              },\n               \"id\":8,\n               \"legend\":{  \n                  \"avg\":false,\n
    \                 \"current\":false,\n                  \"max\":false,\n                  \"min\":false,\n
    \                 \"show\":true,\n                  \"total\":false,\n                  \"values\":false\n
    \              },\n               \"lines\":true,\n               \"linewidth\":1,\n
    \              \"links\":[  \n\n               ],\n               \"nullPointMode\":\"null\",\n
    \              \"percentage\":false,\n               \"pointradius\":5,\n               \"points\":false,\n
    \              \"renderer\":\"flot\",\n               \"seriesOverrides\":[  \n
    \                 {  \n                     \"alias\":\"Panic Mode\",\n                     \"color\":\"#ea6460\",\n
    \                    \"dashes\":true,\n                     \"fill\":2,\n                     \"linewidth\":2,\n
    \                    \"steppedLine\":true,\n                     \"yaxis\":2\n
    \                 },\n                  {  \n                     \"alias\":\"Target
    Concurrency Per Pod\",\n                     \"color\":\"#0a50a1\",\n                     \"dashes\":true,\n
    \                    \"steppedLine\":false\n                  }\n               ],\n
    \              \"spaceLength\":10,\n               \"stack\":false,\n               \"steppedLine\":true,\n
    \              \"targets\":[  \n                  {  \n                     \"expr\":\"sum(autoscaler_stable_request_concurrency{namespace_name=\\\"$namespace\\\",
    configuration_name=\\\"$configuration\\\", revision_name=\\\"$revision\\\"})\",\n
    \                    \"format\":\"time_series\",\n                     \"interval\":\"1s\",\n
    \                    \"intervalFactor\":1,\n                     \"legendFormat\":\"60
    Second Average Concurrency\",\n                     \"refId\":\"A\"\n                  },\n
    \                 {  \n                     \"expr\":\"sum(autoscaler_panic_request_concurrency{namespace_name=\\\"$namespace\\\",
    configuration_name=\\\"$configuration\\\", revision_name=\\\"$revision\\\"})\",\n
    \                    \"format\":\"time_series\",\n                     \"interval\":\"1s\",\n
    \                    \"intervalFactor\":1,\n                     \"legendFormat\":\"6
    Second Average Panic Concurrency\",\n                     \"refId\":\"B\"\n                  },\n
    \                 {  \n                     \"expr\":\"sum(autoscaler_target_concurrency_per_pod{namespace_name=\\\"$namespace\\\",
    configuration_name=\\\"$configuration\\\", revision_name=\\\"$revision\\\"})\",\n
    \                    \"format\":\"time_series\",\n                     \"intervalFactor\":1,\n
    \                    \"legendFormat\":\"60 Second Target Concurrency\",\n                     \"refId\":\"C\"\n
    \                 }\n               ],\n               \"thresholds\":[  \n\n
    \              ],\n               \"timeFrom\":null,\n               \"timeShift\":null,\n
    \              \"title\":\"Observed Concurrency\",\n               \"tooltip\":{
    \ \n                  \"shared\":true,\n                  \"sort\":0,\n                  \"value_type\":\"individual\"\n
    \              },\n               \"type\":\"graph\",\n               \"xaxis\":{
    \ \n                  \"buckets\":null,\n                  \"mode\":\"time\",\n
    \                 \"name\":null,\n                  \"show\":true,\n                  \"values\":[
    \ \n\n                  ]\n               },\n               \"yaxes\":[  \n                  {
    \ \n                     \"format\":\"short\",\n                     \"label\":\"\",\n
    \                    \"logBase\":1,\n                     \"max\":null,\n                     \"min\":null,\n
    \                    \"show\":true\n                  },\n                  {
    \ \n                     \"format\":\"short\",\n                     \"label\":\"\",\n
    \                    \"logBase\":1,\n                     \"max\":null,\n                     \"min\":null,\n
    \                    \"show\":false\n                  }\n               ]\n            },\n
    \           {  \n               \"aliasColors\":{  \n\n               },\n               \"bars\":false,\n
    \              \"dashLength\":10,\n               \"dashes\":false,\n               \"datasource\":\"prometheus\",\n
    \              \"decimals\":null,\n               \"fill\":1,\n               \"gridPos\":{
    \ \n                  \"h\":9,\n                  \"w\":24,\n                  \"x\":0,\n
    \                 \"y\":22\n               },\n               \"id\":12,\n               \"legend\":{
    \ \n                  \"avg\":false,\n                  \"current\":false,\n                  \"hideZero\":false,\n
    \                 \"max\":false,\n                  \"min\":false,\n                  \"show\":false,\n
    \                 \"total\":false,\n                  \"values\":false\n               },\n
    \              \"lines\":true,\n               \"linewidth\":1,\n               \"links\":[
    \ \n\n               ],\n               \"nullPointMode\":\"null\",\n               \"percentage\":false,\n
    \              \"pointradius\":5,\n               \"points\":false,\n               \"renderer\":\"flot\",\n
    \              \"seriesOverrides\":[  \n                  {  \n                     \"alias\":\"Panic
    Mode\",\n                     \"color\":\"#e24d42\",\n                     \"linewidth\":2,\n
    \                    \"yaxis\":2\n                  }\n               ],\n               \"spaceLength\":10,\n
    \              \"stack\":false,\n               \"steppedLine\":true,\n               \"targets\":[
    \ \n                  {  \n                     \"expr\":\"sum(autoscaler_panic_mode{namespace_name=\\\"$namespace\\\",
    configuration_name=\\\"$configuration\\\", revision_name=\\\"$revision\\\"} )\",\n
    \                    \"format\":\"time_series\",\n                     \"intervalFactor\":1,\n
    \                    \"legendFormat\":\"Panic Mode\",\n                     \"refId\":\"A\"\n
    \                 }\n               ],\n               \"thresholds\":[  \n\n
    \              ],\n               \"timeFrom\":null,\n               \"timeShift\":null,\n
    \              \"title\":\"Panic Mode\",\n               \"tooltip\":{  \n                  \"shared\":true,\n
    \                 \"sort\":0,\n                  \"value_type\":\"individual\"\n
    \              },\n               \"type\":\"graph\",\n               \"xaxis\":{
    \ \n                  \"buckets\":null,\n                  \"mode\":\"time\",\n
    \                 \"name\":null,\n                  \"show\":true,\n                  \"values\":[
    \ \n\n                  ]\n               },\n               \"yaxes\":[  \n                  {
    \ \n                     \"format\":\"short\",\n                     \"label\":null,\n
    \                    \"logBase\":1,\n                     \"max\":\"1.0\",\n                     \"min\":\"0\",\n
    \                    \"show\":true\n                  },\n                  {
    \ \n                     \"format\":\"short\",\n                     \"label\":null,\n
    \                    \"logBase\":1,\n                     \"max\":null,\n                     \"min\":null,\n
    \                    \"show\":false\n                  }\n               ]\n            }\n
    \        ],\n         \"title\":\"Autoscaler Metrics\",\n         \"type\":\"row\"\n
    \     },\n      {  \n         \"collapsed\":true,\n         \"gridPos\":{  \n
    \           \"h\":1,\n            \"w\":24,\n            \"x\":0,\n            \"y\":3\n
    \        },\n         \"id\":20,\n         \"panels\":[  \n                     {
    \ \n                        \"aliasColors\":{  \n\n                        },\n
    \                       \"bars\":false,\n                        \"dashLength\":10,\n
    \                       \"dashes\":false,\n                        \"datasource\":\"prometheus\",\n
    \                       \"fill\":1,\n                        \"gridPos\":{  \n
    \                          \"h\":10,\n                           \"w\":24,\n                           \"x\":0,\n
    \                          \"y\":12\n                        },\n                        \"id\":24,\n
    \                       \"legend\":{  \n                           \"avg\":false,\n
    \                          \"current\":false,\n                           \"max\":false,\n
    \                          \"min\":false,\n                           \"show\":true,\n
    \                          \"total\":false,\n                           \"values\":false\n
    \                       },\n                        \"lines\":true,\n                        \"linewidth\":1,\n
    \                       \"links\":[  \n\n                        ],\n                        \"nullPointMode\":\"null\",\n
    \                       \"percentage\":false,\n                        \"pointradius\":5,\n
    \                       \"points\":false,\n                        \"renderer\":\"flot\",\n
    \                       \"seriesOverrides\":[  \n\n                        ],\n
    \                       \"spaceLength\":10,\n                        \"stack\":false,\n
    \                       \"steppedLine\":false,\n                        \"targets\":[
    \ \n                           {  \n                              \"expr\":\"round(sum(increase(activator_request_count{namespace_name=\\\"$namespace\\\",
    configuration_name=~\\\"$configuration\\\",revision_name=~\\\"$revision\\\"}[1m]))
    by (response_code))\",\n                              \"format\":\"time_series\",\n
    \                             \"intervalFactor\":1,\n                              \"legendFormat\":\"{{
    response_code }}\",\n                              \"refId\":\"A\"\n                           }\n
    \                       ],\n                        \"thresholds\":[  \n\n                        ],\n
    \                       \"timeFrom\":null,\n                        \"timeShift\":null,\n
    \                       \"title\":\"Request Count in last minute by Response Code\",\n
    \                       \"tooltip\":{  \n                           \"shared\":true,\n
    \                          \"sort\":0,\n                           \"value_type\":\"individual\"\n
    \                       },\n                        \"type\":\"graph\",\n                        \"xaxis\":{
    \ \n                           \"buckets\":null,\n                           \"mode\":\"time\",\n
    \                          \"name\":null,\n                           \"show\":true,\n
    \                          \"values\":[  \n\n                           ]\n                        },\n
    \                       \"yaxes\":[  \n                           {  \n                              \"format\":\"none\",\n
    \                             \"label\":null,\n                              \"logBase\":1,\n
    \                             \"max\":null,\n                              \"min\":\"0\",\n
    \                             \"show\":true\n                           },\n                           {
    \ \n                              \"format\":\"short\",\n                              \"label\":null,\n
    \                             \"logBase\":1,\n                              \"max\":null,\n
    \                             \"min\":null,\n                              \"show\":true\n
    \                          }\n                        ]\n                     },\n
    \                    {  \n                        \"aliasColors\":{  \n\n                        },\n
    \                       \"bars\":false,\n                        \"dashLength\":10,\n
    \                       \"dashes\":false,\n                        \"datasource\":\"prometheus\",\n
    \                       \"fill\":1,\n                        \"gridPos\":{  \n
    \                          \"h\":10,\n                           \"w\":24,\n                           \"x\":0,\n
    \                          \"y\":32\n                        },\n                        \"id\":28,\n
    \                       \"legend\":{  \n                           \"avg\":true,\n
    \                          \"current\":false,\n                           \"max\":false,\n
    \                          \"min\":false,\n                           \"show\":true,\n
    \                          \"total\":false,\n                           \"values\":true\n
    \                       },\n                        \"lines\":true,\n                        \"linewidth\":1,\n
    \                       \"links\":[  \n\n                        ],\n                        \"nullPointMode\":\"null\",\n
    \                       \"percentage\":false,\n                        \"pointradius\":5,\n
    \                       \"points\":false,\n                        \"renderer\":\"flot\",\n
    \                       \"seriesOverrides\":[  \n\n                        ],\n
    \                       \"spaceLength\":10,\n                        \"stack\":false,\n
    \                       \"steppedLine\":false,\n                        \"targets\":[
    \ \n                           {  \n                              \"expr\":\"label_replace(histogram_quantile(0.50,
    sum(rate(activator_request_latencies_bucket{namespace_name=\\\"$namespace\\\",
    configuration_name=~\\\"$configuration\\\",revision_name=~\\\"$revision\\\"}[1m]))
    by (revision_name, le)), \\\"revision_name\\\", \\\"$2\\\", \\\"revision_name\\\",
    \\\"$configuration(-+)(.*)\\\")\",\n                              \"format\":\"time_series\",\n
    \                             \"intervalFactor\":1,\n                              \"legendFormat\":\"{{
    revision_name }} (p50)\",\n                              \"refId\":\"A\"\n                           },\n
    \                          {  \n                              \"expr\":\"label_replace(histogram_quantile(0.90,
    sum(rate(activator_request_latencies_bucket{namespace_name=\\\"$namespace\\\",
    configuration_name=~\\\"$configuration\\\",revision_name=~\\\"$revision\\\"}[1m]))
    by (revision_name, le)), \\\"revision_name\\\", \\\"$2\\\", \\\"revision_name\\\",
    \\\"$configuration(-+)(.*)\\\")\",\n                              \"format\":\"time_series\",\n
    \                             \"intervalFactor\":1,\n                              \"legendFormat\":\"{{
    revision_name }} (p90)\",\n                              \"refId\":\"B\"\n                           },\n
    \                          {  \n                              \"expr\":\"label_replace(histogram_quantile(0.95,
    sum(rate(activator_request_latencies_bucket{namespace_name=\\\"$namespace\\\",
    configuration_name=~\\\"$configuration\\\",revision_name=~\\\"$revision\\\"}[1m]))
    by (revision_name, le)), \\\"revision_name\\\", \\\"$2\\\", \\\"revision_name\\\",
    \\\"$configuration(-+)(.*)\\\")\",\n                              \"format\":\"time_series\",\n
    \                             \"intervalFactor\":1,\n                              \"legendFormat\":\"{{
    revision_name }} (p95)\",\n                              \"refId\":\"C\"\n                           },\n
    \                          {  \n                              \"expr\":\"label_replace(histogram_quantile(0.99,
    sum(rate(activator_request_latencies_bucket{namespace_name=\\\"$namespace\\\",
    configuration_name=~\\\"$configuration\\\",revision_name=~\\\"$revision\\\"}[1m]))
    by (revision_name, le)), \\\"revision_name\\\", \\\"$2\\\", \\\"revision_name\\\",
    \\\"$configuration(-+)(.*)\\\")\",\n                              \"format\":\"time_series\",\n
    \                             \"intervalFactor\":1,\n                              \"legendFormat\":\"{{
    revision_name }} (p99)\",\n                              \"refId\":\"D\"\n                           }\n
    \                       ],\n                        \"thresholds\":[  \n\n                        ],\n
    \                       \"timeFrom\":null,\n                        \"timeShift\":null,\n
    \                       \"title\":\"Response Time in last minute\",\n                        \"tooltip\":{
    \ \n                           \"shared\":true,\n                           \"sort\":0,\n
    \                          \"value_type\":\"individual\"\n                        },\n
    \                       \"type\":\"graph\",\n                        \"xaxis\":{
    \ \n                           \"buckets\":null,\n                           \"mode\":\"time\",\n
    \                          \"name\":null,\n                           \"show\":true,\n
    \                          \"values\":[  \n\n                           ]\n                        },\n
    \                       \"yaxes\":[  \n                           {  \n                              \"format\":\"ms\",\n
    \                             \"label\":null,\n                              \"logBase\":1,\n
    \                             \"max\":null,\n                              \"min\":null,\n
    \                             \"show\":true\n                           },\n                           {
    \ \n                              \"format\":\"short\",\n                              \"label\":null,\n
    \                             \"logBase\":1,\n                              \"max\":null,\n
    \                             \"min\":null,\n                              \"show\":true\n
    \                          }\n                         ]\n                       }
    \    \n         ],\n         \"title\":\"Activator Metrics\",\n         \"type\":\"row\"\n
    \     }\n   ],\n   \"refresh\":false,\n   \"schemaVersion\":16,\n   \"style\":\"dark\",\n
    \  \"tags\":[  \n\n   ],\n   \"templating\":{  \n      \"list\":[  \n         {
    \ \n            \"allValue\":null,\n            \"current\":{  \n\n            },\n
    \           \"datasource\":\"prometheus\",\n            \"hide\":0,\n            \"includeAll\":false,\n
    \           \"label\":\"Namespace\",\n            \"multi\":false,\n            \"name\":\"namespace\",\n
    \           \"options\":[  \n\n            ],\n            \"query\":\"label_values(autoscaler_desired_pods,
    namespace_name)\",\n            \"refresh\":1,\n            \"regex\":\"\",\n
    \           \"sort\":1,\n            \"tagValuesQuery\":\"\",\n            \"tags\":[
    \ \n\n            ],\n            \"tagsQuery\":\"\",\n            \"type\":\"query\",\n
    \           \"useTags\":false\n         },\n         {  \n            \"allValue\":null,\n
    \           \"current\":{  \n\n            },\n            \"datasource\":\"prometheus\",\n
    \           \"hide\":0,\n            \"includeAll\":false,\n            \"label\":\"Configuration\",\n
    \           \"multi\":false,\n            \"name\":\"configuration\",\n            \"options\":[
    \ \n\n            ],\n            \"query\":\"label_values(autoscaler_desired_pods{namespace_name=\\\"$namespace\\\"},
    configuration_name)\",\n            \"refresh\":1,\n            \"regex\":\"\",\n
    \           \"sort\":1,\n            \"tagValuesQuery\":\"\",\n            \"tags\":[
    \ \n\n            ],\n            \"tagsQuery\":\"\",\n            \"type\":\"query\",\n
    \           \"useTags\":false\n         },\n         {  \n            \"allValue\":null,\n
    \           \"current\":{  \n\n            },\n            \"datasource\":\"prometheus\",\n
    \           \"hide\":0,\n            \"includeAll\":false,\n            \"label\":\"Revision\",\n
    \           \"multi\":false,\n            \"name\":\"revision\",\n            \"options\":[
    \ \n\n            ],\n            \"query\":\"label_values(autoscaler_desired_pods{namespace_name=\\\"$namespace\\\",
    configuration_name=\\\"$configuration\\\"}, revision_name)\",\n            \"refresh\":1,\n
    \           \"regex\":\"\",\n            \"sort\":2,\n            \"tagValuesQuery\":\"\",\n
    \           \"tags\":[  \n\n            ],\n            \"tagsQuery\":\"\",\n
    \           \"type\":\"query\",\n            \"useTags\":false\n         }\n      ]\n
    \  },\n   \"time\":{  \n      \"from\":\"now-15m\",\n      \"to\":\"now\"\n   },\n
    \  \"timepicker\":{  \n      \"refresh_intervals\":[  \n         \"5s\",\n         \"10s\",\n
    \        \"30s\",\n         \"1m\",\n         \"5m\",\n         \"15m\",\n         \"30m\",\n
    \        \"1h\",\n         \"2h\",\n         \"1d\"\n      ],\n      \"time_options\":[
    \ \n         \"5m\",\n         \"15m\",\n         \"1h\",\n         \"6h\",\n
    \        \"12h\",\n         \"24h\",\n         \"2d\",\n         \"7d\",\n         \"30d\"\n
    \     ]\n   },\n   \"timezone\":\"\",\n   \"title\":\"Knative Serving - Scaling
    Debugging\",\n   \"uid\":\"u_-9SIMiz\",\n   \"version\":2\n}\n"
kind: ConfigMap
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: scaling-config
  namespace: knative-monitoring
---

apiVersion: v1
data:
  resource-dashboard.json: |
    {
      "__inputs": [
        {
          "description": "",
          "label": "prometheus",
          "name": "prometheus",
          "pluginId": "prometheus",
          "pluginName": "Prometheus",
          "type": "datasource"
        }
      ],
      "annotations": {
        "list": []
      },
      "description": "Knative Serving - Revision CPU and Memory Usage",
      "editable": false,
      "gnetId": null,
      "graphTooltip": 0,
      "links": [],
      "panels": [
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "prometheus",
          "fill": 1,
          "gridPos": {
            "h": 9,
            "w": 12,
            "x": 0,
            "y": 0
          },
          "id": 2,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(rate(container_cpu_usage_seconds_total{namespace=\"$namespace\", pod_name=~\"$revision.*\", container_name != \"POD\"}[1m])) by (container_name)",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "{{container_name}}",
              "refId": "A"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Total CPU Usage",
          "tooltip": {
            "shared": true,
            "sort": 2,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "s",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": false
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "prometheus",
          "fill": 1,
          "gridPos": {
            "h": 9,
            "w": 12,
            "x": 12,
            "y": 0
          },
          "id": 3,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(container_memory_usage_bytes{namespace=\"$namespace\", pod_name=~\"$revision.*\", container_name != \"POD\"}) by (container_name)",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "{{container_name}}",
              "refId": "A"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Total Memory Usage",
          "tooltip": {
            "shared": true,
            "sort": 2,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "decbytes",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": false
            }
          ]
        }
      ],
      "refresh": "5s",
      "schemaVersion": 16,
      "style": "dark",
      "tags": [],
      "templating": {
        "list": [
          {
            "allValue": null,
            "current": {},
            "datasource": "prometheus",
            "hide": 0,
            "includeAll": false,
            "label": "Namespace",
            "multi": false,
            "name": "namespace",
            "options": [],
            "query": "label_values(kube_pod_labels{label_serving_knative_dev_configuration=~\".+\"}, namespace)",
            "refresh": 2,
            "regex": "",
            "sort": 1,
            "tagValuesQuery": "",
            "tags": [],
            "tagsQuery": "",
            "type": "query",
            "useTags": false
          },
          {
            "allValue": null,
            "current": {},
            "datasource": "prometheus",
            "hide": 0,
            "includeAll": false,
            "label": "Configuration",
            "multi": false,
            "name": "configuration",
            "options": [],
            "query": "label_values(kube_pod_labels{label_serving_knative_dev_configuration=~\".+\", namespace=\"$namespace\"}, label_serving_knative_dev_configuration)",
            "refresh": 2,
            "regex": "",
            "sort": 1,
            "tagValuesQuery": "",
            "tags": [],
            "tagsQuery": "",
            "type": "query",
            "useTags": false
          },
          {
            "allValue": null,
            "current": {},
            "datasource": "prometheus",
            "hide": 0,
            "includeAll": false,
            "label": "Revision",
            "multi": false,
            "name": "revision",
            "options": [],
            "query": "label_values(kube_pod_labels{label_serving_knative_dev_configuration=~\".+\", namespace=\"$namespace\", label_serving_knative_dev_configuration=\"$configuration\"}, label_serving_knative_dev_revision)",
            "refresh": 2,
            "regex": "",
            "sort": 2,
            "tagValuesQuery": "",
            "tags": [],
            "tagsQuery": "",
            "type": "query",
            "useTags": false
          }
        ]
      },
      "time": {
        "from": "now-15m",
        "to": "now"
      },
      "timepicker": {
        "refresh_intervals": [
          "5s",
          "10s",
          "30s",
          "1m",
          "5m",
          "15m",
          "30m",
          "1h",
          "2h",
          "1d"
        ],
        "time_options": [
          "5m",
          "15m",
          "1h",
          "6h",
          "12h",
          "24h",
          "2d",
          "7d",
          "30d"
        ]
      },
      "timezone": "",
      "title": "Knative Serving - Revision CPU and Memory Usage",
      "uid": "bKOoE9Wmk",
      "version": 4
    }
  revision-dashboard.json: |
    {
      "__inputs": [
        {
          "description": "",
          "label": "prometheus",
          "name": "prometheus",
          "pluginId": "prometheus",
          "pluginName": "Prometheus",
          "type": "datasource"
        }
      ],
      "annotations": {
        "list": []
      },
      "description": "Knative Serving - Revision HTTP Requests",
      "editable": false,
      "gnetId": null,
      "graphTooltip": 0,
      "links": [],
      "panels": [
        {
          "collapsed": false,
          "gridPos": {
            "h": 1,
            "w": 24,
            "x": 0,
            "y": 0
          },
          "id": 8,
          "panels": [],
          "title": "Overview (average over the selected time range)",
          "type": "row"
        },
        {
          "cacheTimeout": null,
          "colorBackground": false,
          "colorValue": false,
          "colors": [
            "#299c46",
            "rgba(237, 129, 40, 0.89)",
            "#d44a3a"
          ],
          "datasource": "prometheus",
          "format": "ops",
          "gauge": {
            "maxValue": 100,
            "minValue": 0,
            "show": false,
            "thresholdLabels": false,
            "thresholdMarkers": true
          },
          "gridPos": {
            "h": 4,
            "w": 6,
            "x": 0,
            "y": 1
          },
          "id": 2,
          "interval": null,
          "links": [],
          "mappingType": 1,
          "mappingTypes": [
            {
              "name": "value to text",
              "value": 1
            },
            {
              "name": "range to text",
              "value": 2
            }
          ],
          "maxDataPoints": 100,
          "nullPointMode": "connected",
          "nullText": null,
          "postfix": "",
          "postfixFontSize": "50%",
          "prefix": "",
          "prefixFontSize": "50%",
          "rangeMaps": [
            {
              "from": "null",
              "text": "N/A",
              "to": "null"
            }
          ],
          "repeat": null,
          "repeatDirection": "v",
          "sparkline": {
            "fillColor": "rgba(31, 118, 189, 0.18)",
            "full": true,
            "lineColor": "rgb(31, 120, 193)",
            "show": true
          },
          "tableColumn": "",
          "targets": [
            {
              "expr": "round(sum(rate(revision_request_count{namespace_name=\"$namespace\", revision_name=~\"$revision\", configuration_name=~\"$configuration\"}[1m])), 0.001)",
              "format": "time_series",
              "hide": false,
              "interval": "",
              "intervalFactor": 1,
              "refId": "A"
            }
          ],
          "thresholds": "",
          "title": "Request Volume",
          "type": "singlestat",
          "valueFontSize": "80%",
          "valueMaps": [
            {
              "op": "=",
              "text": "N/A",
              "value": "null"
            }
          ],
          "valueName": "avg"
        },
        {
          "cacheTimeout": null,
          "colorBackground": false,
          "colorValue": false,
          "colors": [
            "#d44a3a",
            "rgba(237, 129, 40, 0.89)",
            "#299c46"
          ],
          "datasource": "prometheus",
          "format": "percentunit",
          "gauge": {
            "maxValue": 100,
            "minValue": 0,
            "show": false,
            "thresholdLabels": false,
            "thresholdMarkers": true
          },
          "gridPos": {
            "h": 4,
            "w": 6,
            "x": 6,
            "y": 1
          },
          "id": 4,
          "interval": null,
          "links": [],
          "mappingType": 1,
          "mappingTypes": [
            {
              "name": "value to text",
              "value": 1
            },
            {
              "name": "range to text",
              "value": 2
            }
          ],
          "maxDataPoints": 100,
          "nullPointMode": "connected",
          "nullText": null,
          "postfix": "",
          "postfixFontSize": "50%",
          "prefix": "",
          "prefixFontSize": "50%",
          "rangeMaps": [
            {
              "from": "null",
              "text": "N/A",
              "to": "null"
            }
          ],
          "sparkline": {
            "fillColor": "rgba(31, 118, 189, 0.18)",
            "full": true,
            "lineColor": "rgb(31, 120, 193)",
            "show": true
          },
          "tableColumn": "",
          "targets": [
            {
              "expr": "sum(rate(revision_request_count{response_code_class!=\"5xx\", namespace_name=\"$namespace\", revision_name=~\"$revision\", configuration_name=~\"$configuration\"}[1m])) / sum(rate(revision_request_count{namespace_name=\"$namespace\", revision_name=~\"$revision\", configuration_name=~\"$configuration\"}[1m]))",
              "format": "time_series",
              "interval": "",
              "intervalFactor": 1,
              "refId": "A"
            }
          ],
          "thresholds": "95, 99, 99.5",
          "title": "Success Rate (non-5xx responses)",
          "type": "singlestat",
          "valueFontSize": "80%",
          "valueMaps": [
            {
              "op": "=",
              "text": "N/A",
              "value": "null"
            }
          ],
          "valueName": "avg"
        },
        {
          "cacheTimeout": null,
          "colorBackground": false,
          "colorValue": false,
          "colors": [
            "#d44a3a",
            "rgba(237, 129, 40, 0.89)",
            "#299c46"
          ],
          "datasource": "prometheus",
          "format": "ops",
          "gauge": {
            "maxValue": 100,
            "minValue": 0,
            "show": false,
            "thresholdLabels": false,
            "thresholdMarkers": true
          },
          "gridPos": {
            "h": 4,
            "w": 6,
            "x": 12,
            "y": 1
          },
          "id": 5,
          "interval": null,
          "links": [],
          "mappingType": 1,
          "mappingTypes": [
            {
              "name": "value to text",
              "value": 1
            },
            {
              "name": "range to text",
              "value": 2
            }
          ],
          "maxDataPoints": 100,
          "nullPointMode": "connected",
          "nullText": null,
          "postfix": "",
          "postfixFontSize": "50%",
          "prefix": "",
          "prefixFontSize": "50%",
          "rangeMaps": [
            {
              "from": "null",
              "text": "N/A",
              "to": "null"
            }
          ],
          "sparkline": {
            "fillColor": "rgba(31, 118, 189, 0.18)",
            "full": true,
            "lineColor": "rgb(31, 120, 193)",
            "show": true
          },
          "tableColumn": "",
          "targets": [
            {
              "expr": "sum(rate(revision_request_count{response_code_class=\"4xx\", namespace_name=\"$namespace\", revision_name=~\"$revision\", configuration_name=~\"$configuration\"}[1m])) ",
              "format": "time_series",
              "interval": "",
              "intervalFactor": 1,
              "refId": "A"
            }
          ],
          "thresholds": "",
          "title": "4xx",
          "type": "singlestat",
          "valueFontSize": "80%",
          "valueMaps": [
            {
              "op": "=",
              "text": "N/A",
              "value": "null"
            }
          ],
          "valueName": "avg"
        },
        {
          "cacheTimeout": null,
          "colorBackground": false,
          "colorValue": false,
          "colors": [
            "#d44a3a",
            "rgba(237, 129, 40, 0.89)",
            "#299c46"
          ],
          "datasource": "prometheus",
          "format": "ops",
          "gauge": {
            "maxValue": 100,
            "minValue": 0,
            "show": false,
            "thresholdLabels": false,
            "thresholdMarkers": true
          },
          "gridPos": {
            "h": 4,
            "w": 6,
            "x": 18,
            "y": 1
          },
          "id": 9,
          "interval": null,
          "links": [],
          "mappingType": 1,
          "mappingTypes": [
            {
              "name": "value to text",
              "value": 1
            },
            {
              "name": "range to text",
              "value": 2
            }
          ],
          "maxDataPoints": 100,
          "nullPointMode": "connected",
          "nullText": null,
          "postfix": "",
          "postfixFontSize": "50%",
          "prefix": "",
          "prefixFontSize": "50%",
          "rangeMaps": [
            {
              "from": "null",
              "text": "N/A",
              "to": "null"
            }
          ],
          "sparkline": {
            "fillColor": "rgba(31, 118, 189, 0.18)",
            "full": true,
            "lineColor": "rgb(31, 120, 193)",
            "show": true
          },
          "tableColumn": "",
          "targets": [
            {
              "expr": "sum(rate(revision_request_count{response_code_class=\"5xx\", namespace_name=\"$namespace\", revision_name=~\"$revision\", configuration_name=~\"$configuration\"}[1m])) ",
              "format": "time_series",
              "interval": "",
              "intervalFactor": 1,
              "refId": "A"
            }
          ],
          "thresholds": "",
          "title": "5xx",
          "type": "singlestat",
          "valueFontSize": "80%",
          "valueMaps": [
            {
              "op": "=",
              "text": "N/A",
              "value": "null"
            }
          ],
          "valueName": "avg"
        },
        {
          "collapsed": false,
          "gridPos": {
            "h": 1,
            "w": 24,
            "x": 0,
            "y": 5
          },
          "id": 11,
          "panels": [],
          "title": "Request Volume",
          "type": "row"
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "prometheus",
          "fill": 1,
          "gridPos": {
            "h": 10,
            "w": 12,
            "x": 0,
            "y": 6
          },
          "id": 17,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "label_replace(round(sum(rate(revision_request_count{namespace_name=\"$namespace\", configuration_name=~\"$configuration\",revision_name=~\"$revision\"}[1m])) by (revision_name), 0.001), \"revision_name\", \"$2\", \"revision_name\", \"$configuration(-+)(.*)\")",
              "format": "time_series",
              "interval": "",
              "intervalFactor": 1,
              "legendFormat": "{{revision_name}}",
              "refId": "A"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Request Volume by Revision",
          "tooltip": {
            "shared": true,
            "sort": 2,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "ops",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "prometheus",
          "fill": 1,
          "gridPos": {
            "h": 10,
            "w": 12,
            "x": 12,
            "y": 6
          },
          "id": 18,
          "legend": {
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "show": true,
            "total": false,
            "values": false
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "round(sum(rate(revision_request_count{namespace_name=\"$namespace\", configuration_name=~\"$configuration\",revision_name=~\"$revision\"}[1m])) by (response_code_class), 0.001)",
              "format": "time_series",
              "interval": "",
              "intervalFactor": 1,
              "legendFormat": "{{ response_code_class }}",
              "refId": "A"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Request Volume by Response Code Class",
          "tooltip": {
            "shared": true,
            "sort": 2,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "ops",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            }
          ]
        },
        {
          "collapsed": false,
          "gridPos": {
            "h": 1,
            "w": 24,
            "x": 0,
            "y": 16
          },
          "id": 15,
          "panels": [],
          "title": "Response Time",
          "type": "row"
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "prometheus",
          "fill": 1,
          "gridPos": {
            "h": 9,
            "w": 12,
            "x": 0,
            "y": 17
          },
          "id": 20,
          "legend": {
            "alignAsTable": false,
            "avg": true,
            "current": false,
            "hideEmpty": false,
            "hideZero": false,
            "max": false,
            "min": false,
            "rightSide": false,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "label_replace(histogram_quantile(0.50, sum(rate(revision_request_latencies_bucket{namespace_name=\"$namespace\", configuration_name=~\"$configuration\",revision_name=~\"$revision\"}[1m])) by (revision_name, le)), \"revision_name\", \"$2\", \"revision_name\", \"$configuration(-+)(.*)\")",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "{{ revision_name }} (p50)",
              "refId": "A"
            },
            {
              "expr": "label_replace(histogram_quantile(0.90, sum(rate(revision_request_latencies_bucket{namespace_name=\"$namespace\", configuration_name=~\"$configuration\",revision_name=~\"$revision\"}[1m])) by (revision_name, le)), \"revision_name\", \"$2\", \"revision_name\", \"$configuration(-+)(.*)\")",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "{{ revision_name }} (p90)",
              "refId": "B"
            },
            {
              "expr": "label_replace(histogram_quantile(0.95, sum(rate(revision_request_latencies_bucket{namespace_name=\"$namespace\", configuration_name=~\"$configuration\",revision_name=~\"$revision\"}[1m])) by (revision_name, le)), \"revision_name\", \"$2\", \"revision_name\", \"$configuration(-+)(.*)\")",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "{{ revision_name }} (p95)",
              "refId": "C"
            },
            {
              "expr": "label_replace(histogram_quantile(0.99, sum(rate(revision_request_latencies_bucket{namespace_name=\"$namespace\", configuration_name=~\"$configuration\",revision_name=~\"$revision\"}[1m])) by (revision_name, le)), \"revision_name\", \"$2\", \"revision_name\", \"$configuration(-+)(.*)\")",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "{{ revision_name }} (p99)",
              "refId": "D"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Response Time by Revision",
          "tooltip": {
            "shared": true,
            "sort": 2,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "s",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": false
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": false,
          "dashLength": 10,
          "dashes": false,
          "datasource": "prometheus",
          "fill": 1,
          "gridPos": {
            "h": 9,
            "w": 12,
            "x": 12,
            "y": 17
          },
          "id": 21,
          "legend": {
            "alignAsTable": false,
            "avg": true,
            "current": false,
            "hideEmpty": false,
            "hideZero": false,
            "max": false,
            "min": false,
            "rightSide": false,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "spaceLength": 10,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "label_replace(histogram_quantile(0.50, sum(rate(revision_request_latencies_bucket{namespace_name=\"$namespace\", configuration_name=~\"$configuration\",revision_name=~\"$revision\",response_code_class=\"2xx\"}[1m])) by (le)), \"revision_name\", \"$2\", \"revision_name\", \"$configuration(-+)(.*)\")",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "2xx (p50)",
              "refId": "C"
            },
            {
              "expr": "label_replace(histogram_quantile(0.50, sum(rate(revision_request_latencies_bucket{namespace_name=\"$namespace\", configuration_name=~\"$configuration\",revision_name=~\"$revision\",response_code_class=\"3xx\"}[1m])) by (le)), \"revision_name\", \"$2\", \"revision_name\", \"$configuration(-+)(.*)\")",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "3xx (p50)",
              "refId": "D"
            },
            {
              "expr": "label_replace(histogram_quantile(0.50, sum(rate(revision_request_latencies_bucket{namespace_name=\"$namespace\", configuration_name=~\"$configuration\",revision_name=~\"$revision\",response_code_class=\"4xx\"}[1m])) by (le)), \"revision_name\", \"$2\", \"revision_name\", \"$configuration(-+)(.*)\")",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "4xx (p50)",
              "refId": "A"
            },
            {
              "expr": "label_replace(histogram_quantile(0.50, sum(rate(revision_request_latencies_bucket{namespace_name=\"$namespace\", configuration_name=~\"$configuration\",revision_name=~\"$revision\",response_code_class=\"5xx\"}[1m])) by (le)), \"revision_name\", \"$2\", \"revision_name\", \"$configuration(-+)(.*)\")",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "5xx (p50)",
              "refId": "B"
            },
            {
              "expr": "label_replace(histogram_quantile(0.95, sum(rate(revision_request_latencies_bucket{namespace_name=\"$namespace\", configuration_name=~\"$configuration\",revision_name=~\"$revision\",response_code_class=\"2xx\"}[1m])) by (le)), \"revision_name\", \"$2\", \"revision_name\", \"$configuration(-+)(.*)\")",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "2xx (p95)",
              "refId": "E"
            },
            {
              "expr": "label_replace(histogram_quantile(0.95, sum(rate(revision_request_latencies_bucket{namespace_name=\"$namespace\", configuration_name=~\"$configuration\",revision_name=~\"$revision\",response_code_class=\"3xx\"}[1m])) by (le)), \"revision_name\", \"$2\", \"revision_name\", \"$configuration(-+)(.*)\")",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "3xx (p95)",
              "refId": "F"
            },
            {
              "expr": "label_replace(histogram_quantile(0.95, sum(rate(revision_request_latencies_bucket{namespace_name=\"$namespace\", configuration_name=~\"$configuration\",revision_name=~\"$revision\",response_code_class=\"4xx\"}[1m])) by (le)), \"revision_name\", \"$2\", \"revision_name\", \"$configuration(-+)(.*)\")",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "4xx (p95)",
              "refId": "G"
            },
            {
              "expr": "label_replace(histogram_quantile(0.95, sum(rate(revision_request_latencies_bucket{namespace_name=\"$namespace\", configuration_name=~\"$configuration\",revision_name=~\"$revision\",response_code_class=\"5xx\"}[1m])) by (le)), \"revision_name\", \"$2\", \"revision_name\", \"$configuration(-+)(.*)\")",
              "format": "time_series",
              "intervalFactor": 1,
              "legendFormat": "5xx (p95)",
              "refId": "H"
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Response Time by Response Code Class",
          "tooltip": {
            "shared": true,
            "sort": 2,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "buckets": null,
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "s",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": false
            }
          ]
        }
      ],
      "refresh": "5s",
      "schemaVersion": 16,
      "style": "dark",
      "tags": [],
      "templating": {
        "list": [
          {
            "allValue": null,
            "current": {},
            "datasource": "prometheus",
            "hide": 0,
            "includeAll": false,
            "label": "Namespace",
            "multi": false,
            "name": "namespace",
            "options": [],
            "query": "label_values(revision_request_count{namespace_name!=\"unknown\"}, namespace_name)",
            "refresh": 1,
            "regex": "",
            "sort": 1,
            "tagValuesQuery": "",
            "tags": [],
            "tagsQuery": "",
            "type": "query",
            "useTags": false
          },
          {
            "allValue": null,
            "current": {},
            "datasource": "prometheus",
            "hide": 0,
            "includeAll": false,
            "label": "Configuration",
            "multi": false,
            "name": "configuration",
            "options": [],
            "query": "label_values(revision_request_count{namespace_name=\"$namespace\", configuration_name!=\"unknown\"}, configuration_name)",
            "refresh": 1,
            "regex": "",
            "sort": 1,
            "tagValuesQuery": "$tag",
            "tags": [],
            "tagsQuery": "",
            "type": "query",
            "useTags": false
          },
          {
            "allValue": null,
            "current": {},
            "datasource": "prometheus",
            "hide": 0,
            "includeAll": true,
            "label": "Revision",
            "multi": true,
            "name": "revision",
            "options": [],
            "query": "label_values(revision_request_count{namespace_name=\"$namespace\", configuration_name=~\"$configuration\", revision_name!=\"unknown\"}, revision_name)",
            "refresh": 1,
            "regex": "",
            "sort": 2,
            "tagValuesQuery": "",
            "tags": [],
            "tagsQuery": "",
            "type": "query",
            "useTags": false
          }
        ]
      },
      "time": {
        "from": "now-15m",
        "to": "now"
      },
      "timepicker": {
        "refresh_intervals": [
          "5s",
          "10s",
          "30s",
          "1m",
          "5m",
          "15m",
          "30m",
          "1h",
          "2h",
          "1d"
        ],
        "time_options": [
          "5m",
          "15m",
          "1h",
          "6h",
          "12h",
          "24h",
          "2d",
          "7d",
          "30d"
        ]
      },
      "timezone": "",
      "title": "Knative Serving - Revision HTTP Requests",
      "uid": "im_gFbWik",
      "version": 2
    }
kind: ConfigMap
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: grafana-dashboard-definition-knative
  namespace: knative-monitoring
---

apiVersion: v1
data:
  prometheus.yaml: |
    datasources:
     - name: prometheus
       type: prometheus
       access: proxy
       org_id: 1
       url: http://prometheus-system-np:8080
       version: 1
       editable: false
kind: ConfigMap
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: grafana-datasources
  namespace: knative-monitoring---

apiVersion: v1
data:
  dashboards.yaml: |
    - name: 'knative'
      org_id: 1
      folder: ''
      type: file
      options:
        folder: /grafana-dashboard-definition/knative
    - name: 'knative-efficiency'
      org_id: 1
      folder: ''
      type: file
      options:
        folder: /grafana-dashboard-definition/knative-efficiency
    - name: 'knative-reconciler'
      org_id: 1
      folder: ''
      type: file
      options:
        folder: /grafana-dashboard-definition/knative-reconciler
    - name: 'istio'
      org_id: 1
      folder: ''
      type: file
      options:
        folder: /grafana-dashboard-definition/istio
    - name: 'mixer'
      org_id: 1
      folder: ''
      type: file
      options:
        folder: /grafana-dashboard-definition/mixer
    - name: 'pilot'
      org_id: 1
      folder: ''
      type: file
      options:
        folder: /grafana-dashboard-definition/pilot
    - name: 'kubernetes-deployment'
      org_id: 1
      folder: ''
      type: file
      options:
        folder: /grafana-dashboard-definition/kubernetes-deployment
    - name: 'kubernetes-capacity-planning'
      org_id: 1
      folder: ''
      type: file
      options:
        folder: /grafana-dashboard-definition/kubernetes-capacity-planning
    - name: 'kubernetes-cluster-health'
      org_id: 1
      folder: ''
      type: file
      options:
        folder: /grafana-dashboard-definition/kubernetes-cluster-health
    - name: 'kubernetes-cluster-status'
      org_id: 1
      folder: ''
      type: file
      options:
        folder: /grafana-dashboard-definition/kubernetes-cluster-status
    - name: 'kubernetes-control-plane-status'
      org_id: 1
      folder: ''
      type: file
      options:
        folder: /grafana-dashboard-definition/kubernetes-control-plane-status
    - name: 'kubernetes-resource-requests'
      org_id: 1
      folder: ''
      type: file
      options:
        folder: /grafana-dashboard-definition/kubernetes-resource-requests
    - name: 'kubernetes-nodes'
      org_id: 1
      folder: ''
      type: file
      options:
        folder: /grafana-dashboard-definition/kubernetes-nodes
    - name: 'kubernetes-pods'
      org_id: 1
      folder: ''
      type: file
      options:
        folder: /grafana-dashboard-definition/kubernetes-pods
    - name: 'kubernetes-statefulset'
      org_id: 1
      folder: ''
      type: file
      options:
        folder: /grafana-dashboard-definition/kubernetes-statefulset
    - name: 'knative-serving-scaling'
      org_id: 1
      folder: ''
      type: file
      options:
        folder: /grafana-dashboard-definition/scaling
kind: ConfigMap
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: grafana-dashboards
  namespace: knative-monitoring---

apiVersion: v1
kind: Service
metadata:
  labels:
    app: grafana
    serving.knative.dev/release: "v0.7.0"
  name: grafana
  namespace: knative-monitoring
spec:
  ports:
  - port: 30802
    protocol: TCP
    targetPort: 3000
  selector:
    app: grafana
  type: NodePort---

apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: grafana
  namespace: knative-monitoring
spec:
  replicas: 1
  selector:
    matchLabels:
      app: grafana
  template:
    metadata:
      labels:
        app: grafana
        serving.knative.dev/release: "v0.7.0"
    spec:
      containers:
      - image: quay.io/coreos/monitoring-grafana:5.0.3
        name: grafana
        ports:
        - containerPort: 3000
          name: web
        resources:
          limits:
            cpu: 200m
            memory: 200Mi
          requests:
            cpu: 100m
            memory: 100Mi
        volumeMounts:
        - mountPath: /data
          name: grafana-storage
        - mountPath: /grafana/conf/provisioning/datasources
          name: grafana-datasources
        - mountPath: /grafana/conf/provisioning/dashboards
          name: grafana-dashboards
        - mountPath: /grafana-dashboard-definition/knative
          name: grafana-dashboard-definition-knative
        - mountPath: /grafana-dashboard-definition/knative-efficiency
          name: grafana-dashboard-definition-knative-efficiency
        - mountPath: /grafana-dashboard-definition/knative-reconciler
          name: grafana-dashboard-definition-knative-reconciler
        - mountPath: /grafana-dashboard-definition/kubernetes-deployment
          name: grafana-dashboard-definition-kubernetes-deployment
        - mountPath: /grafana-dashboard-definition/kubernetes-capacity-planning
          name: grafana-dashboard-definition-kubernetes-capacity-planning
        - mountPath: /grafana-dashboard-definition/kubernetes-cluster-health
          name: grafana-dashboard-definition-kubernetes-cluster-health
        - mountPath: /grafana-dashboard-definition/kubernetes-cluster-status
          name: grafana-dashboard-definition-kubernetes-cluster-status
        - mountPath: /grafana-dashboard-definition/kubernetes-control-plane-status
          name: grafana-dashboard-definition-kubernetes-control-plane-status
        - mountPath: /grafana-dashboard-definition/kubernetes-resource-requests
          name: grafana-dashboard-definition-kubernetes-resource-requests
        - mountPath: /grafana-dashboard-definition/kubernetes-nodes
          name: grafana-dashboard-definition-kubernetes-nodes
        - mountPath: /grafana-dashboard-definition/kubernetes-pods
          name: grafana-dashboard-definition-kubernetes-pods
        - mountPath: /grafana-dashboard-definition/kubernetes-statefulset
          name: grafana-dashboard-definition-kubernetes-statefulset
        - mountPath: /grafana-dashboard-definition/scaling
          name: scaling-config
      securityContext:
        runAsNonRoot: true
        runAsUser: 65534
      volumes:
      - emptyDir: {}
        name: grafana-storage
      - configMap:
          name: grafana-datasources
        name: grafana-datasources
      - configMap:
          name: grafana-dashboards
        name: grafana-dashboards
      - configMap:
          name: grafana-dashboard-definition-knative
        name: grafana-dashboard-definition-knative
      - configMap:
          name: grafana-dashboard-definition-knative-efficiency
        name: grafana-dashboard-definition-knative-efficiency
      - configMap:
          name: grafana-dashboard-definition-knative-reconciler
        name: grafana-dashboard-definition-knative-reconciler
      - configMap:
          name: grafana-dashboard-definition-kubernetes-deployment
        name: grafana-dashboard-definition-kubernetes-deployment
      - configMap:
          name: grafana-dashboard-definition-kubernetes-capacity-planning
        name: grafana-dashboard-definition-kubernetes-capacity-planning
      - configMap:
          name: grafana-dashboard-definition-kubernetes-cluster-health
        name: grafana-dashboard-definition-kubernetes-cluster-health
      - configMap:
          name: grafana-dashboard-definition-kubernetes-cluster-status
        name: grafana-dashboard-definition-kubernetes-cluster-status
      - configMap:
          name: grafana-dashboard-definition-kubernetes-control-plane-status
        name: grafana-dashboard-definition-kubernetes-control-plane-status
      - configMap:
          name: grafana-dashboard-definition-kubernetes-resource-requests
        name: grafana-dashboard-definition-kubernetes-resource-requests
      - configMap:
          name: grafana-dashboard-definition-kubernetes-nodes
        name: grafana-dashboard-definition-kubernetes-nodes
      - configMap:
          name: grafana-dashboard-definition-kubernetes-pods
        name: grafana-dashboard-definition-kubernetes-pods
      - configMap:
          name: grafana-dashboard-definition-kubernetes-statefulset
        name: grafana-dashboard-definition-kubernetes-statefulset
      - configMap:
          name: scaling-config
        name: scaling-config
---

apiVersion: v1
data:
  prometheus.yml: |-
    global:
      scrape_interval: 30s
      scrape_timeout: 10s
      evaluation_interval: 30s
    scrape_configs:
    # Controller endpoint
    - job_name: controller
      scrape_interval: 3s
      scrape_timeout: 3s
      kubernetes_sd_configs:
      - role: pod
      relabel_configs:
      # Scrape only the the targets matching the following metadata
      - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_pod_label_app, __meta_kubernetes_pod_container_port_name]
        action: keep
        regex: knative-serving;controller;metrics
      # Rename metadata labels to be reader friendly
      - source_labels: [__meta_kubernetes_namespace]
        target_label: namespace
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod
      - source_labels: [__meta_kubernetes_service_name]
        target_label: service
    # Autoscaler endpoint
    - job_name: autoscaler
      scrape_interval: 3s
      scrape_timeout: 3s
      kubernetes_sd_configs:
      - role: pod
      relabel_configs:
      # Scrape only the the targets matching the following metadata
      - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_pod_label_app, __meta_kubernetes_pod_container_port_name]
        action: keep
        regex: knative-serving;autoscaler;metrics
      # Rename metadata labels to be reader friendly
      - source_labels: [__meta_kubernetes_namespace]
        target_label: namespace
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod
      - source_labels: [__meta_kubernetes_service_name]
        target_label: service
    # Activator pods
    - job_name: activator
      scrape_interval: 3s
      scrape_timeout: 3s
      kubernetes_sd_configs:
      - role: pod
      relabel_configs:
      # Scrape only the the targets matching the following metadata
      - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_pod_label_app, __meta_kubernetes_pod_container_port_name]
        action: keep
        regex: knative-serving;activator;metrics-port
      # Rename metadata labels to be reader friendly
      - source_labels: [__meta_kubernetes_namespace]
        target_label: namespace
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod
      - source_labels: [__meta_kubernetes_service_name]
        target_label: service
    # Queue proxy metrics
    - job_name: queue-proxy
      scrape_interval: 3s
      scrape_timeout: 3s
      kubernetes_sd_configs:
      - role: pod
      relabel_configs:
      # Scrape only the the targets matching the following metadata
      - source_labels: [__meta_kubernetes_pod_label_serving_knative_dev_revision, __meta_kubernetes_pod_container_port_name]
        action: keep
        regex: .+;user-metrics
      # Rename metadata labels to be reader friendly
      - source_labels: [__meta_kubernetes_namespace]
        target_label: namespace
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod
      - source_labels: [__meta_kubernetes_service_name]
        target_label: service
    # Fluentd daemonset
    - job_name: fluentd-ds
      kubernetes_sd_configs:
      - role: endpoints
      relabel_configs:
      # Scrape only the the targets matching the following metadata
      - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_service_label_app, __meta_kubernetes_endpoint_port_name]
        action: keep
        regex: knative-monitoring;fluentd-ds;prometheus-metrics
      # Rename metadata labels to be reader friendly
      - source_labels: [__meta_kubernetes_namespace]
        target_label: namespace
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod
      - source_labels: [__meta_kubernetes_service_name]
        target_label: service
    # Istio mesh
    - job_name: istio-mesh
      scrape_interval: 5s
      kubernetes_sd_configs:
      - role: endpoints
      relabel_configs:
      # Scrape only the the targets matching the following metadata
      - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_service_name, __meta_kubernetes_endpoint_port_name]
        action: keep
        regex: istio-system;istio-telemetry;prometheus
      # Rename metadata labels to be reader friendly
      - source_labels: [__meta_kubernetes_namespace]
        target_label: namespace
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod
      - source_labels: [__meta_kubernetes_service_name]
        target_label: service
    # Istio Envoy
    # These are very noisy and not enabled by default.
    # - job_name: istio-envoy
    #   scrape_interval: 5s
    #   kubernetes_sd_configs:
    #   - role: endpoints
    #   relabel_configs:
    #   # Scrape only the the targets matching the following metadata
    #   - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_service_name, __meta_kubernetes_endpoint_port_name]
    #     action: keep
    #     regex: istio-system;istio-statsd-prom-bridge;statsd-prom
    #   # Rename metadata labels to be reader friendly
    #   - source_labels: [__meta_kubernetes_namespace]
    #     target_label: namespace
    #   - source_labels: [__meta_kubernetes_pod_name]
    #     target_label: pod
    #   - source_labels: [__meta_kubernetes_service_name]
    #     target_label: service
    # Istio policy
    - job_name: istio-policy
      scrape_interval: 5s
      kubernetes_sd_configs:
      - role: endpoints
      relabel_configs:
      # Scrape only the the targets matching the following metadata
      - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_service_name, __meta_kubernetes_endpoint_port_name]
        action: keep
        regex: istio-system;istio-policy;http-monitoring
      # Rename metadata labels to be reader friendly
      - source_labels: [__meta_kubernetes_namespace]
        target_label: namespace
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod
      - source_labels: [__meta_kubernetes_service_name]
        target_label: service
    # Istio telemetry
    - job_name: istio-telemetry
      scrape_interval: 5s
      kubernetes_sd_configs:
      - role: endpoints
      relabel_configs:
      # Scrape only the the targets matching the following metadata
      - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_service_name, __meta_kubernetes_endpoint_port_name]
        action: keep
        regex: istio-system;istio-telemetry;http-monitoring
      # Rename metadata labels to be reader friendly
      - source_labels: [__meta_kubernetes_namespace]
        target_label: namespace
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod
      - source_labels: [__meta_kubernetes_service_name]
        target_label: service
    # Istio pilot
    - job_name: istio-pilot
      scrape_interval: 5s
      kubernetes_sd_configs:
      - role: endpoints
      relabel_configs:
      # Scrape only the the targets matching the following metadata
      - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_service_name, __meta_kubernetes_endpoint_port_name]
        action: keep
        regex: istio-system;istio-pilot;http-monitoring
      # Rename metadata labels to be reader friendly
      - source_labels: [__meta_kubernetes_namespace]
        target_label: namespace
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod
      - source_labels: [__meta_kubernetes_service_name]
        target_label: service
    # Kube API server
    - job_name: kube-apiserver
      scheme: https
      kubernetes_sd_configs:
      - role: endpoints
      bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
      tls_config:
        ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
        server_name: kubernetes
        insecure_skip_verify: false
      relabel_configs:
      # Scrape only the the targets matching the following metadata
      - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_service_label_component, __meta_kubernetes_service_label_provider, __meta_kubernetes_endpoint_port_name]
        action: keep
        regex: default;apiserver;kubernetes;https
      # Rename metadata labels to be reader friendly
      - source_labels: [__meta_kubernetes_namespace]
        target_label: namespace
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod
      - source_labels: [__meta_kubernetes_service_name]
        target_label: service
    # Kube controller manager
    - job_name: kube-controller-manager
      kubernetes_sd_configs:
      - role: endpoints
      relabel_configs:
      # Scrape only the the targets matching the following metadata
      - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_service_label_app, __meta_kubernetes_endpoint_port_name]
        action: keep
        regex: knative-monitoring;kube-controller-manager;http-metrics
      # Rename metadata labels to be reader friendly
      - source_labels: [__meta_kubernetes_namespace]
        target_label: namespace
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod
      - source_labels: [__meta_kubernetes_service_name]
        target_label: service
    # Kube scheduler
    - job_name: kube-scheduler
      kubernetes_sd_configs:
      - role: endpoints
      relabel_configs:
      # Scrape only the the targets matching the following metadata
      - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_service_label_k8s_app, __meta_kubernetes_endpoint_port_name]
        action: keep
        regex: kube-system;kube-scheduler;http-metrics
      # Rename metadata labels to be reader friendly
      - source_labels: [__meta_kubernetes_namespace]
        target_label: namespace
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod
      - source_labels: [__meta_kubernetes_service_name]
        target_label: service
    # Kube state metrics on https-main port
    - job_name: kube-state-metrics-https-main
      honor_labels: true
      scheme: https
      kubernetes_sd_configs:
      - role: endpoints
      bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
      tls_config:
        insecure_skip_verify: true
      relabel_configs:
      # Scrape only the the targets matching the following metadata
      - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_service_label_app, __meta_kubernetes_endpoint_port_name]
        action: keep
        regex: knative-monitoring;kube-state-metrics;https-main
      # Rename metadata labels to be reader friendly
      - source_labels: [__meta_kubernetes_namespace]
        target_label: namespace
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod
      - source_labels: [__meta_kubernetes_service_name]
        target_label: service
    # Kube state metrics on https-self port
    - job_name: kube-state-metrics-https-self
      scheme: https
      kubernetes_sd_configs:
      - role: endpoints
      bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
      tls_config:
        insecure_skip_verify: true
      relabel_configs:
      # Scrape only the the targets matching the following metadata
      - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_service_label_app, __meta_kubernetes_endpoint_port_name]
        action: keep
        regex: knative-monitoring;kube-state-metrics;https-self
      # Rename metadata labels to be reader friendly
      - source_labels: [__meta_kubernetes_namespace]
        target_label: namespace
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod
      - source_labels: [__meta_kubernetes_service_name]
        target_label: service
    # Kubelet - nodes
    # Rather than connecting directly to the node, the scrape is proxied though the
    # Kubernetes apiserver.  This means it will work if Prometheus is running out of
    # cluster, or can't connect to nodes for some other reason (e.g. because of
    # firewalling).
    - job_name: kubernetes-nodes
      scheme: https
      tls_config:
        ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
      bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
      kubernetes_sd_configs:
      - role: node
      relabel_configs:
      - action: labelmap
        regex: __meta_kubernetes_node_label_(.+)
      - target_label: __address__
        replacement: kubernetes.default.svc:443
      - source_labels: [__meta_kubernetes_node_name]
        target_label: __metrics_path__
        replacement: /api/v1/nodes/${1}/proxy/metrics
    # Kubelet - cAdvisor
    #
    # This is required for Kubernetes 1.7.3 and later, where cAdvisor metrics
    # (those whose names begin with 'container_') have been removed from the
    # Kubelet metrics endpoint.  This job scrapes the cAdvisor endpoint to
    # retrieve those metrics.
    #
    # In Kubernetes 1.7.0-1.7.2, these metrics are only exposed on the cAdvisor
    # HTTP endpoint; use "replacement: /api/v1/nodes/${1}:4194/proxy/metrics"
    # in that case (and ensure cAdvisor's HTTP server hasn't been disabled with
    # the --cadvisor-port=0 Kubelet flag).
    - job_name: kubernetes-cadvisor
      scrape_interval: 15s
      scheme: https
      tls_config:
        ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
      bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
      kubernetes_sd_configs:
      - role: node
      relabel_configs:
      - action: labelmap
        regex: __meta_kubernetes_node_label_(.+)
      - target_label: __address__
        replacement: kubernetes.default.svc:443
      - source_labels: [__meta_kubernetes_node_name]
        target_label: __metrics_path__
        replacement: /api/v1/nodes/${1}/proxy/metrics/cadvisor
    # Node exporter
    - job_name: node-exporter
      scheme: https
      kubernetes_sd_configs:
      - role: endpoints
      bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
      tls_config:
        insecure_skip_verify: true
      relabel_configs:
      # Scrape only the the targets matching the following metadata
      - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_service_label_app, __meta_kubernetes_endpoint_port_name]
        action: keep
        regex: knative-monitoring;node-exporter;https
      # Rename metadata labels to be reader friendly
      - source_labels: [__meta_kubernetes_namespace]
        target_label: namespace
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod
      - source_labels: [__meta_kubernetes_service_name]
        target_label: service
    # Prometheus
    - job_name: prometheus
      kubernetes_sd_configs:
      - role: endpoints
      relabel_configs:
      # Scrape only the the targets matching the following metadata
      - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_service_label_app, __meta_kubernetes_endpoint_port_name]
        action: keep
        regex: knative-monitoring;prometheus;web
      # Rename metadata labels to be reader friendly
      - source_labels: [__meta_kubernetes_namespace]
        target_label: namespace
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod
      - source_labels: [__meta_kubernetes_service_name]
        target_label: service
kind: ConfigMap
metadata:
  labels:
    name: prometheus-scrape-config
    serving.knative.dev/release: "v0.7.0"
  name: prometheus-scrape-config
  namespace: knative-monitoring
---

apiVersion: v1
kind: Service
metadata:
  labels:
    app: kube-controller-manager
    serving.knative.dev/release: "v0.7.0"
  name: kube-controller-manager
  namespace: knative-monitoring
spec:
  clusterIP: None
  ports:
  - name: http-metrics
    port: 10252
    protocol: TCP
    targetPort: 10252
  selector:
    k8s-app: kube-controller-manager
  type: ClusterIP
---

apiVersion: v1
kind: Service
metadata:
  labels:
    app: prometheus
    serving.knative.dev/release: "v0.7.0"
  name: prometheus-system-discovery
  namespace: knative-monitoring
spec:
  clusterIP: None
  ports:
  - name: web
    port: 9090
    protocol: TCP
    targetPort: web
  selector:
    app: prometheus
  sessionAffinity: None
  type: ClusterIP---

apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: prometheus-system
  namespace: knative-monitoring---

apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: prometheus-system
  namespace: default
rules:
- apiGroups:
  - ""
  resources:
  - nodes
  - services
  - endpoints
  - pods
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get---

apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: prometheus-system
  namespace: knative-monitoring
rules:
- apiGroups:
  - ""
  resources:
  - nodes
  - services
  - endpoints
  - pods
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get---

apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: prometheus-system
  namespace: kube-system
rules:
- apiGroups:
  - ""
  resources:
  - services
  - endpoints
  - pods
  verbs:
  - get
  - list
  - watch---

apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: prometheus-system
  namespace: istio-system
rules:
- apiGroups:
  - ""
  resources:
  - nodes/metrics
  - nodes
  - services
  - endpoints
  - pods
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: prometheus-system
  namespace: knative-monitoring
rules:
- apiGroups:
  - ""
  resources:
  - nodes/metrics
  - nodes
  - services
  - endpoints
  - pods
  - nodes/proxy
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
- nonResourceURLs:
  - /metrics
  verbs:
  - get---

apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: prometheus-system
  namespace: default
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: prometheus-system
subjects:
- kind: ServiceAccount
  name: prometheus-system
  namespace: knative-monitoring---

apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: prometheus-system
  namespace: knative-monitoring
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: prometheus-system
subjects:
- kind: ServiceAccount
  name: prometheus-system
  namespace: knative-monitoring---

apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: prometheus-system
  namespace: kube-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: prometheus-system
subjects:
- kind: ServiceAccount
  name: prometheus-system
  namespace: knative-monitoring---

apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: prometheus-system
  namespace: istio-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: prometheus-system
subjects:
- kind: ServiceAccount
  name: prometheus-system
  namespace: knative-monitoring---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: prometheus-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: prometheus-system
subjects:
- kind: ServiceAccount
  name: prometheus-system
  namespace: knative-monitoring---

apiVersion: v1
kind: Service
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: prometheus-system-np
  namespace: knative-monitoring
spec:
  ports:
  - port: 8080
    targetPort: 9090
  selector:
    app: prometheus
  type: NodePort---

apiVersion: apps/v1
kind: StatefulSet
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: prometheus-system
  namespace: knative-monitoring
spec:
  podManagementPolicy: Parallel
  replicas: 2
  selector:
    matchLabels:
      app: prometheus
  serviceName: prometheus-system
  template:
    metadata:
      labels:
        app: prometheus
        serving.knative.dev/release: "v0.7.0"
    spec:
      containers:
      - args:
        - --config.file=/etc/prometheus/prometheus.yml
        - --storage.tsdb.path=/prometheus
        - --storage.tsdb.retention=2d
        - --storage.tsdb.no-lockfile
        - --web.enable-lifecycle
        - --web.route-prefix=/
        image: prom/prometheus:v2.2.1
        imagePullPolicy: IfNotPresent
        livenessProbe:
          failureThreshold: 10
          httpGet:
            path: /-/healthy
            port: web
            scheme: HTTP
          initialDelaySeconds: 30
          timeoutSeconds: 3
        name: prometheus
        ports:
        - containerPort: 9090
          name: web
          protocol: TCP
        readinessProbe:
          failureThreshold: 10
          httpGet:
            path: /-/ready
            port: web
            scheme: HTTP
          timeoutSeconds: 3
        resources:
          limits:
            memory: 1000Mi
          requests:
            memory: 400Mi
        terminationMessagePath: /dev/termination-log
        volumeMounts:
        - mountPath: /etc/prometheus
          name: prometheus-config-volume
        - mountPath: /prometheus
          name: prometheus-storage-volume
      serviceAccountName: prometheus-system
      terminationGracePeriodSeconds: 600
      volumes:
      - configMap:
          defaultMode: 420
          name: prometheus-scrape-config
        name: prometheus-config-volume
      - emptyDir: {}
        name: prometheus-storage-volume
  updateStrategy:
    type: RollingUpdate
---

apiVersion: v1
kind: Service
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: zipkin
  namespace: istio-system
spec:
  ports:
  - name: http
    port: 9411
  selector:
    app: zipkin---

apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    serving.knative.dev/release: "v0.7.0"
  name: zipkin
  namespace: istio-system
spec:
  replicas: 1
  selector:
    matchLabels:
      app: zipkin
  template:
    metadata:
      annotations:
        sidecar.istio.io/inject: "false"
      labels:
        app: zipkin
        serving.knative.dev/release: "v0.7.0"
    spec:
      containers:
      - env:
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: STORAGE_TYPE
          value: elasticsearch
        - name: ES_HOSTS
          value: elasticsearch-logging.knative-monitoring.svc.cluster.local:9200
        - name: ES_INDEX
          value: zipkin
        - name: ZIPKIN_UI_LOGS_URL
          value: http://localhost:8001/api/v1/namespaces/knative-monitoring/services/kibana-logging/proxy/app/kibana#/
        image: docker.io/openzipkin/zipkin:2.13.0
        imagePullPolicy: IfNotPresent
        name: zipkin
        ports:
        - containerPort: 9411
        resources:
          limits:
            memory: 1000Mi
          requests:
            memory: 256Mi
`
