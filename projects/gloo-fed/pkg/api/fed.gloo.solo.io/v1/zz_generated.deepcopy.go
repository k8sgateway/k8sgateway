// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for fed.gloo.solo.io/v1 resources

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for FederatedUpstream

func (in *FederatedUpstream) DeepCopyInto(out *FederatedUpstream) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *FederatedUpstream) DeepCopy() *FederatedUpstream {
	if in == nil {
		return nil
	}
	out := new(FederatedUpstream)
	in.DeepCopyInto(out)
	return out
}

func (in *FederatedUpstream) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *FederatedUpstreamList) DeepCopyInto(out *FederatedUpstreamList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedUpstream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *FederatedUpstreamList) DeepCopy() *FederatedUpstreamList {
	if in == nil {
		return nil
	}
	out := new(FederatedUpstreamList)
	in.DeepCopyInto(out)
	return out
}

func (in *FederatedUpstreamList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for FederatedUpstreamGroup

func (in *FederatedUpstreamGroup) DeepCopyInto(out *FederatedUpstreamGroup) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *FederatedUpstreamGroup) DeepCopy() *FederatedUpstreamGroup {
	if in == nil {
		return nil
	}
	out := new(FederatedUpstreamGroup)
	in.DeepCopyInto(out)
	return out
}

func (in *FederatedUpstreamGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *FederatedUpstreamGroupList) DeepCopyInto(out *FederatedUpstreamGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedUpstreamGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *FederatedUpstreamGroupList) DeepCopy() *FederatedUpstreamGroupList {
	if in == nil {
		return nil
	}
	out := new(FederatedUpstreamGroupList)
	in.DeepCopyInto(out)
	return out
}

func (in *FederatedUpstreamGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// Generated Deepcopy methods for FederatedSettings

func (in *FederatedSettings) DeepCopyInto(out *FederatedSettings) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)
	// deepcopy status
	in.Status.DeepCopyInto(&out.Status)

	return
}

func (in *FederatedSettings) DeepCopy() *FederatedSettings {
	if in == nil {
		return nil
	}
	out := new(FederatedSettings)
	in.DeepCopyInto(out)
	return out
}

func (in *FederatedSettings) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *FederatedSettingsList) DeepCopyInto(out *FederatedSettingsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedSettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *FederatedSettingsList) DeepCopy() *FederatedSettingsList {
	if in == nil {
		return nil
	}
	out := new(FederatedSettingsList)
	in.DeepCopyInto(out)
	return out
}

func (in *FederatedSettingsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
