package krtcollections_test

import (
	"context"
	"testing"

	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gateway2/krtcollections"
	. "github.com/solo-io/gloo/projects/gateway2/krtcollections"
	"istio.io/istio/pkg/kube/krt"
	"istio.io/istio/pkg/kube/krt/krttest"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestPods(t *testing.T) {
	g := gomega.NewWithT(t)
	testCases := []struct {
		name   string
		inputs []any
		result LocalityPod
	}{
		{
			name: "basic",
			inputs: []any{
				&corev1.Pod{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "name",
						Namespace: "ns",
						Labels:    map[string]string{"a": "b"},
					},
					Spec: corev1.PodSpec{
						NodeName: "node",
					},
				},
				&corev1.Node{
					ObjectMeta: metav1.ObjectMeta{
						Name: "node",
						Labels: map[string]string{
							corev1.LabelTopologyRegion: "region",
							corev1.LabelTopologyZone:   "zone",
						},
					},
				},
			},
			result: LocalityPod{
				Named: krt.Named{
					Name:      "name",
					Namespace: "ns",
				},
				Locality: PodLocality{
					Region:  "region",
					Zone:    "zone",
					Subzone: "",
				},
				AugmentedLabels: map[string]string{
					corev1.LabelTopologyRegion: "region",
					corev1.LabelTopologyZone:   "zone",
					"a":                        "b",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := krttest.NewMock(t, tc.inputs)
			nodes := krtcollections.NewNodeMetadataCollection(krttest.GetMockCollection[*corev1.Node](mock))
			pods := krtcollections.NewLocalityPodsCollection(nodes, krttest.GetMockCollection[*corev1.Pod](mock))
			pods.Synced().WaitUntilSynced(context.Background().Done())
			lp := pods.List()[0]

			g.Expect(tc.result.Equals(lp)).To(BeTrue(), "expected %#v, got %#v", lp, tc.result)
		})
	}

}
