package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BungeeConfigTemplateSpec defines the desired state of BungeeConfigTemplate
type BungeeConfigTemplateSpec struct {

	// BungeeCordのconfig.yamlを含むConfigMapマニフェスト。go-templateで記載されている
	BungeeConfigTemplate string `json:"goTemplate"`
}

// BungeeConfigTemplateStatus defines the observed state of BungeeConfigTemplate
// +kubebuilder:validation:Enum=Applying;Applied;Error
type BungeeConfigTemplateStatus string

const (
	BungeeConfigApplying = BungeeConfigTemplateStatus("Applying")
	BungeeConfigApplied  = BungeeConfigTemplateStatus("Applied")
	BungeeConfigError    = BungeeConfigTemplateStatus("Error")
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// BungeeConfigTemplate is the Schema for the BungeeConfigTemplates API
type BungeeConfigTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BungeeConfigTemplateSpec   `json:"spec,omitempty"`
	Status BungeeConfigTemplateStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BungeeConfigTemplateList contains a list of BungeeConfigTemplate
type BungeeConfigTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BungeeConfigTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BungeeConfigTemplate{}, &BungeeConfigTemplateList{})
}
