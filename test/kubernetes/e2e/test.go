package e2e

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/solo-io/gloo/test/kubernetes/testutils/actions"
	"github.com/solo-io/gloo/test/kubernetes/testutils/assertions"
	"github.com/solo-io/gloo/test/kubernetes/testutils/cluster"
	"github.com/solo-io/gloo/test/kubernetes/testutils/gloogateway"
	"github.com/solo-io/gloo/test/kubernetes/testutils/helper"
	testruntime "github.com/solo-io/gloo/test/kubernetes/testutils/runtime"
	"github.com/solo-io/gloo/test/testutils"
)

// MustTestHelper returns the SoloTestHelper used for e2e tests
// The SoloTestHelper is a wrapper around `glooctl` and we should eventually phase it out
// in favor of using the exact tool that users rely on
func MustTestHelper(ctx context.Context, installation *TestInstallation) *helper.SoloTestHelper {
	testHelper, err := helper.GetTestHelperForRootDir(ctx, testutils.GitRootDirectory(), installation.Metadata.InstallNamespace)
	if err != nil {
		panic(err)
	}

	testHelper.SetKubeCli(installation.ClusterContext.Cli)

	return testHelper
}

// CreateTestInstallation is the simplest way to construct a TestInstallation in Gloo Gateway OSS
// It is syntactic sugar on top of CreateTestInstallationForCluster
func CreateTestInstallation(
	t *testing.T,
	glooGatewayContext *gloogateway.Context,
) *TestInstallation {
	runtimeContext := testruntime.NewContext()
	clusterContext := cluster.MustKindContext(runtimeContext.ClusterName)

	return CreateTestInstallationForCluster(t, runtimeContext, clusterContext, glooGatewayContext)
}

// CreateTestInstallationForCluster is the standard way to construct a TestInstallation
// It accepts context objects from 3 relevant sources:
//
//	runtime - These are properties that are supplied at runtime and will impact how tests are executed
//	cluster - These are properties that are used to connect to the Kubernetes cluster
//	glooGateway - These are properties that are relevant to how Gloo Gateway will be configured
func CreateTestInstallationForCluster(
	t *testing.T,
	runtimeContext testruntime.Context,
	clusterContext *cluster.Context,
	glooGatewayContext *gloogateway.Context,
) *TestInstallation {
	installation := &TestInstallation{
		// RuntimeContext contains the set of properties that are defined at runtime by whoever is invoking tests
		RuntimeContext: runtimeContext,

		// ClusterContext contains the metadata about the Kubernetes Cluster that is used for this TestCluster
		ClusterContext: clusterContext,

		// Maintain a reference to the Metadata used for this installation
		Metadata: glooGatewayContext,

		// ResourceClients are only available _after_ installing Gloo Gateway
		ResourceClients: nil,

		// Create an operations provider, and point it to the running installation
		Actions: actions.NewActionsProvider().
			WithClusterContext(clusterContext).
			WithGlooGatewayContext(glooGatewayContext),

		// Create an assertions provider, and point it to the running installation
		Assertions: assertions.NewProvider(t).
			WithClusterContext(clusterContext).
			WithGlooGatewayContext(glooGatewayContext),

		// GeneratedFiles contains the unique location where files generated during the execution
		// of tests against this installation will be stored
		// By creating a unique location, per TestInstallation and per Cluster.Name we guarantee isolation
		// between TestInstallation outputs per CI run
		GeneratedFiles: MustGeneratedFiles(glooGatewayContext.InstallNamespace, clusterContext.Name),
	}
	runtime.SetFinalizer(installation, func(i *TestInstallation) { i.finalize() })
	return installation
}

// TestInstallation is the structure around a set of tests that validate behavior for an installation
// of Gloo Gateway.
type TestInstallation struct {
	fmt.Stringer

	// RuntimeContext contains the set of properties that are defined at runtime by whoever is invoking tests
	RuntimeContext testruntime.Context

	// ClusterContext contains the metadata about the Kubernetes Cluster that is used for this TestCluster
	ClusterContext *cluster.Context

	// Metadata contains the properties used to install Gloo Gateway
	Metadata *gloogateway.Context

	// ResourceClients is a set of clients that can manipulate resources owned by Gloo Gateway
	ResourceClients gloogateway.ResourceClients

	// Actions is the entity that creates actions that can be executed by the Operator
	Actions *actions.Provider

	// Assertions is the entity that creates assertions that can be executed by the Operator
	Assertions *assertions.Provider

	// GeneratedFiles is the collection of directories and files that this test installation _may_ create
	GeneratedFiles GeneratedFiles

	// IstioctlBinary is the path to the istioctl binary that can be used to interact with Istio
	IstioctlBinary string
}

func (i *TestInstallation) String() string {
	return i.Metadata.InstallNamespace
}

func (i *TestInstallation) finalize() {
	if err := os.RemoveAll(i.GeneratedFiles.TempDir); err != nil {
		panic(fmt.Sprintf("Failed to remove temporary directory: %s", i.GeneratedFiles.TempDir))
	}
}

func (i *TestInstallation) AddIstioctl(ctx context.Context) error {
	istioctl, err := cluster.GetIstioctl(ctx)
	if err != nil {
		return fmt.Errorf("failed to download istio: %w", err)
	}
	i.IstioctlBinary = istioctl
	return nil
}

func (i *TestInstallation) InstallMinimalIstio(ctx context.Context) error {
	return cluster.InstallMinimalIstio(ctx, i.IstioctlBinary, i.ClusterContext.KubeContext)
}

func (i *TestInstallation) UninstallIstio() error {
	return cluster.UninstallIstio(i.IstioctlBinary, i.ClusterContext.KubeContext)
}

func (i *TestInstallation) CreateIstioBugReport(ctx context.Context) {
	cluster.CreateIstioBugReport(ctx, i.IstioctlBinary, i.ClusterContext.KubeContext, i.GeneratedFiles.FailureDir)
}

func (i *TestInstallation) InstallGlooGateway(ctx context.Context, installFn func(ctx context.Context) error) {
	if !testutils.ShouldSkipInstall() {
		err := installFn(ctx)
		i.Assertions.Require.NoError(err)
		i.Assertions.EventuallyInstallationSucceeded(ctx)
		i.Assertions.EventuallyGlooReachesConsistentState(i.Metadata.InstallNamespace)
	}

	// We can only create the ResourceClients after the CRDs exist in the Cluster
	clients, err := gloogateway.NewResourceClients(ctx, i.ClusterContext)
	i.Assertions.Require.NoError(err)
	i.ResourceClients = clients
}

func (i *TestInstallation) UninstallGlooGateway(ctx context.Context, uninstallFn func(ctx context.Context) error) {
	if testutils.ShouldSkipInstall() {
		return
	}
	err := uninstallFn(ctx)
	i.Assertions.Require.NoError(err)
	i.Assertions.EventuallyUninstallationSucceeded(ctx)
}

func (i *TestInstallation) UpgradeGlooGateway(ctx context.Context, serverVersion string, upgradeFn func(ctx context.Context) error) {
	err := upgradeFn(ctx)
	i.Assertions.Require.NoError(err)
	i.Assertions.EventuallyUpgradeSucceeded(ctx, serverVersion)
	i.Assertions.EventuallyGlooReachesConsistentState(i.Metadata.InstallNamespace)
}

// PreFailHandler is the function that is invoked if a test in the given TestInstallation fails
func (i *TestInstallation) PreFailHandler(ctx context.Context) {
	// This is a work in progress
	// The idea here is we want to accumulate ALL information about this TestInstallation into a single directory
	// That way we can upload it in CI, or inspect it locally

	failureDir := i.GeneratedFiles.FailureDir
	err := os.Mkdir(failureDir, os.ModePerm)
	i.Assertions.Require.NoError(err)

	glooLogFilePath := filepath.Join(failureDir, "gloo.log")
	glooLogFile, err := os.OpenFile(glooLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	i.Assertions.Require.NoError(err)
	defer glooLogFile.Close()

	glooLogsCmd := i.Actions.Kubectl().Command(ctx, "logs", "-n", i.Metadata.InstallNamespace, "deployments/gloo")
	_ = glooLogsCmd.WithStdout(glooLogFile).WithStderr(glooLogFile).Run()

	edgeGatewayLogFilePath := filepath.Join(failureDir, "edge_gateway.log")
	edgeGatewayLogFile, err := os.OpenFile(edgeGatewayLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	i.Assertions.Require.NoError(err)
	defer edgeGatewayLogFile.Close()

	kubeGatewayLogFilePath := filepath.Join(failureDir, "kube_gateway.log")
	kubeGatewayLogFile, err := os.OpenFile(kubeGatewayLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	i.Assertions.Require.NoError(err)
	defer kubeGatewayLogFile.Close()

	namespaces, err := i.Actions.Kubectl().Namespaces(ctx)
	i.Assertions.Require.NoError(err)
	for _, n := range namespaces {
		edgeGatewayLogFile.WriteString(fmt.Sprintf("Logs for edge gateway proxies in namespace %s\n", n))
		edgeGatewayLogsCmd := i.Actions.Kubectl().Command(ctx, "logs", "--all-containers", "--namespace", n, "--prefix", "-l", "gloo=gateway-proxy")
		_ = edgeGatewayLogsCmd.WithStdout(edgeGatewayLogFile).WithStderr(edgeGatewayLogFile).Run()
		edgeGatewayLogFile.WriteString("----------------------------------------------------------------------------------------------------------\n")

		kubeGatewayLogFile.WriteString(fmt.Sprintf("Logs for kube gateway proxies in namespace %s\n", n))
		kubeGatewayLogsCmd := i.Actions.Kubectl().Command(ctx, "logs", "--all-containers", "--namespace", n, "--prefix", "-l", "gloo=kube-gateway")
		_ = kubeGatewayLogsCmd.WithStdout(kubeGatewayLogFile).WithStderr(kubeGatewayLogFile).Run()
		kubeGatewayLogFile.WriteString("----------------------------------------------------------------------------------------------------------\n")
	}

	clusterStateFilePath := filepath.Join(failureDir, "cluster_state.log")
	clusterStateFile, err := os.OpenFile(clusterStateFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	i.Assertions.Require.NoError(err)
	defer clusterStateFile.Close()

	kubectlGetAllCmd := i.Actions.Kubectl().Command(ctx, "get", "all", "-A")
	_ = kubectlGetAllCmd.WithStdout(clusterStateFile).WithStderr(clusterStateFile).Run()
	clusterStateFile.WriteString("\n")

	resourcesToGet := []string{
		"gateways",
		"gatewayclasses",
		"gatewayparameters",
		"routeoptions",
		"virtualhostoptions",
		"upstreams",
		"upstreamgroups",
		"authconfigs",
		"ratelimitconfigs",
		"virtualservices",
		"httproutes",
		"secrets",
	}
	kubectlGetResourcesCmd := i.Actions.Kubectl().Command(ctx, "get", strings.Join(resourcesToGet, ","), "-A")
	_ = kubectlGetResourcesCmd.WithStdout(clusterStateFile).WithStderr(clusterStateFile).Run()
	clusterStateFile.Write([]byte{'\n'})
}

// GeneratedFiles is a collection of files that are generated during the execution of a set of tests
type GeneratedFiles struct {
	// TempDir is the directory where any temporary files should be created
	// Tests may create files for any number of reasons:
	// - A: When a test renders objects in a file, and then uses this file to create and delete values
	// - B: When a test invokes a command that produces a file as a side effect (glooctl, for example)
	// Files in this directory are an implementation detail of the test itself.
	// As a result, it is the callers responsibility to clean up the TempDir when the tests complete
	TempDir string

	// FailureDir is the directory where any assets that are produced on failure will be created
	FailureDir string
}

// MustGeneratedFiles returns GeneratedFiles, or panics if there was an error generating the directories
func MustGeneratedFiles(tmpDirId, clusterId string) GeneratedFiles {
	tmpDir, err := os.MkdirTemp("", tmpDirId)
	if err != nil {
		panic(err)
	}

	// output path is in the format of bug_report/cluster_name/tmp_dir_id
	failureDir := filepath.Join(testruntime.PathToBugReport(), clusterId, tmpDirId)
	err = os.MkdirAll(failureDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	return GeneratedFiles{
		TempDir:    tmpDir,
		FailureDir: failureDir,
	}
}
