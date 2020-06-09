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
	"github.com/presslabs/controller-util/syncer"
	appsv1 "k8s.io/api/apps/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.libre.sh/application"
	"k8s.libre.sh/objects"
	ctrl "sigs.k8s.io/controller-runtime"

	appsv1alpha1 "k8s.libre.sh/apps/nextcloud/api/v1alpha1"
	"k8s.libre.sh/apps/nextcloud/util"

	oplib "github.com/redhat-cop/operator-utils/pkg/util"
)

// NextcloudReconciler reconciles a Nextcloud object
type NextcloudReconciler struct {
	Log logr.Logger
	application.ReconcilerBase
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
func (r *NextcloudReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("nextcloud", req.NamespacedName)
	log.Info("reconciling")

	app := &appsv1alpha1.Nextcloud{}
	if err := r.GetClient().Get(ctx, req.NamespacedName, app); err != nil {
		log.Error(err, "unable to fetch Nextcloud")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, ignoreNotFound(err)
	}

	// Initiatilize & Defaults
	app.Init()
	app.SetDefaultMeta()
	app.SetDefaults()

	// Create and Init Settings
	sett, err := application.CreateAndInitSettings(app, r)

	if err != nil {
		return ctrl.Result{}, err
	}

	// Init Components
	cpts := app.GetComponents()
	application.InitComponentsFromSettings(cpts, sett)

	// Syncers
	syncers := make(map[string]map[int]syncer.Interface)
	settingsSyncers := make(map[int]syncer.Interface)
	syncers["settings"] = settingsSyncers

	// Settings Object Syncers
	for _, s := range sett {
		for k, obj := range s.GetObjects() {
			settingsSyncers[k] = objects.NewObjectSyncer(obj, app, r)
		}
	}

	// Migration, installer or upgrader component
	var phase appsv1alpha1.Phase

	phase, err = util.GetAppPhase(app.Status, app.Spec.Version)
	if err != nil {
		return ctrl.Result{}, err
	}

	switch phase {
	case appsv1alpha1.PhaseInstalling:
		cpts["cli"].(*appsv1alpha1.CLI).Job.Args = []string{"install"}
		fmt.Println(cpts["cli"].(*appsv1alpha1.CLI).Job.Args)
	case appsv1alpha1.PhaseUpgrading:
		cpts["cli"].(*appsv1alpha1.CLI).Job.Args = []string{"upgrade"}
	default:
		delete(cpts, "cli")
	}

	// Components Object Syncers
	for _, c := range cpts {
		syncers[c.GetComponent()] = application.NewObjectSyncersFromComponent(c, r, app)

	}

	cptsOrder := app.GetComponentsSyncOrder()
	err = application.Sync(ctx, r, syncers, cptsOrder)
	if err != nil {
		return ctrl.Result{}, err
	}

	settingsStatus := appsv1alpha1.SettingsStatus{}

	if app.Status.Settings == nil {
		app.Status.Settings = make(map[string]appsv1alpha1.SettingsStatus)
	}
	for _, s := range sett {
		// Update Status
		if len(s.GetConfig().GetSources()) > 0 {
			settingsStatus.Sources = s.GetConfig().GetSources()
			app.Status.Settings[s.GetMeta().GetComponent()] = settingsStatus
		}
	}

	// Reconcile status
	app.Status.Version = app.Spec.Version
	app.Status.Phase = appsv1alpha1.PhaseRunning

	oldStatus := app.Status.DeepCopy()
	if oldStatus != &app.Status {
		if err := r.GetClient().Status().Update(ctx, app); err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *NextcloudReconciler) SetupWithManager(mgr ctrl.Manager) error {
	dep := &appsv1.Deployment{}
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1alpha1.Nextcloud{}).
		Owns(dep).
		WithEventFilter(oplib.ResourceGenerationOrFinalizerChangedPredicate{}).
		Complete(r)
}
