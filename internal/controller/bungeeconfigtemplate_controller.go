package controller

import (
	"bytes"
	"context"
	"errors"
	"text/template"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	seichiclickv1alpha1 "github.com/GiganticMinecraft/seichi-gateway-operator/api/v1alpha1"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
)

// BungeeConfigTemplateReconciler reconciles a BungeeConfigTemplate object
type BungeeConfigTemplateReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=seichi.click,resources=bungeeconfigtemplates,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=seichi.click,resources=bungeeconfigtemplates/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=seichi.click,resources=bungeeconfigtemplates/finalizers,verbs=update
//+kubebuilder:rbac:groups=seichi.click,resources=seichireviewgateways,verbs=get;list;watch
//+kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.0/pkg/reconcile
func (r *BungeeConfigTemplateReconciler) Reconcile(ctx context.Context, _ ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	seichiReviewGatewayList := &seichiclickv1alpha1.SeichiReviewGatewayList{}
	if err := r.Client.List(ctx, seichiReviewGatewayList); err != nil {
		return ctrl.Result{}, err
	}

	pullRequestNumbers := lo.Map(seichiReviewGatewayList.Items, func(item seichiclickv1alpha1.SeichiReviewGateway, index int) int {
		return item.Spec.PullRequestNo
	})

	bungeeConfigTemplates := &seichiclickv1alpha1.BungeeConfigTemplateList{}
	if err := r.Client.List(ctx, bungeeConfigTemplates); err != nil {
		return ctrl.Result{}, err
	}

	for _, bungeeConfigTemplate := range bungeeConfigTemplates.Items {
		setErrorStatusToTemplateResourceAndReturnBecauseOf := func(err error) (ctrl.Result, error) {
			logger.Error(
				err, "unable to reconcile the config template",
				"template", bungeeConfigTemplate,
			)

			bungeeConfigTemplate.Status = seichiclickv1alpha1.BungeeConfigError
			if updateErr := r.Client.Status().Update(ctx, &bungeeConfigTemplate); err != nil {
				return ctrl.Result{}, errors.Join(err, updateErr)
			} else {
				return ctrl.Result{}, err
			}
		}

		templateObject, err := template.New("template").Parse(bungeeConfigTemplate.Spec.BungeeConfigTemplate)
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

		// Manifest 文字列を Kubernetes object へと変換する
		var configMap corev1.ConfigMap
		decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(bungeeConfigMapManifest.Bytes()), bungeeConfigMapManifest.Len())
		if err := decoder.Decode(&configMap); err != nil {
			return setErrorStatusToTemplateResourceAndReturnBecauseOf(err)
		}

		// クラスタに ConfigMap を適用する
		if err := r.Client.Update(ctx, &configMap); err != nil {
			return setErrorStatusToTemplateResourceAndReturnBecauseOf(err)
		}

		// 適用が完了したらステータスを更新する
		bungeeConfigTemplate.Status = seichiclickv1alpha1.BungeeConfigApplied
		if err := r.Client.Status().Update(ctx, &bungeeConfigTemplate); err != nil {
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

// SetupWithManager sets up the controller with the Manager.
func (r *BungeeConfigTemplateReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&seichiclickv1alpha1.BungeeConfigTemplate{}).
		Complete(r)
}
