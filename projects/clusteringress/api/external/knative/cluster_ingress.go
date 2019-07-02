package knative

import (
	"reflect"

	"github.com/knative/serving/pkg/apis/networking/v1alpha1"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/utils/kubeutils"
)

type ClusterIngress v1alpha1.ClusterIngress

func (p *ClusterIngress) GetMetadata() core.Metadata {
	return kubeutils.FromKubeMeta(p.ObjectMeta)
}

func (p *ClusterIngress) SetMetadata(meta core.Metadata) {
	p.ObjectMeta = kubeutils.ToKubeMeta(meta)
}

func (p *ClusterIngress) Equal(that interface{}) bool {
	return reflect.DeepEqual(p, that)
}

func (p *ClusterIngress) Clone() *ClusterIngress {
	ci := v1alpha1.ClusterIngress(*p)
	copy := ci.DeepCopy()
	newCi := ClusterIngress(*copy)
	return &newCi
}
