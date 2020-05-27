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
	"k8s.libre.sh/application"
	"k8s.libre.sh/objects"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/presslabs/controller-util/syncer"

	appsv1alpha1 "git.indie.host/operators/nextcloud-operator/api/v1alpha1"
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

func (r *NextcloudReconciler) GetClient() client.Client          { return r.Client }
func (r *NextcloudReconciler) GetScheme() *runtime.Scheme        { return r.Scheme }
func (r *NextcloudReconciler) GetRecorder() record.EventRecorder { return r.Recorder }
func (r *NextcloudReconciler) GetLogger() logr.Logger            { return r.Log }

// +kubebuilder:rbac:groups=apps.libre.sh,resources=nextclouds,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps.libre.sh,resources=nextclouds/status,verbs=get;update;patch

func (r *NextcloudReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("nextcloud", req.NamespacedName)
	log.Info("reconciling")

	app := &appsv1alpha1.Nextcloud{}
	if err := r.Get(ctx, req.NamespacedName, app); err != nil {
		log.Error(err, "unable to fetch Nextcloud")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, ignoreNotFound(err)
	}

	// var phase appsv1alpha1.Phase

	/* 	phase, err := util.GetAppPhase(app.Status, app.Spec.Version)
	   	if err != nil {
	   		return ctrl.Result{}, err
	   	}

	   	fmt.Println(phase) */

	application.Init(app, r)

	sett, err := application.CreateAndInitSettings(app, r)
	if err != nil {
		return ctrl.Result{}, err
	}

	for _, obj := range sett.GetObjects() {
		s := objects.NewObjectSyncer(obj, app, r)

		if err := syncer.Sync(context.TODO(), s, r.GetRecorder()); err != nil {
			return ctrl.Result{}, err
		}

	}

	/* 	syncers, err := application.NewSyncers(app, r, app)
	   	if err != nil {
	   		return ctrl.Result{}, err
	   	}

	   	err = r.sync(syncers)
	   	if err != nil {
	   		return ctrl.Result{}, err
	   	} */

	//	for _, obj := range sett.GetObjects() {
	//	meta.SetObjectMeta(sett.CommonMeta, obj)
	//	s := objects.NewObjectSyncer(obj, owner, r)

	//	if err := syncer.Sync(context.TODO(), s, r.GetRecorder()); err != nil {
	//		return syncers, err
	//	}

	//		fmt.Println(obj)

	//	}

	// fmt.Println(sett)

	/*
		if phase == appsv1alpha1.PhaseInstalling || phase == appsv1alpha1.PhaseRunning {
			jobSyncers := []syncer.Interface{
				componentCLI.NewJobSyncer(r),
			}

			err = r.sync(jobSyncers)
		}
	*/

	app.Status.Version = app.Spec.Version
	app.Status.Phase = appsv1alpha1.PhaseRunning

	oldStatus := app.Status.DeepCopy()
	if oldStatus != &app.Status {
		if err := r.Status().Update(ctx, app); err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *NextcloudReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1alpha1.Nextcloud{}).
		Complete(r)
}

func (r *NextcloudReconciler) sync(syncers []syncer.Interface) error {
	for _, s := range syncers {
		if err := syncer.Sync(context.TODO(), s, r.Recorder); err != nil {
			return err
		}
	}
	return nil
}
