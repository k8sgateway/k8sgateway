package utils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/solo-io/gloo/projects/gloo/pkg/plugins/utils"

	"github.com/gogo/protobuf/types"
)

var _ = Describe("Any", func() {

	It("should deserialized a proto message from map", func() {

		duration := types.Duration{
			Nanos:   1,
			Seconds: 2,
		}
		anyduration, err := types.MarshalAny(&duration)
		Expect(err).NotTo(HaveOccurred())

		protos := map[string]*types.Any{
			"duration": anyduration,
		}

		var outm types.Duration
		err = UnmarshalAnyFromMap(protos, "duration", &outm)
		Expect(err).NotTo(HaveOccurred())

		Expect(outm).To(Equal(duration))
	})

	It("should error if no name found with expected error", func() {

		protos := map[string]*types.Any{}
		var outm types.Duration
		err := UnmarshalAnyFromMap(protos, "duration", &outm)
		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal(NotFoundError))
	})

	It("should error if proto is bad with other error", func() {

		anyduration, err := types.MarshalAny(&types.Duration{})
		Expect(err).NotTo(HaveOccurred())
		anyduration.Value = []byte("bad proto")
		protos := map[string]*types.Any{
			"duration": anyduration,
		}
		var outm types.Duration
		err = UnmarshalAnyFromMap(protos, "duration", &outm)
		Expect(err).To(HaveOccurred())
		Expect(err).NotTo(Equal(NotFoundError))
	})

	Describe("Any from plugins", func() {

		It("should return not found for nil plugins", func() {
			var outm types.Duration
			err := UnmarshalAnyPlugins(nil, "duration", &outm)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(NotFoundError))
		})

		It("should return not found for typed nil plugins", func() {
			var p *plugins
			var outm types.Duration
			err := UnmarshalAnyPlugins(p, "duration", &outm)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(NotFoundError))
		})

		It("should return not found for nil plugin map", func() {
			var p plugins
			var outm types.Duration
			err := UnmarshalAnyPlugins(&p, "duration", &outm)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(NotFoundError))
		})

	})

})

type plugins struct {
	plugins map[string]*types.Any
}

func (p *plugins) GetPlugins() map[string]*types.Any { return p.plugins }
