package controller

import (
	"context"

	seichiclickv1alpha1 "github.com/GiganticMinecraft/seichi-gateway-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SeichiReviewGatewayReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *SeichiReviewGatewayReconciler) Reconcile(ctx context.Context, _ ctrl.Request) (ctrl.Result, error) {
	return ReconcileAllManagedResources(ctx, r.Client)
}

func (r *SeichiReviewGatewayReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&seichiclickv1alpha1.SeichiReviewGateway{}).
		Complete(r)
}
