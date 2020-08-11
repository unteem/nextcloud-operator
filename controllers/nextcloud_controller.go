/*

Licensed under the GNU AFFERO GENERAL PUBLIC LICENSE Version 3 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.gnu.org/licenses/agpl-3.0.html

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.libre.sh/controller-utils/application"
	"k8s.libre.sh/controller-utils/application/reconciler"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	appsv1alpha1 "k8s.libre.sh/apps/nextcloud/api/v1alpha1"

	oplib "github.com/redhat-cop/operator-utils/pkg/util"
)

// NextcloudReconciler reconciles a Nextcloud object
type NextcloudReconciler struct {
	Log logr.Logger
	reconciler.ReconcilerBase
	Manager manager.Manager
}

func ignoreNotFound(err error) error {
	if apierrs.IsNotFound(err) {
		return nil
	}
	return err
}

func (r *NextcloudReconciler) GetLogger() logr.Logger { return r.Log }

// +kubebuilder:rbac:groups=apps.libre.sh,resources=nextclouds,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps.libre.sh,resources=nextclouds/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=batch,resources=jobs/status,verbs=get
// +kubebuilder:rbac:groups=,resources=secrets;configmaps;services;events,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
func (r *NextcloudReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("nextcloud", req.NamespacedName)
	log.Info("reconciling")

	nc := &appsv1alpha1.Nextcloud{}
	if err := r.GetClient().Get(ctx, req.NamespacedName, nc); err != nil {
		log.Error(err, "unable to fetch Nextcloud")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, ignoreNotFound(err)
	}

	nc.Init()
	nc.SetDefaults()

	// app := application.NewApplication(nc)
	app := application.NewApplication(nc.GetSettings(), nc.GetComponents(), nc.GetJobs(), nc.GetOwner(), nc)

	instanceReconciler := reconciler.NewInstanceReconciler(&r.ReconcilerBase, app, nc.GetComponentsSyncOrder())
	res, err := instanceReconciler.Reconcile(req)

	return res, err

}

func (r *NextcloudReconciler) SetupWithManager(mgr ctrl.Manager) error {
	// dep := &appsv1.Deployment{}

	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1alpha1.Nextcloud{}).
		//	Owns(dep).
		WithEventFilter(oplib.ResourceGenerationOrFinalizerChangedPredicate{}).
		Complete(r)
}
