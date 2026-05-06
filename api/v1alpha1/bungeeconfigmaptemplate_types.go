/*
Copyright 2026.

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

// BungeeConfigMapTemplateSpec defines the desired state of BungeeConfigMapTemplate
type BungeeConfigMapTemplateSpec struct {
	// BungeeCord の config.yaml を含む ConfigMap.spec.data を生成するための Go Template 文字列
	ConfigMapDataTemplate string `json:"dataGoTemplate"`
}

// +kubebuilder:validation:Enum=Applied;Error

// BungeeConfigMapTemplateStatus は適用結果を示す string enum.
type BungeeConfigMapTemplateStatus string

const (
	BungeeConfigMapTemplateApplied = BungeeConfigMapTemplateStatus("Applied")
	BungeeConfigMapTemplateError   = BungeeConfigMapTemplateStatus("Error")
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// BungeeConfigMapTemplate is the Schema for the bungeeconfigmaptemplates API
type BungeeConfigMapTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BungeeConfigMapTemplateSpec   `json:"spec,omitempty"`
	Status BungeeConfigMapTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BungeeConfigMapTemplateList contains a list of BungeeConfigMapTemplate
type BungeeConfigMapTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BungeeConfigMapTemplate `json:"items"`
}
