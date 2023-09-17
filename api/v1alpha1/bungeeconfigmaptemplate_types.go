package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BungeeConfigMapTemplateSpec struct {
	// BungeeCord の config.yaml を含む ConfigMap.spec.data を生成するための Go Template 文字列
	ConfigMapDataTemplate string `json:"dataGoTemplate"`
}

// +kubebuilder:validation:Enum=Applied;Error

type BungeeConfigMapTemplateStatus string

const (
	BungeeConfigMapTemplateApplied = BungeeConfigMapTemplateStatus("Applied")
	BungeeConfigMapTemplateError   = BungeeConfigMapTemplateStatus("Error")
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

type BungeeConfigMapTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BungeeConfigMapTemplateSpec   `json:"spec,omitempty"`
	Status BungeeConfigMapTemplateStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

type BungeeConfigMapTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BungeeConfigMapTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BungeeConfigMapTemplate{}, &BungeeConfigMapTemplateList{})
}
