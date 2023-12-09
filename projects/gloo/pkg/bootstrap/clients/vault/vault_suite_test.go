package vault

import (
    "github.com/solo-io/gloo/test/gomega/assertions"
    "go.opencensus.io/stats/view"
    "testing"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

func TestVaultClient(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Vault Client Suite")
}

var _ = BeforeSuite(func() {
    resetViews()
})

var _ = AfterSuite(func() {
    resetViews()
})

func resetViews() {
    views := []*view.View{
        mLastLoginSuccessView,
        mLoginFailuresView,
        mLoginSuccessesView,
        mLastLoginFailureView,
    }
    view.Unregister(views...)
    _ = view.Register(metricViews...)
    assertions.ExpectStatLastValueMatches(mLastLoginSuccess, BeZero())
    assertions.ExpectStatLastValueMatches(mLastLoginFailure, BeZero())
    assertions.ExpectStatSumMatches(mLoginSuccesses, BeZero())
    assertions.ExpectStatSumMatches(mLoginFailures, BeZero())
}
