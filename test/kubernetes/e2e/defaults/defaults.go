package defaults

import (
	"context"
	"path/filepath"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/solo-io/gloo/pkg/utils/kubeutils/kubectl"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	"github.com/solo-io/skv2/codegen/util"
)

type CommonTestSuite interface {
	Ctx() context.Context
	TestInstallation() *e2e.TestInstallation // DO_NOT_SUBMIT: what to do about solo-projects? uses a struct that embeds this struct
	Assert() *assert.Assertions
	Resources() []Resource
}

type CommonTestSuiteImpl struct {
	suite.Suite
	ctx       context.Context
	ti        *e2e.TestInstallation
	resources []Resource
}

func (s *CommonTestSuiteImpl) Ctx() context.Context {
	return s.ctx
}

func (s *CommonTestSuiteImpl) TestInstallation() *e2e.TestInstallation {
	return s.ti
}

func (s *CommonTestSuiteImpl) Resources() []Resource {
	return s.resources
}

func NewCommonTestSuiteImpl(ctx context.Context, testInst *e2e.TestInstallation, resources []Resource) *CommonTestSuiteImpl {
	return &CommonTestSuiteImpl{
		ctx:       ctx,
		ti:        testInst,
		resources: resources,
	}
}

var (
	CurlPodExecOpt = kubectl.PodExecOptions{
		Name:      "curl",
		Namespace: "curl",
		Container: "curl",
	}

	CurlPod = &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "curl",
			Namespace: "curl",
		},
	}

	CurlPodManifest = filepath.Join(util.MustGetThisDir(), "testdata", "curl_pod.yaml")

	HttpEchoPod = &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "http-echo",
			Namespace: "http-echo",
		},
	}

	HttpEchoPodManifest = filepath.Join(util.MustGetThisDir(), "testdata", "http_echo.yaml")

	TcpEchoPod = &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "tcp-echo",
			Namespace: "tcp-echo",
		},
	}

	TcpEchoPodManifest = filepath.Join(util.MustGetThisDir(), "testdata", "tcp_echo.yaml")

	NginxPod = &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx",
			Namespace: "nginx",
		},
	}

	NginxSvc = &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx",
			Namespace: "nginx",
		},
	}

	NginxPodManifest = filepath.Join(util.MustGetThisDir(), "testdata", "nginx_pod.yaml")

	NginxResponse = `<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>`
)

// Resource interface
type Resource interface {
	Install(s CommonTestSuite)
	Delete(s CommonTestSuite)
}

// InstallResources installs multiple resources for the test suite
func InstallResources(s CommonTestSuite) {
	for _, r := range s.Resources() {
		r.Install(s)
	}
}

// DeleteResources deletes multiple resources for the test suite
func DeleteResources(s CommonTestSuite) {
	for _, r := range s.Resources() {
		r.Delete(s)
	}
}

// CurlPodResource
type CurlPodResource struct{}

func (c *CurlPodResource) eventuallyRunning(s CommonTestSuite) {
	s.TestInstallation().Assertions.EventuallyPodsRunning(s.Ctx(), CurlPod.ObjectMeta.GetNamespace(), metav1.ListOptions{
		LabelSelector: "app.kubernetes.io/name=curl",
	})
}

func (c *CurlPodResource) Install(s CommonTestSuite) {
	err := s.TestInstallation().Actions.Kubectl().ApplyFile(s.Ctx(), CurlPodManifest)
	s.Assert().NoError(err)
	c.eventuallyRunning(s)
}

func (c *CurlPodResource) Delete(s CommonTestSuite) {
	output, err := s.TestInstallation().Actions.Kubectl().DeleteFileWithOutput(s.Ctx(), CurlPodManifest)
	s.Assert().NoError(err, "can delete curl pod")
	s.TestInstallation().Assertions.ExpectObjectDeleted(CurlPodManifest, err, output)
}

// func CurlPodEventuallyRunning(s CommonTestSuite) {
// 	// Check that test resources are running
// 	s.TestInstallation().Assertions.EventuallyPodsRunning(s.Ctx(), CurlPod.ObjectMeta.GetNamespace(), metav1.ListOptions{
// 		LabelSelector: "app.kubernetes.io/name=curl",
// 	})
// }

// func InstallCurlPod(s CommonTestSuite) {
// 	err := s.TestInstallation().Actions.Kubectl().ApplyFile(s.Ctx(), CurlPodManifest)
// 	s.Assert().NoError(err)
// 	CurlPodEventuallyRunning(s)
// }

// func DeleteCurlPod(s CommonTestSuite) {
// 	output, err := s.TestInstallation().Actions.Kubectl().DeleteFileWithOutput(s.Ctx(), CurlPodManifest)
// 	s.Assert().NoError(err, "can delete curl pod")
// 	s.TestInstallation().Assertions.ExpectObjectDeleted(CurlPodManifest, err, output)
// }

// // Or like this?
// func CurlPodEventuallyRunning(ctx context.Context, ti *e2e.TestInstallation) {
// 	// Check that test resources are running
// 	ti.Assertions.EventuallyPodsRunning(ctx, CurlPod.ObjectMeta.GetNamespace(), metav1.ListOptions{
// 		LabelSelector: "app.kubernetes.io/name=curl",
// 	})
// }

// func InstallCurlPod(ctx context.Context, ti *e2e.TestInstallation, s *suite.Suite) {
// 	err := ti.Actions.Kubectl().ApplyFile(ctx, CurlPodManifest)
// 	s.Assert().NoError(err)
// 	CurlPodEventuallyRunning(ctx, ti)
// }

// func DeleteCurlPod(ctx context.Context, ti *e2e.TestInstallation, s *suite.Suite) {
// 	output, err := ti.Actions.Kubectl().DeleteFileWithOutput(ctx, CurlPodManifest)
// 	s.Assert().NoError(err, "can delete curl pod")
// 	ti.Assertions.ExpectObjectDeleted(CurlPodManifest, err, output)
// }
