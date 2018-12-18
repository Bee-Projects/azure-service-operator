package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StorageAccountSpec defines the desired state of StorageAccount
type StorageAccountSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	ResourceGroupName string `json:"resource_group_name"`
	Location string `json:"location"`
	AccountType string `json:"account_type"`
	Kind string `json:"kind"`
	AccessTier string `json:"access_tier"`
	SupportsHttpsTrafficOnly bool `json:"supports_https_traffic_only"`
}

// StorageAccountStatus defines the observed state of StorageAccount
type StorageAccountStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StorageAccount is the Schema for the storageaccounts API
// +k8s:openapi-gen=true
type StorageAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageAccountSpec   `json:"spec,omitempty"`
	Status StorageAccountStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StorageAccountList contains a list of StorageAccount
type StorageAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccount `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageAccount{}, &StorageAccountList{})
}
