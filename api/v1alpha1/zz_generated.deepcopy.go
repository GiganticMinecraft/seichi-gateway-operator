//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BungeeConfigMapTemplate) DeepCopyInto(out *BungeeConfigMapTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BungeeConfigMapTemplate.
func (in *BungeeConfigMapTemplate) DeepCopy() *BungeeConfigMapTemplate {
	if in == nil {
		return nil
	}
	out := new(BungeeConfigMapTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BungeeConfigMapTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BungeeConfigMapTemplateList) DeepCopyInto(out *BungeeConfigMapTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BungeeConfigMapTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BungeeConfigMapTemplateList.
func (in *BungeeConfigMapTemplateList) DeepCopy() *BungeeConfigMapTemplateList {
	if in == nil {
		return nil
	}
	out := new(BungeeConfigMapTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BungeeConfigMapTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BungeeConfigMapTemplateSpec) DeepCopyInto(out *BungeeConfigMapTemplateSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BungeeConfigMapTemplateSpec.
func (in *BungeeConfigMapTemplateSpec) DeepCopy() *BungeeConfigMapTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(BungeeConfigMapTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeichiAssistDebugEnvironmentRequest) DeepCopyInto(out *SeichiAssistDebugEnvironmentRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeichiAssistDebugEnvironmentRequest.
func (in *SeichiAssistDebugEnvironmentRequest) DeepCopy() *SeichiAssistDebugEnvironmentRequest {
	if in == nil {
		return nil
	}
	out := new(SeichiAssistDebugEnvironmentRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SeichiAssistDebugEnvironmentRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeichiAssistDebugEnvironmentRequestList) DeepCopyInto(out *SeichiAssistDebugEnvironmentRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SeichiAssistDebugEnvironmentRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeichiAssistDebugEnvironmentRequestList.
func (in *SeichiAssistDebugEnvironmentRequestList) DeepCopy() *SeichiAssistDebugEnvironmentRequestList {
	if in == nil {
		return nil
	}
	out := new(SeichiAssistDebugEnvironmentRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SeichiAssistDebugEnvironmentRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeichiAssistDebugEnvironmentRequestSpec) DeepCopyInto(out *SeichiAssistDebugEnvironmentRequestSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeichiAssistDebugEnvironmentRequestSpec.
func (in *SeichiAssistDebugEnvironmentRequestSpec) DeepCopy() *SeichiAssistDebugEnvironmentRequestSpec {
	if in == nil {
		return nil
	}
	out := new(SeichiAssistDebugEnvironmentRequestSpec)
	in.DeepCopyInto(out)
	return out
}
