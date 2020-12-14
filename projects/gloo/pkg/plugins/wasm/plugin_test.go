package wasm

import (
	"fmt"

	envoy_extensions_filters_http_wasm_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/wasm/v3"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opencontainers/go-digest"
	"github.com/rotisserie/eris"
	configcore "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/core/v3"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/wasm"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	mock_cache "github.com/solo-io/gloo/projects/gloo/pkg/plugins/wasm/mocks"
	"github.com/solo-io/solo-kit/test/matchers"
)

var _ = Describe("wasm plugin", func() {
	var (
		p         *Plugin
		ctrl      *gomock.Controller
		mockCache *mock_cache.MockCache
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockCache = mock_cache.NewMockCache(ctrl)
		imageCache = mockCache
		p = NewPlugin()
	})

	It("should not add filter if wasm config is nil", func() {
		f, err := p.HttpFilters(plugins.Params{}, nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(f).To(BeNil())
	})

	It("will err if plugin cache returns an error", func() {
		image := "hello"
		hl := &v1.HttpListener{
			Options: &v1.HttpListenerOptions{
				Wasm: &wasm.PluginSource{
					Filters: []*wasm.WasmFilter{
						{
							Image: image,
						},
					},
				},
			},
		}

		fakeErr := eris.New("hello")
		mockCache.EXPECT().Add(gomock.Any(), image).Return(digest.Digest(""), fakeErr)
		f, err := p.HttpFilters(plugins.Params{}, hl)
		Expect(err).To(HaveOccurred())
		Expect(f).To(BeNil())
		Expect(err).To(Equal(fakeErr))
	})

	It("will return the proper config", func() {
		sha := "test-sha"
		image := "image"
		wasmFilter := &wasm.WasmFilter{
			Image: image,
			Config: &any.Any{
				TypeUrl: "type.googleapis.com/google.protobuf.StringValue",
				Value:   []byte("test-config"),
			},
			Name:   "test",
			RootId: "test-root",
			VmType: wasm.WasmFilter_V8,
		}
		hl := &v1.HttpListener{
			Options: &v1.HttpListenerOptions{
				Wasm: &wasm.PluginSource{
					Filters: []*wasm.WasmFilter{wasmFilter},
				},
			},
		}

		mockCache.EXPECT().Add(gomock.Any(), image).Return(digest.Digest(sha), nil)
		f, err := p.HttpFilters(plugins.Params{}, hl)
		Expect(err).NotTo(HaveOccurred())
		Expect(f).To(HaveLen(1))
		goTypedConfig := f[0].HttpFilter.GetTypedConfig()
		var pc envoy_extensions_filters_http_wasm_v3.Wasm
		Expect(ptypes.UnmarshalAny(goTypedConfig, &pc)).NotTo(HaveOccurred())
		Expect(pc.Config.RootId).To(Equal(wasmFilter.RootId))
		Expect(pc.Config.Name).To(Equal(wasmFilter.Name))
		Expect(pc.Config.Configuration).To(matchers.MatchProto(wasmFilter.Config))
		Expect(pc.Config.GetVmConfig().GetVmId()).To(Equal(VmId))
		Expect(pc.Config.GetVmConfig().GetRuntime()).To(Equal(V8Runtime))
		remote := pc.Config.GetVmConfig().GetCode().GetRemote()
		Expect(remote).NotTo(BeNil())
		Expect(remote.Sha256).To(Equal(sha))
		Expect(remote.HttpUri.Uri).To(Equal(fmt.Sprintf("http://gloo/images/%s", sha)))
		Expect(remote.HttpUri.HttpUpstreamType).To(BeEquivalentTo(&configcore.HttpUri_Cluster{
			Cluster: WasmCacheCluster,
		}))
	})
	Context("filter stage transformations", func() {
		testCases := []struct {
			wasmFilterStage *wasm.FilterStage
			glooFilterStage plugins.FilterStage
		}{
			// Nil case
			{
				wasmFilterStage: nil,
				glooFilterStage: defaultPluginStage,
			},
			// Fault stage
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_FaultStage,
					Predicate: wasm.FilterStage_During,
				},
				glooFilterStage: plugins.DuringStage(plugins.FaultStage),
			},
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_FaultStage,
					Predicate: wasm.FilterStage_Before,
				},
				glooFilterStage: plugins.BeforeStage(plugins.FaultStage),
			},
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_FaultStage,
					Predicate: wasm.FilterStage_After,
				},
				glooFilterStage: plugins.AfterStage(plugins.FaultStage),
			},
			// Cors stage
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_CorsStage,
					Predicate: wasm.FilterStage_During,
				},
				glooFilterStage: plugins.DuringStage(plugins.CorsStage),
			},
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_CorsStage,
					Predicate: wasm.FilterStage_Before,
				},
				glooFilterStage: plugins.BeforeStage(plugins.CorsStage),
			},
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_CorsStage,
					Predicate: wasm.FilterStage_After,
				},
				glooFilterStage: plugins.AfterStage(plugins.CorsStage),
			},
			// Waf stage
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_WafStage,
					Predicate: wasm.FilterStage_During,
				},
				glooFilterStage: plugins.DuringStage(plugins.WafStage),
			},
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_WafStage,
					Predicate: wasm.FilterStage_Before,
				},
				glooFilterStage: plugins.BeforeStage(plugins.WafStage),
			},
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_WafStage,
					Predicate: wasm.FilterStage_After,
				},
				glooFilterStage: plugins.AfterStage(plugins.WafStage),
			},
			// AuthNstage
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_AuthNStage,
					Predicate: wasm.FilterStage_During,
				},
				glooFilterStage: plugins.DuringStage(plugins.AuthNStage),
			},
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_AuthNStage,
					Predicate: wasm.FilterStage_Before,
				},
				glooFilterStage: plugins.BeforeStage(plugins.AuthNStage),
			},
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_AuthNStage,
					Predicate: wasm.FilterStage_After,
				},
				glooFilterStage: plugins.AfterStage(plugins.AuthNStage),
			},
			// AuthZStage
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_AuthZStage,
					Predicate: wasm.FilterStage_During,
				},
				glooFilterStage: plugins.DuringStage(plugins.AuthZStage),
			},
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_AuthZStage,
					Predicate: wasm.FilterStage_Before,
				},
				glooFilterStage: plugins.BeforeStage(plugins.AuthZStage),
			},
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_AuthZStage,
					Predicate: wasm.FilterStage_After,
				},
				glooFilterStage: plugins.AfterStage(plugins.AuthZStage),
			},
			// RateLimit Stage
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_RateLimitStage,
					Predicate: wasm.FilterStage_During,
				},
				glooFilterStage: plugins.DuringStage(plugins.RateLimitStage),
			},
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_RateLimitStage,
					Predicate: wasm.FilterStage_Before,
				},
				glooFilterStage: plugins.BeforeStage(plugins.RateLimitStage),
			},
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_RateLimitStage,
					Predicate: wasm.FilterStage_After,
				},
				glooFilterStage: plugins.AfterStage(plugins.RateLimitStage),
			},
			// Accepted
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_AcceptedStage,
					Predicate: wasm.FilterStage_During,
				},
				glooFilterStage: plugins.DuringStage(plugins.AcceptedStage),
			},
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_AcceptedStage,
					Predicate: wasm.FilterStage_Before,
				},
				glooFilterStage: plugins.BeforeStage(plugins.AcceptedStage),
			},
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_AcceptedStage,
					Predicate: wasm.FilterStage_After,
				},
				glooFilterStage: plugins.AfterStage(plugins.AcceptedStage),
			},
			// OutAuth
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_OutAuthStage,
					Predicate: wasm.FilterStage_During,
				},
				glooFilterStage: plugins.DuringStage(plugins.OutAuthStage),
			},
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_OutAuthStage,
					Predicate: wasm.FilterStage_Before,
				},
				glooFilterStage: plugins.BeforeStage(plugins.OutAuthStage),
			},
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_OutAuthStage,
					Predicate: wasm.FilterStage_After,
				},
				glooFilterStage: plugins.AfterStage(plugins.OutAuthStage),
			},
			// Route
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_RouteStage,
					Predicate: wasm.FilterStage_During,
				},
				glooFilterStage: plugins.DuringStage(plugins.RouteStage),
			},
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_RouteStage,
					Predicate: wasm.FilterStage_Before,
				},
				glooFilterStage: plugins.BeforeStage(plugins.RouteStage),
			},
			{
				wasmFilterStage: &wasm.FilterStage{
					Stage:     wasm.FilterStage_RouteStage,
					Predicate: wasm.FilterStage_After,
				},
				glooFilterStage: plugins.AfterStage(plugins.RouteStage),
			},
		}

		It("can properly translate all test cases", func() {
			for _, v := range testCases {
				Expect(TransformWasmFilterStage(v.wasmFilterStage)).To(Equal(v.glooFilterStage))
			}
		})
	})
})
