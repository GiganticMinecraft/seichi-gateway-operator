package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SeichiAssistDebugEnvironmentRequestSpec struct {
	// SeichiAssist リポジトリの Pull Request で Ready for review になっていて
	// デバッグ環境を必要としているものの番号
	PullRequestNo int `json:"pullRequestNo"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

type SeichiAssistDebugEnvironmentRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec SeichiAssistDebugEnvironmentRequestSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

type SeichiAssistDebugEnvironmentRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SeichiAssistDebugEnvironmentRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SeichiAssistDebugEnvironmentRequest{}, &SeichiAssistDebugEnvironmentRequestList{})
}
