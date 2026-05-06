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

package controller

import (
	"bytes"
	"context"
	"errors"
	"text/template"

	"github.com/samber/lo"
	"k8s.io/apimachinery/pkg/util/yaml"
	corev1ac "k8s.io/client-go/applyconfigurations/core/v1"
	metav1ac "k8s.io/client-go/applyconfigurations/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	seichiclickv1alpha1 "github.com/GiganticMinecraft/seichi-gateway-operator/api/v1alpha1"
)

// +kubebuilder:rbac:groups=seichi.click,resources=seichiassistdebugenvrequests,verbs=get;list;watch
// +kubebuilder:rbac:groups=seichi.click,resources=seichiassistdebugenvrequests/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=seichi.click,resources=seichiassistdebugenvrequests/finalizers,verbs=update
// +kubebuilder:rbac:groups=seichi.click,resources=bungeeconfigmaptemplates,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=seichi.click,resources=bungeeconfigmaptemplates/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=seichi.click,resources=bungeeconfigmaptemplates/finalizers,verbs=update
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;create;update;patch

func ReconcileAllManagedResources(ctx context.Context, k8sClient client.Client) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	debugEnvRequestList := &seichiclickv1alpha1.SeichiAssistDebugEnvRequestList{}
	if err := k8sClient.List(ctx, debugEnvRequestList); err != nil {
		return ctrl.Result{}, err
	}

	pullRequestNumbers := lo.Map(debugEnvRequestList.Items, func(item seichiclickv1alpha1.SeichiAssistDebugEnvRequest, _ int) int {
		return item.Spec.PullRequestNo
	})

	bungeeConfigTemplates := &seichiclickv1alpha1.BungeeConfigMapTemplateList{}
	if err := k8sClient.List(ctx, bungeeConfigTemplates); err != nil {
		return ctrl.Result{}, err
	}

	for _, bungeeConfigTemplate := range bungeeConfigTemplates.Items {
		setErrorStatusToTemplateResourceAndReturnBecauseOf := func(err error) (ctrl.Result, error) {
			logger.Error(
				err, "unable to reconcile the config template",
				"template", bungeeConfigTemplate,
			)

			bungeeConfigTemplate.Status = seichiclickv1alpha1.BungeeConfigMapTemplateError
			if updateErr := k8sClient.Status().Update(ctx, &bungeeConfigTemplate); updateErr != nil {
				return ctrl.Result{}, errors.Join(err, updateErr)
			}
			return ctrl.Result{}, err
		}

		templateObject, err := template.New("template").Parse(bungeeConfigTemplate.Spec.ConfigMapDataTemplate)
		if err != nil {
			return setErrorStatusToTemplateResourceAndReturnBecauseOf(err)
		}

		// templateObject を pullRequestNumbers を渡して展開したものを BungeeCord の ConfigMap を定義するマニフェストとして扱う
		var bungeeConfigMapManifest bytes.Buffer
		if err := templateObject.Execute(&bungeeConfigMapManifest, pullRequestNumbers); err != nil {
			return setErrorStatusToTemplateResourceAndReturnBecauseOf(err)
		}

		logger.Info(
			"Trying to apply the ConfigMap to the cluster",
			"manifest string", bungeeConfigMapManifest.String(),
		)

		// Manifest 文字列を ConfigMap の data として扱う
		var configMapData map[string]string
		decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(bungeeConfigMapManifest.Bytes()), bungeeConfigMapManifest.Len())
		if err := decoder.Decode(&configMapData); err != nil {
			return setErrorStatusToTemplateResourceAndReturnBecauseOf(err)
		}

		configMapApply := corev1ac.ConfigMap(bungeeConfigTemplate.Name, bungeeConfigTemplate.Namespace).
			WithOwnerReferences(metav1ac.OwnerReference().
				WithAPIVersion(bungeeConfigTemplate.APIVersion).
				WithKind(bungeeConfigTemplate.Kind).
				WithName(bungeeConfigTemplate.Name).
				WithUID(bungeeConfigTemplate.UID)).
			WithData(configMapData)

		// クラスタに ConfigMap を apply する (Server-Side Apply)
		if err := k8sClient.Apply(
			ctx, configMapApply,
			client.ForceOwnership, client.FieldOwner("seichi-gateway-operator"),
		); err != nil {
			return setErrorStatusToTemplateResourceAndReturnBecauseOf(err)
		}

		// 適用が完了したらステータスを更新する
		bungeeConfigTemplate.Status = seichiclickv1alpha1.BungeeConfigMapTemplateApplied
		if err := k8sClient.Status().Update(ctx, &bungeeConfigTemplate); err != nil {
			return ctrl.Result{}, err
		}

		logger.Info(
			"Applied the ConfigMap to the cluster",
			"template", bungeeConfigTemplate,
		)
	}

	return ctrl.Result{}, nil
}
