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

package controller

import (
	"bytes"
	"context"
	"text/template"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	seichiclickv1alpha1 "github.com/GiganticMinecraft/seichi-gateway-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

// BungeeConfigTemplateReconciler reconciles a BungeeConfigTemplate object
type BungeeConfigTemplateReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=seichiclick.seichi.click,resources=bungeeconfigtemplates,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=seichiclick.seichi.click,resources=bungeeconfigtemplates/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=seichiclick.seichi.click,resources=bungeeconfigtemplates/finalizers,verbs=update
//+kubebuilder:rbac:groups=seichiclick.seichi.click,resources=seichireviewgateways,verbs=get;list;watch
//+kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the BungeeConfigTemplate object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.0/pkg/reconcile
func (r *BungeeConfigTemplateReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// seichiReviewGatewayList に seichireviewgateways のリストを格納する
	seichiReviewGatewayList := &seichiclickv1alpha1.SeichiReviewGatewayList{}
	if err := r.Client.List(ctx, seichiReviewGatewayList); err != nil {
		// Handle error
		return ctrl.Result{}, err
	}

	//  seichiReviewGatewayList と同じ長さの配列 reviewPullRequestNumberList を作る
	reviewPullRequestNumberList := make([]int, len(seichiReviewGatewayList.Items))

	// seichiReviewGatewayList.Items の要素1つ1つから、
	// seichiReviewGatewayList.Items[i].Spec.PullRequestNo の値を取り出して、
	// reviewPullRequestNumberList[i] に格納する
	for i, seichiReviewGateway := range seichiReviewGatewayList.Items {
		reviewPullRequestNumberList[i] = seichiReviewGateway.Spec.PullRequestNo
	}

	// BungeeConfigTemplateList に BungeeConfigTemplates のリストを格納する
	BungeeConfigTemplateList := &seichiclickv1alpha1.BungeeConfigTemplateList{}
	if err := r.Client.List(ctx, BungeeConfigTemplateList); err != nil {
		// Handle error
		return ctrl.Result{}, err
	}

	// BungeeConfigTemplateList.Items の要素1つ1つから、
	// BungeeConfigTemplateList.Items[i].Spec.BungeeConfigTemplate
	// の値を取り出す
	for _, BungeeConfigTemplate := range BungeeConfigTemplateList.Items {
		bungeeConfigTemplate := BungeeConfigTemplate.Spec.BungeeConfigTemplate

		// BungeeConfigTemplate.Status を更新する
		BungeeConfigTemplate.Status = seichiclickv1alpha1.BungeeConfigApplying

		// bungeeConfigTemplate に格納されているテンプレートをもとに、
		// go template を作成する
		tmp, err := template.New("template").Parse(bungeeConfigTemplate)
		if err != nil {
			BungeeConfigTemplate.Status = seichiclickv1alpha1.BungeeConfigError
			panic(err)
		}

		// go template に reviewPullRequestNumberList を適用して、
		// bungeeConfig に格納する
		var bungeeConfig bytes.Buffer
		if err := tmp.Execute(&bungeeConfig, reviewPullRequestNumberList); err != nil {
			BungeeConfigTemplate.Status = seichiclickv1alpha1.BungeeConfigError
			panic(err)
		}

		// Convert YAML to Kubernetes object
		var configMap corev1.ConfigMap
		yamlManifest := bungeeConfig.String()

		// YAML manifest as a string
		decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(yamlManifest)), len(yamlManifest))
		if err := decoder.Decode(&configMap); err != nil {
			panic(err.Error())
		}

		// Apply the ConfigMap to the cluster
		if err := r.Client.Update(ctx, &configMap); err != nil {
			// Handle error
			logger.Error(err, "unable to apply the ConfigMap to the cluster", "name", req.NamespacedName)
			BungeeConfigTemplate.Status = seichiclickv1alpha1.BungeeConfigError
			return ctrl.Result{}, err
		}

		// BungeeConfigTemplate.Status を更新する
		BungeeConfigTemplate.Status = seichiclickv1alpha1.BungeeConfigApplied
		if err := r.Client.Status().Update(ctx, &BungeeConfigTemplate); err != nil {
			// Handle error
			BungeeConfigTemplate.Status = seichiclickv1alpha1.BungeeConfigError
			return ctrl.Result{}, err
		}

	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *BungeeConfigTemplateReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&seichiclickv1alpha1.BungeeConfigTemplate{}).
		Complete(r)
}
