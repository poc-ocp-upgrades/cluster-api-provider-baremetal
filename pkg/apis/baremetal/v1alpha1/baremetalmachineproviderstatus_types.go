package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BareMetalMachineProviderStatusSpec struct{}
type BareMetalMachineProviderStatusStatus struct{}
type BareMetalMachineProviderStatus struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec				BareMetalMachineProviderStatusSpec		`json:"spec,omitempty"`
	Status				BareMetalMachineProviderStatusStatus	`json:"status,omitempty"`
}
type BareMetalMachineProviderStatusList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata,omitempty"`
	Items			[]BareMetalMachineProviderStatus	`json:"items"`
}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	SchemeBuilder.Register(&BareMetalMachineProviderStatus{}, &BareMetalMachineProviderStatusList{})
}
