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
	"fmt"

	"github.com/go-logr/logr"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/presslabs/controller-util/syncer"

	appsv1beta1 "git.indie.host/nextcloud-operator/api/v1beta1"
	application "git.indie.host/nextcloud-operator/components/app"
	"git.indie.host/nextcloud-operator/components/common"
	cron "git.indie.host/nextcloud-operator/components/cron"
	"git.indie.host/nextcloud-operator/components/web"
	"git.indie.host/nextcloud-operator/util"
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

	app := &appsv1beta1.Nextcloud{}
	if err := r.Get(ctx, req.NamespacedName, app); err != nil {
		log.Error(err, "unable to fetch Nextcloud")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, ignoreNotFound(err)
	}

	// var phase appsv1beta1.Phase

	phase, err := util.GetAppPhase(app.Status, app.Spec.Version)
	if err != nil {
		return ctrl.Result{}, err
	}

	fmt.Println(phase)

	common := common.CreateAndInit(app)

	componentApp := application.CreateAndInit(common)
	componentCron := cron.CreateAndInit(common)
	componentWeb := web.CreateAndInit(common)
	// 	componentCLI := cli.CreateAndInit(common)

	appSyncers := []syncer.Interface{
		componentApp.NewSecretSyncer(r),
		componentApp.NewDeploymentSyncer(r),
		componentApp.NewServiceSyncer(r),
	}

	cronSyncers := []syncer.Interface{
		componentCron.NewCronJobSyncer(r),
	}

	jobSyncers := []syncer.Interface{
		// jobSyncer := componentCLI.NewJobSyncer(r)
	}

	webSyncers := []syncer.Interface{
		componentWeb.NewConfigMapSyncer(r),
		componentWeb.NewDeploymentSyncer(r),
		componentWeb.NewIngressSyncer(r),
		componentWeb.NewServiceSyncer(r),
	}

	err = r.sync(appSyncers)
	err = r.sync(cronSyncers)
	err = r.sync(jobSyncers)
	err = r.sync(webSyncers)
	if err != nil {
		return ctrl.Result{}, err
	}

	app.Status.Version = app.Spec.Version
	app.Status.Phase = appsv1beta1.PhaseRunning

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
		For(&appsv1beta1.Nextcloud{}).
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
