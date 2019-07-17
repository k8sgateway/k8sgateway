package consul

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Conversions", func() {

	It("correctly generates the name for the fake upstream", func() {
		Expect(fakeUpstreamName("my-consul-service")).To(Equal(upstreamNamePrefix + "my-consul-service"))
	})

	It("correctly detects upstreams derived from Kubernetes services", func() {
		Expect(IsConsulUpstream(upstreamNamePrefix + "my-service")).To(BeTrue())
		Expect(IsConsulUpstream("my-" + upstreamNamePrefix + "service")).To(BeFalse())
		Expect(IsConsulUpstream("consul:my-service-8080")).To(BeFalse())
	})

	It("correctly converts a list of services to upstreams", func() {
		servicesWithDataCenters := []*ServiceMeta{
			{Name: "svc-1", DataCenters: []string{"dc1", "dc2"}},
			{Name: "svc-2", DataCenters: []string{"dc1", "dc3", "dc4"}},
		}

		usList := toUpstreamList(servicesWithDataCenters)
		usList.Sort()

		Expect(usList).To(HaveLen(2))

		Expect(usList[0].Metadata.Name).To(Equal(upstreamNamePrefix + "svc-1"))
		Expect(usList[0].Metadata.Namespace).To(BeEmpty())
		Expect(usList[0].UpstreamSpec.GetConsul()).NotTo(BeNil())
		Expect(usList[0].UpstreamSpec.GetConsul().ServiceName).To(Equal("svc-1"))
		Expect(usList[0].UpstreamSpec.GetConsul().DataCenters).To(ConsistOf("dc1", "dc2"))

		Expect(usList[1].Metadata.Name).To(Equal(upstreamNamePrefix + "svc-2"))
		Expect(usList[1].Metadata.Namespace).To(BeEmpty())
		Expect(usList[1].UpstreamSpec.GetConsul()).NotTo(BeNil())
		Expect(usList[1].UpstreamSpec.GetConsul().ServiceName).To(Equal("svc-2"))
		Expect(usList[1].UpstreamSpec.GetConsul().DataCenters).To(ConsistOf("dc1", "dc3", "dc4"))
	})

	It("correctly consolidates service information from different data centers", func() {
		input := []*dataCenterServicesTuple{
			{
				dataCenter: "dc-1",
				services: map[string][]string{
					"svc-1": {"tag-1", "tag-2"},
					"svc-2": {"tag-2"},
					"svc-3": {},
				},
			},
			{
				dataCenter: "dc-2",
				services: map[string][]string{
					"svc-1": {"tag-3"},
					"svc-2": {},
					"svc-4": nil,
				},
			},
		}

		result := toServiceMetaSlice(input)

		Expect(result).To(ConsistOf(
			[]*ServiceMeta{
				{
					Name:        "svc-1",
					DataCenters: []string{"dc-1", "dc-2"},
					Tags:        []string{"tag-1", "tag-2", "tag-3"},
				},
				{
					Name:        "svc-2",
					DataCenters: []string{"dc-1", "dc-2"},
					Tags:        []string{"tag-2"},
				},
				{
					Name:        "svc-3",
					DataCenters: []string{"dc-1"},
					Tags:        nil,
				},
				{
					Name:        "svc-4",
					DataCenters: []string{"dc-2"},
					Tags:        nil,
				},
			},
		))

	})
})
