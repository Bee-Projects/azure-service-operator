package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


// AcrSpec defines the desired state of Acr
type AcrSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	ResourceGroupName string `json:"resource_group_name"`
	Location string `json:"location"`
	Sku string `json:"sku"`
	AdminEnabled bool `json:"admin_enabled"`
}

// AcrStatus defines the observed state of Acr
type AcrStatus struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	ResourceGroupName string `json:"resource_group_name"`
	Location string `json:"location"`
	Sku string `json:"sku"`
	AdminEnabled bool `json:"admin_enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Acr is the Schema for the acrs API
// +k8s:openapi-gen=true
type Acr struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AcrSpec   `json:"spec,omitempty"`
	Status AcrStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AcrList contains a list of Acr
type AcrList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Acr `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Acr{}, &AcrList{})
}
