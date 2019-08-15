package benchmark_test

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"hash"
	"hash/fnv"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/go-utils/protoutils"
)

var _ = Describe("SnapshotBenchmark", func() {
	var allUpstreams v1.UpstreamList
	BeforeEach(func() {
		var upstreams []interface{}
		data := helpers.MustReadFile("upstream_list.json")
		err := json.Unmarshal(data, &upstreams)
		Expect(err).NotTo(HaveOccurred())
		for _, usInt := range upstreams {
			usMap := usInt.(map[string]interface{})
			var us v1.Upstream
			err = protoutils.UnmarshalMap(usMap, &us)
			Expect(err).NotTo(HaveOccurred())
			allUpstreams = append(allUpstreams, &us)
		}

		// Sort merged slice for consistent hashing
		allUpstreams.Sort()

	})
	Measure("it should do something hard efficiently", func(b Benchmarker) {
		const times = 100
		runtime1 := b.Time(fmt.Sprintf("runtime of %d reflect-based hash calls", times), func() {
			for i := 0; i < times; i++ {
				for _, us := range allUpstreams {
					us.Hash()
				}
			}
		})
		runtime2 := b.Time(fmt.Sprintf("runtime of %d generated hash calls", times), func() {
			for i := 0; i < times; i++ {
				for _, us := range allUpstreams {
					Hashi(us)
				}
			}
		})

		// divide by 1e3 to get time in micro seconds instead of nano seconds
		b.RecordValue("Runtime per reflection call in µ seconds", float64(int64(runtime1)/times)/1e3)
		b.RecordValue("Runtime per generated call in µ seconds", float64(int64(runtime2)/times)/1e3)

	}, 10)
})

type pair struct {
	k, v string
}

func Hashi(us *v1.Upstream) uint64 {
	pairs := []pair{
		{k: "us.Metadata.Namespace", v: us.Metadata.Namespace},
		{k: "us.Metadata.Name", v: us.Metadata.Name},
		{k: "us.UpstreamSpec.GetKube().ServiceNamespace", v: us.UpstreamSpec.GetKube().ServiceNamespace},
		{k: "us.UpstreamSpec.GetKube().ServiceName", v: us.UpstreamSpec.GetKube().ServiceName},
		{k: "strconv.Itoa(int(us.UpstreamSpec.GetKube().ServicePort))", v: strconv.Itoa(int(us.UpstreamSpec.GetKube().ServicePort))},
	}

	for k, v := range us.Metadata.Annotations {
		pairs = append(pairs, pair{k, v})
	}
	for k, v := range us.Metadata.Labels {
		pairs = append(pairs, pair{k, v})
	}

	h := fnv.New64()
	var hash uint64
	for _, p := range pairs {
		pairHash := hashPair(h, p)
		hash = hashUpdateUnordered(hash, pairHash)
	}
	return hash
}

func hashPair(h hash.Hash64, p pair) uint64 {
	k := hashString(h, p.k)
	v := hashString(h, p.v)
	return hashUpdateOrdered(h, k, v)
}

func hashUpdateOrdered(h hash.Hash64, a, b uint64) uint64 {
	// For ordered updates, use a real hash function
	h.Reset()

	// We just panic if the binary writes fail because we are writing
	// an int64 which should never be fail-able.
	e1 := binary.Write(h, binary.LittleEndian, a)
	e2 := binary.Write(h, binary.LittleEndian, b)
	if e1 != nil {
		panic(e1)
	}
	if e2 != nil {
		panic(e2)
	}

	return h.Sum64()
}

func hashUpdateUnordered(a, b uint64) uint64 {
	return a ^ b
}

func hashString(h hash.Hash64, s string) uint64 {
	h.Reset()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}
