package kube2e

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"

	"go.uber.org/zap/zapcore"

	"github.com/golang/protobuf/proto"
	"github.com/solo-io/go-utils/testutils/goimpl"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/check"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	clienthelpers "github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/k8s-utils/kubeutils"
	"github.com/solo-io/k8s-utils/testutils/helper"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	. "github.com/onsi/gomega"
	errors "github.com/rotisserie/eris"
	"k8s.io/client-go/kubernetes"
)

func MustKubeClient() kubernetes.Interface {
	restConfig, err := kubeutils.GetConfig("", "")
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	kubeClient, err := kubernetes.NewForConfig(restConfig)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	return kubeClient
}

// Check that everything is OK by running `glooctl check`
func GlooctlCheckEventuallyHealthy(offset int, testHelper *helper.SoloTestHelper, timeoutInterval string) {
	EventuallyWithOffset(offset, func() error {
		opts := &options.Options{
			Metadata: core.Metadata{
				Namespace: testHelper.InstallNamespace,
			},
			Top: options.Top{
				Ctx: context.Background(),
			},
		}
		err := check.CheckResources(opts)
		if err != nil {
			return errors.New("glooctl check detected a problem with the installation")
		}
		return nil
	}, timeoutInterval, "5s").Should(BeNil())
}

func GetHelmValuesOverrideFile() (filename string, cleanup func()) {
	values, err := ioutil.TempFile("", "values-*.yaml")
	Expect(err).NotTo(HaveOccurred())

	// disabling usage statistics is not important to the functionality of the tests,
	// but we don't want to report usage in CI since we only care about how our users are actually using Gloo.
	// install to a single namespace so we can run multiple invocations of the regression tests against the
	// same cluster in CI.
	_, err = values.Write([]byte(`
global:
  image:
    pullPolicy: IfNotPresent
  glooRbac:
    namespaced: true
    nameSuffix: e2e-test-rbac-suffix
settings:
  singleNamespace: true
  create: true
  replaceInvalidRoutes: true
gatewayProxies:
  gatewayProxy:
    healthyPanicThreshold: 0
`))
	Expect(err).NotTo(HaveOccurred())

	err = values.Close()
	Expect(err).NotTo(HaveOccurred())

	return values.Name(), func() { _ = os.Remove(values.Name()) }
}

func EventuallyReachesConsistentState(installNamespace string) {
	metricsPort := 9091
	metricsPortString := strconv.Itoa(metricsPort)
	portFwd := exec.Command("kubectl", "port-forward", "-n", installNamespace,
		"deployment/gloo", metricsPortString)
	portFwd.Stdout = os.Stderr
	portFwd.Stderr = os.Stderr
	err := portFwd.Start()
	Expect(err).ToNot(HaveOccurred())

	defer func() {
		if portFwd.Process != nil {
			portFwd.Process.Kill()
		}
	}()

	// make sure we eventually reach an eventually consistent state
	lastSnapOut := getSnapOut(metricsPortString)

	eventuallyConsistentPollingInterval := 7 * time.Second // >= 5s for metrics reporting, which happens every 5s
	time.Sleep(eventuallyConsistentPollingInterval)

	Eventually(func() bool {
		currentSnapOut := getSnapOut(metricsPortString)
		consistent := lastSnapOut == currentSnapOut
		lastSnapOut = currentSnapOut
		return consistent
	}, "30s", eventuallyConsistentPollingInterval).Should(Equal(true))

	Consistently(func() string {
		currentSnapOut := getSnapOut(metricsPortString)
		return currentSnapOut
	}, "30s", eventuallyConsistentPollingInterval).Should(Equal(lastSnapOut))

	// Gloo components are configured to log to the Info level by default
	EventuallyLogLevel(metricsPort, zapcore.InfoLevel)
}

// Copied from: https://github.com/solo-io/go-utils/blob/176c4c008b4d7cde836269c7a817f657b6981236/testutils/assertions.go#L20
func ExpectEqualProtoMessages(g Gomega, a, b proto.Message, optionalDescription ...interface{}) {
	if proto.Equal(a, b) {
		return
	}

	g.Expect(a.String()).To(Equal(b.String()), optionalDescription...)
}

// needs a port-forward of the metrics port before a call to this will work
func getSnapOut(metricsPort string) string {
	var bodyResp string
	Eventually(func() string {
		res, err := http.Post("http://localhost:"+metricsPort+"/metrics", "", nil)
		if err != nil || res.StatusCode != 200 {
			return ""
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		Expect(err).ToNot(HaveOccurred())
		bodyResp = string(body)
		return bodyResp
	}, "5s", "1s").ShouldNot(BeEmpty())

	Expect(bodyResp).To(ContainSubstring("api_gloo_solo_io_emitter_snap_out"))
	findSnapOut := regexp.MustCompile("api_gloo_solo_io_emitter_snap_out ([\\d]+)")
	matches := findSnapOut.FindAllStringSubmatch(bodyResp, -1)
	Expect(matches).To(HaveLen(1))
	snapOut := matches[0][1]
	return snapOut
}

// EventuallyLogLevel ensures that we can query the endpoint responsible for getting the current
// log level of a gloo component, and updating the log level dynamically
func EventuallyLogLevel(port int, logLevel zapcore.Level) {
	url := fmt.Sprintf("http://localhost:%d/logging", port)
	body := bytes.NewReader([]byte(url))

	request, err := http.NewRequest(http.MethodGet, url, body)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	expectedResponse := fmt.Sprintf("{\"level\":\"%s\"}\n", logLevel.String())
	EventuallyWithOffset(1, func() (string, error) {
		return goimpl.ExecuteRequest(request)
	}, time.Second*5, time.Millisecond*100).Should(Equal(expectedResponse))
}

func UpdateDisableTransformationValidationSetting(ctx context.Context, shouldDisable bool, installNamespace string) {
	UpdateSettings(func(settings *v1.Settings) {
		Expect(settings.GetGateway().GetValidation()).NotTo(BeNil())
		settings.GetGateway().GetValidation().DisableTransformationValidation = &wrappers.BoolValue{Value: shouldDisable}
	}, ctx, installNamespace)
}

// enable/disable strict validation
func UpdateAlwaysAcceptSetting(ctx context.Context, alwaysAccept bool, installNamespace string) {
	UpdateSettings(func(settings *v1.Settings) {
		Expect(settings.GetGateway().GetValidation()).NotTo(BeNil())
		settings.GetGateway().GetValidation().AlwaysAccept = &wrappers.BoolValue{Value: alwaysAccept}
	}, ctx, installNamespace)
}

func UpdateRestEdsSetting(ctx context.Context, enableRestEds bool, installNamespace string) {
	UpdateSettings(func(settings *v1.Settings) {
		Expect(settings.GetGloo()).NotTo(BeNil())
		settings.GetGloo().EnableRestEds = &wrappers.BoolValue{Value: enableRestEds}
	}, ctx, installNamespace)
}

func UpdateReplaceInvalidRoutes(ctx context.Context, replaceInvalidRoutes bool, installNamespace string) {
	UpdateSettings(func(settings *v1.Settings) {
		Expect(settings.GetGloo().GetInvalidConfigPolicy()).NotTo(BeNil())
		settings.GetGloo().GetInvalidConfigPolicy().ReplaceInvalidRoutes = replaceInvalidRoutes
	}, ctx, installNamespace)
}

func UpdateSettings(updateSettings func(settings *v1.Settings), ctx context.Context, installNamespace string) {
	// when validation config changes, the validation server restarts -- give time for it to come up again.
	// without the wait, the validation webhook may temporarily fallback to it's failurePolicy, which is not
	// what we want to test.
	// TODO (samheilbron) We should avoid relying on time.Sleep in our tests as these tend to cause flakes
	waitForSettingsToPropagate := func() {
		time.Sleep(3 * time.Second)
	}
	UpdateSettingsWithPropagationDelay(updateSettings, waitForSettingsToPropagate, ctx, installNamespace)
}

func UpdateSettingsWithPropagationDelay(updateSettings func(settings *v1.Settings), waitForSettingsToPropagate func(), ctx context.Context, installNamespace string) {
	settingsClient := clienthelpers.MustSettingsClient(ctx)
	settings, err := settingsClient.Read(installNamespace, "default", clients.ReadOpts{})
	Expect(err).NotTo(HaveOccurred())

	updateSettings(settings)

	_, err = settingsClient.Write(settings, clients.WriteOpts{OverwriteExisting: true})
	Expect(err).NotTo(HaveOccurred())

	waitForSettingsToPropagate()
}
