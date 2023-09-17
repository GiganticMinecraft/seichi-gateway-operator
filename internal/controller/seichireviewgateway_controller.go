package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	seichiclickv1alpha1 "github.com/GiganticMinecraft/seichi-gateway-operator/api/v1alpha1"
)

// SeichiReviewGatewayReconciler reconciles a SeichiReviewGateway object
type SeichiReviewGatewayReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=seichi.click,resources=seichireviewgateways,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=seichi.click,resources=seichireviewgateways/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=seichi.click,resources=seichireviewgateways/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.0/pkg/reconcile
func (r *SeichiReviewGatewayReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SeichiReviewGatewayReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&seichiclickv1alpha1.SeichiReviewGateway{}).
		Complete(r)
}
