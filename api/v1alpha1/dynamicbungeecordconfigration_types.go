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

// DynamicBungeeCordConfigrationSpec defines the desired state of DynamicBungeeCordConfigration
type DynamicBungeeCordConfigrationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of DynamicBungeeCordConfigration. Edit dynamicbungeecordconfigration_types.go to remove/update
	// BungeeCordの設定のテンプレート
	BungeeConfigTemplate string `json:"bungeeConfigTemplate"`

	// apply先のconfigmapの情報 具体的には:
	// - namespace.metadata.name
	// - configmap.metadata.name
	// - configmap.data.??? <-どこか
	TargetToApplyNamespace     string `json:"targetToApplyNamespace"`
	TargetToApplyConfigMapName string `json:"targetToApplyConfigmapName"`
	TargetToApplyConfigMapKey  string `json:"targetToApplyConfigmapKey"`
}

// DynamicBungeeCordConfigrationStatus defines the observed state of DynamicBungeeCordConfigration
// +kubebuilder:validation:Enum=NotReady;Available;Healthy
type DynamicBungeeCordConfigrationStatus string

const (
	MarkdownViewApplying = DynamicBungeeCordConfigrationStatus("Applying")
	MarkdownViewApplied  = DynamicBungeeCordConfigrationStatus("Applied")
	MarkdownViewError    = DynamicBungeeCordConfigrationStatus("Error")
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DynamicBungeeCordConfigration is the Schema for the dynamicbungeecordconfigrations API
type DynamicBungeeCordConfigration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DynamicBungeeCordConfigrationSpec   `json:"spec,omitempty"`
	Status DynamicBungeeCordConfigrationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DynamicBungeeCordConfigrationList contains a list of DynamicBungeeCordConfigration
type DynamicBungeeCordConfigrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DynamicBungeeCordConfigration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DynamicBungeeCordConfigration{}, &DynamicBungeeCordConfigrationList{})
}
