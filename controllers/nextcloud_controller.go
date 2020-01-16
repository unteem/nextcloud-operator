/*

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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/presslabs/controller-util/syncer"

	"k8s.io/apimachinery/pkg/labels"

	appsv1beta1 "git.indie.host/nextcloud-operator/api/v1beta1"
)

// NextcloudReconciler reconciles a Nextcloud object
type NextcloudReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

func ignoreNotFound(err error) error {
	if apierrs.IsNotFound(err) {
		return nil
	}
	return err
}

// +kubebuilder:rbac:groups=apps.libre.sh,resources=nextclouds,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps.libre.sh,resources=nextclouds/status,verbs=get;update;patch

func (r *NextcloudReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("nextcloud", req.NamespacedName)
	log.Info("reconciling")

	app := &appsv1beta1.Nextcloud{}
	if err := r.Get(ctx, req.NamespacedName, app); err != nil {
		log.Error(err, "unable to fetch Nextcloud")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, ignoreNotFound(err)
	}
	// fmt.Println(app)

	componentApp := NewApp(app)
	//	componentApp.MutateDeployment()
	componentCron := NewCron(app)

	objectSyncer := componentApp.NewDeploymentSyncer(r)
	serviceSyncer := componentApp.NewServiceSyncer(r)
	ingressSyncer := componentApp.NewIngressSyncer(r)
	cronSyncer := componentCron.NewCronJobSyncer(r)

	if err := syncer.Sync(context.TODO(), objectSyncer, r.Recorder); err != nil {
		return ctrl.Result{}, err
	}

	if err := syncer.Sync(context.TODO(), serviceSyncer, r.Recorder); err != nil {
		return ctrl.Result{}, err
	}

	if err := syncer.Sync(context.TODO(), ingressSyncer, r.Recorder); err != nil {
		return ctrl.Result{}, err
	}

	if err := syncer.Sync(context.TODO(), cronSyncer, r.Recorder); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *NextcloudReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1beta1.Nextcloud{}).
		Complete(r)
}

func (c *Common) Labels(component string) labels.Set {
	partOf := "nextcloud"
	//	if o.ObjectMeta.Labels != nil && len(o.ObjectMeta.Labels["app.kubernetes.io/part-of"]) > 0 {
	//		partOf = o.ObjectMeta.Labels["app.kubernetes.io/part-of"]
	//	}

	labels := labels.Set{
		"app.kubernetes.io/name":     "nextcloud",
		"app.kubernetes.io/part-of":  partOf,
		"app.kubernetes.io/instance": c.Owner.ObjectMeta.Name,
		//	"app.kubernetes.io/version":    c.Owner.Spec.AppVersion,
		"app.kubernetes.io/component":  component,
		"app.kubernetes.io/managed-by": "nextcloud-operator.libre.sh",
	}

	return labels
}

type Component interface {
	GetObjects() runtime.Object
}

//type Cron struct {
//	Deployment
//	Service
//	Secret
//	Confimap
//}

//type Cli struct {
//	Deployment
//	Service
//	Secret
//	Confimap
//}

//type Web struct {
//	Deployment
//	Service
//	Secret
//	Confimap
//}

//type Database struct {
//	Container
//	Secret
//	Confimap
//}

//type Storage struct {
//	Container
//	Secret
//	Confimap
//}
