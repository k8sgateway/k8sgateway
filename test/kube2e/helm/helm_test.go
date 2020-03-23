package helm_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/version"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/gloo/test/kube2e"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var _ = Describe("Kube2e: helm", func() {

	It("uses helm to upgrade to a higher 1.3.x version without errors", func() {

		// check that the version is 1.3.0
		Expect(GetGlooServerVersion(testHelper.InstallNamespace)).To(Equal("1.3.0"))

		// upgrade to most recent gloo version
		runAndCleanCommand("helm", "upgrade", "gloo", "gloo/gloo",
			"-n", testHelper.InstallNamespace)

		// check that the version is the most recent one by comparing with the client version
		glooVersion, err := version.GetClientServerVersions(version.NewKube(testHelper.InstallNamespace))
		Expect(err).To(BeNil())
		Expect(GetGlooServerVersion(testHelper.InstallNamespace)).To(Equal(glooVersion.GetClient().Version))

		kube2e.GlooctlCheckEventuallyHealthy(testHelper)
	})

	It("uses helm to update the settings without errors", func() {

		// check that the setting is the default to start
		client := helpers.MustSettingsClient()
		settings, err := client.Read(testHelper.InstallNamespace, defaults.SettingsName, clients.ReadOpts{})
		Expect(err).To(BeNil())
		Expect(settings.GetGloo().GetInvalidConfigPolicy().GetInvalidRouteResponseCode()).To(Equal(uint32(404)))

		// update the settings with `helm upgrade` (without updating the gloo version)
		runAndCleanCommand("helm", "upgrade", "gloo", "gloo/gloo",
			"-n", testHelper.InstallNamespace,
			"--set", "settings.invalidConfigPolicy.invalidRouteResponseCode=400",
			"--version", GetGlooServerVersion(testHelper.InstallNamespace))

		// check that the setting updated
		settings, err = client.Read(testHelper.InstallNamespace, defaults.SettingsName, clients.ReadOpts{})
		Expect(err).To(BeNil())
		Expect(settings.GetGloo().GetInvalidConfigPolicy().GetInvalidRouteResponseCode()).To(Equal(uint32(400)))

		kube2e.GlooctlCheckEventuallyHealthy(testHelper)
	})

})

func GetGlooServerVersion(namespace string) (v string) {
	glooVersion, err := version.GetClientServerVersions(version.NewKube(namespace))
	Expect(err).To(BeNil())
	Expect(len(glooVersion.GetServer())).To(Equal(1))
	for _, container := range glooVersion.GetServer()[0].GetKubernetes().GetContainers() {
		if v == "" {
			v = container.Tag
		} else {
			Expect(container.Tag).To(Equal(v))
		}
	}
	return v
}
