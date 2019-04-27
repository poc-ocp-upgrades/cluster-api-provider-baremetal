package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *BareMetalMachineProviderSpec) DeepCopyInto(out *BareMetalMachineProviderSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}
func (in *BareMetalMachineProviderSpec) DeepCopy() *BareMetalMachineProviderSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BareMetalMachineProviderSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *BareMetalMachineProviderSpec) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *BareMetalMachineProviderSpecList) DeepCopyInto(out *BareMetalMachineProviderSpecList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BareMetalMachineProviderSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *BareMetalMachineProviderSpecList) DeepCopy() *BareMetalMachineProviderSpecList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BareMetalMachineProviderSpecList)
	in.DeepCopyInto(out)
	return out
}
func (in *BareMetalMachineProviderSpecList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *BareMetalMachineProviderSpecSpec) DeepCopyInto(out *BareMetalMachineProviderSpecSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *BareMetalMachineProviderSpecSpec) DeepCopy() *BareMetalMachineProviderSpecSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BareMetalMachineProviderSpecSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *BareMetalMachineProviderSpecStatus) DeepCopyInto(out *BareMetalMachineProviderSpecStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *BareMetalMachineProviderSpecStatus) DeepCopy() *BareMetalMachineProviderSpecStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BareMetalMachineProviderSpecStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *BareMetalMachineProviderStatus) DeepCopyInto(out *BareMetalMachineProviderStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}
func (in *BareMetalMachineProviderStatus) DeepCopy() *BareMetalMachineProviderStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BareMetalMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *BareMetalMachineProviderStatus) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *BareMetalMachineProviderStatusList) DeepCopyInto(out *BareMetalMachineProviderStatusList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BareMetalMachineProviderStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *BareMetalMachineProviderStatusList) DeepCopy() *BareMetalMachineProviderStatusList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BareMetalMachineProviderStatusList)
	in.DeepCopyInto(out)
	return out
}
func (in *BareMetalMachineProviderStatusList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *BareMetalMachineProviderStatusSpec) DeepCopyInto(out *BareMetalMachineProviderStatusSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *BareMetalMachineProviderStatusSpec) DeepCopy() *BareMetalMachineProviderStatusSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BareMetalMachineProviderStatusSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *BareMetalMachineProviderStatusStatus) DeepCopyInto(out *BareMetalMachineProviderStatusStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *BareMetalMachineProviderStatusStatus) DeepCopy() *BareMetalMachineProviderStatusStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BareMetalMachineProviderStatusStatus)
	in.DeepCopyInto(out)
	return out
}
