package tests

import (
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	"github.com/solo-io/gloo/test/kubernetes/e2e/features/helm"
	"github.com/solo-io/gloo/test/kubernetes/e2e/features/helm_settings"
)

func HelmSuiteRunner() e2e.SuiteRunner {
	helmSuiteRunner := e2e.NewSuiteRunner(false)
	helmSuiteRunner.Register("Helm", helm.NewTestingSuite)
	helmSuiteRunner.Register("HelmSettings", helm_settings.NewTestingSuite)
	return helmSuiteRunner
}
