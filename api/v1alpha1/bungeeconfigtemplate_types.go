package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BungeeConfigTemplateSpec struct {
	// BungeeCordのconfig.yamlを含むConfigMapマニフェスト。go-templateで記載されている
	BungeeConfigTemplate string `json:"goTemplate"`
}

// +kubebuilder:validation:Enum=Applied;Error

type BungeeConfigTemplateStatus string

const (
	BungeeConfigApplied = BungeeConfigTemplateStatus("Applied")
	BungeeConfigError   = BungeeConfigTemplateStatus("Error")
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

type BungeeConfigTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BungeeConfigTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BungeeConfigTemplate{}, &BungeeConfigTemplateList{})
}
