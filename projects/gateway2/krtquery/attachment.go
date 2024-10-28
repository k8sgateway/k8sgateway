package krtquery

import (
	"errors"

	"github.com/golang/protobuf/ptypes/wrappers"
	"istio.io/istio/pkg/kube/krt"
	"istio.io/istio/pkg/ptr"
	"istio.io/istio/pkg/slices"
	"k8s.io/apimachinery/pkg/types"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

var (
	AttachmentErrTargetNotFound  = errors.New("Unresolved target or parent reference.")
	AttachmentErrSectionNotFound = errors.New("Unresolved section in target resource.")
)

// Attachment indexes a resource by its attachments.
// The mechanism for attachment can take several shapes:
// * TargetRef
// * ParentRef
// * Annotations/labels (like istio.io/use-waypoint)
type Attachment[A comparableNamespaced, T comparableNamespaced] struct {
	Resource A
	Target   T
	Section  string
	Error    error
}

func (a Attachment[A, T]) Equals(o Attachment[A, T]) bool {
	// don't compare the Target!
	return a.Resource == o.Resource && a.Error == o.Error && a.Section == o.Section
}

// ResourceName provides a unique key for krt collections.
func (h Attachment[A, T]) ResourceName() string {
	return h.Resource.GetNamespace() + "/" + h.Resource.GetName() +
		"~" +
		h.Target.GetNamespace() + "/" + h.Target.GetName() + "/" + h.Section
}

func ExtractResources[A comparableNamespaced, T comparableNamespaced](a []Attachment[A, T]) []A {
	return slices.Map(a, func(a Attachment[A, T]) A {
		return a.Resource
	})
}

// AttachmentSourceKey is based on both the source resource and target resource (and optionally section name).
type AttachmentSourceKey string

// MakeAttachmentSourceKey is used to key/query for individual attachments.
func MakeAttachmentSourceKey(resource Namespaced, target Namespaced, section string) AttachmentSourceKey {
	return AttachmentSourceKey(resource.GetNamespace() + "/" + resource.GetName() +
		"~" +
		target.GetNamespace() + "/" + target.GetName() + "/" + section)
}

// AttachmentKey is based on the target resource (and optionally section).
type AttachmentKey string

// MakeAttachmentKey is used to key/query for attachments to a given target resource.
func MakeAttachmentKey(target Namespaced, section string) AttachmentKey {
	return AttachmentKey(target.GetNamespace() + "/" + target.GetName() + "/" + section)
}

func (i *IndexedAttachments[A, T]) FetchOne(
	ctx krt.HandlerContext,
	resource Namespaced,
	target Namespaced,
	section string,
) *Attachment[A, T] {
	key := string(MakeAttachmentSourceKey(resource, target, section))
	return krt.FetchOne(ctx, i.Collection, krt.FilterKey(key))
}

func filterAttachedTo[A comparableNamespaced, T comparableNamespaced](target T, section string) krt.FetchOption {
	key := MakeAttachmentKey(target, section)
	return krt.FilterGeneric(func(o any) bool {
		a := o.(Attachment[A, T])
		return MakeAttachmentKey(a.Target, a.Section) == key
	})
}

func (i *IndexedAttachments[A, T]) FetchAttachmentsTo(
	ctx krt.HandlerContext,
	target T,
	section string,
) []Attachment[A, T] {
	var sectionFilter []krt.FetchOption
	if section != "" {
		sectionFilter = []krt.FetchOption{krt.FilterKey(string(MakeAttachmentKey(target, "")))}
	}
	return krt.Fetch(
		ctx,
		i.Collection,
		append(
			sectionFilter,
			krt.FilterIndex(i.Index, namespacedName(target)),
		)...,
	)
}

func (i *IndexedAttachments[A, T]) FetchErroredAttachments(
	ctx krt.HandlerContext,
	target T,
	section string,
) []Attachment[A, T] {
	var sectionFilter []krt.FetchOption
	if section != "" {
		sectionFilter = []krt.FetchOption{krt.FilterKey(string(MakeAttachmentKey(target, "")))}
	}
	return krt.Fetch(
		ctx,
		i.Collection,
		append(
			sectionFilter,
			krt.FilterIndex(i.ErrorIndex, namespacedName(target)),
		)...,
	)
}

type comparableNamespaced interface {
	comparable
	Namespaced
}

type IndexedAttachments[A comparableNamespaced, T comparableNamespaced] struct {
	Collection krt.Collection[Attachment[A, T]]
	Index      krt.Index[types.NamespacedName, Attachment[A, T]]
	ErrorIndex krt.Index[types.NamespacedName, Attachment[A, T]]
}

// CreateAttachmentIndexes allows consumers to query these attachments by the resource
// that's being targeted.
func CreateAttachmentIndexes[A comparableNamespaced, T comparableNamespaced](
	collection krt.Collection[Attachment[A, T]],
) IndexedAttachments[A, T] {
	return IndexedAttachments[A, T]{
		Collection: collection,
		Index: krt.NewIndex(collection, func(o Attachment[A, T]) []types.NamespacedName {
			if o.Error != nil {
				return nil
			}
			return []types.NamespacedName{namespacedName(o.Target)}
		}),
		ErrorIndex: krt.NewIndex(collection, func(o Attachment[A, T]) []types.NamespacedName {
			if o.Error == nil {
				return nil
			}
			return []types.NamespacedName{namespacedName(o.Target)}
		}),
	}
}

// common struct for parentRef and targetRef
type ref struct {
	ns, name, section string
}

// targetResourceKey gives the target without section name
func (r ref) targetResourceKey() string {
	return r.ns + "/" + r.name
}

func (r ref) GetName() string {
	return r.name
}

func (r ref) GetNamespace() string {
	return r.ns
}

func parentRef[T Namespaced](resource T, parentRef gwv1.ParentReference) ref {
	return ref{
		ns:      string(ptr.OrDefault(parentRef.Namespace, gwv1.Namespace(resource.GetNamespace()))),
		name:    string(parentRef.Name),
		section: string(ptr.OrEmpty(parentRef.SectionName)),
	}
}

type TargetRef interface {
	GetNamespace() *wrappers.StringValue
	GetName() string
	GetSectionName() *wrappers.StringValue
}

func targetRef[T Namespaced](resource T, targetRef TargetRef) ref {
	return ref{
		ns:      ptr.NonEmptyOrDefault(targetRef.GetNamespace().GetValue(), resource.GetNamespace()),
		name:    string(targetRef.GetName()),
		section: string(ptr.OrEmpty(targetRef.GetSectionName()).Value),
	}
}

func NewGatewayAttachment[T comparableNamespaced](
	ctx krt.HandlerContext,
	Gateways krt.Collection[*gwv1.Gateway],
	resource T,
	ref ref,
) (Attachment[T, *gwv1.Gateway], error) {
	gw := ptr.Flatten(krt.FetchOne(ctx, Gateways, krt.FilterKey(ref.targetResourceKey())))
	// gw not found
	if gw == nil {
		return Attachment[T, *gwv1.Gateway]{
			Resource: resource, Target: nil, Section: ref.section,
			Error: AttachmentErrTargetNotFound,
		}, AttachmentErrTargetNotFound
	}
	// no need for section
	if ref.section == "" {
		return Attachment[T, *gwv1.Gateway]{Resource: resource, Target: gw}, nil
	}
	// try find section
	for _, l := range gw.Spec.Listeners {
		if string(l.Name) == ref.section {
			return Attachment[T, *gwv1.Gateway]{Resource: resource, Target: gw, Section: ref.section}, nil
		}
	}
	// no section found
	return Attachment[T, *gwv1.Gateway]{
		Resource: resource, Target: gw, Section: ref.section,
		Error: AttachmentErrSectionNotFound,
	}, AttachmentErrSectionNotFound
}
