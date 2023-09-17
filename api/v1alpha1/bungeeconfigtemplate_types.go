/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BungeeConfigTemplateSpec defines the desired state of BungeeConfigTemplate
type BungeeConfigTemplateSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

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
