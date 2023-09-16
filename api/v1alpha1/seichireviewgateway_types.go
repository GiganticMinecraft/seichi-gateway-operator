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

// SeichiReviewGatewaySpec defines the desired state of SeichiReviewGateway
type SeichiReviewGatewaySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// 稼働環境を生成する時にどのプルリクエスト番号時点のリポジトリデータを参照するかの設定
	PullRequestNo int `json:"pull-request-no"`
}

// SeichiReviewGatewayStatus defines the observed state of SeichiReviewGateway
type SeichiReviewGatewayStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SeichiReviewGateway is the Schema for the seichireviewgateways API
type SeichiReviewGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SeichiReviewGatewaySpec   `json:"spec,omitempty"`
	Status SeichiReviewGatewayStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SeichiReviewGatewayList contains a list of SeichiReviewGateway
type SeichiReviewGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SeichiReviewGateway `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SeichiReviewGateway{}, &SeichiReviewGatewayList{})
}
