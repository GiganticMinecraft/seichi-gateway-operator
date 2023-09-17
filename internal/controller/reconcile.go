package controller

import (
	"bytes"
	"context"
	"errors"
	seichiclickv1alpha1 "github.com/GiganticMinecraft/seichi-gateway-operator/api/v1alpha1"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"text/template"
)

//+kubebuilder:rbac:groups=seichi.click,resources=seichiassistdebugenvrequests,verbs=get;list;watch
//+kubebuilder:rbac:groups=seichi.click,resources=seichiassistdebugenvrequests/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=seichi.click,resources=seichiassistdebugenvrequests/finalizers,verbs=update
//+kubebuilder:rbac:groups=seichi.click,resources=bungeeconfigmaptemplates,verbs=get;list;watch;update;patch
//+kubebuilder:rbac:groups=seichi.click,resources=bungeeconfigmaptemplates/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=seichi.click,resources=bungeeconfigmaptemplates/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;create;update;patch

func ReconcileAllManagedResources(ctx context.Context, k8sClient client.Client) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	seichiReviewGatewayList := &seichiclickv1alpha1.SeichiAssistDebugEnvRequestList{}
	if err := k8sClient.List(ctx, seichiReviewGatewayList); err != nil {
		return ctrl.Result{}, err
	}

	pullRequestNumbers := lo.Map(seichiReviewGatewayList.Items, func(item seichiclickv1alpha1.SeichiAssistDebugEnvRequest, index int) int {
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
			if updateErr := k8sClient.Status().Update(ctx, &bungeeConfigTemplate); err != nil {
				return ctrl.Result{}, errors.Join(err, updateErr)
			} else {
				return ctrl.Result{}, err
			}
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
		configMap := &corev1.ConfigMap{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "v1",
				Kind:       "ConfigMap",
			},
			ObjectMeta: metav1.ObjectMeta{
				Namespace: bungeeConfigTemplate.Namespace,
				Name:      bungeeConfigTemplate.Name,
				OwnerReferences: []metav1.OwnerReference{
					{
						APIVersion: bungeeConfigTemplate.APIVersion,
						Kind:       bungeeConfigTemplate.Kind,
						Name:       bungeeConfigTemplate.Name,
						UID:        bungeeConfigTemplate.UID,
					},
				},
			},
			Data: configMapData,
		}

		// クラスタに ConfigMap を apply する
		if err := k8sClient.Patch(
			ctx, configMap, client.Apply,
			client.ForceOwnership, client.FieldOwner("seichi-gateway-operator"),
		); err != nil {
			return setErrorStatusToTemplateResourceAndReturnBecauseOf(err)
		}

		// 適用が完了したらステータスを更新する
		bungeeConfigTemplate.Status = seichiclickv1alpha1.BungeeConfigMapTemplateApplied
		if err := k8sClient.Status().Update(ctx, &bungeeConfigTemplate); err != nil {
			// 適用には成功しているので Status は Applied のままにしておく
			return ctrl.Result{}, err
		}

		logger.Info(
			"Applied the ConfigMap to the cluster",
			"template", bungeeConfigTemplate,
		)
	}

	return ctrl.Result{}, nil
}
