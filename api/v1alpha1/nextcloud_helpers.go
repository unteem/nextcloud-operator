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

package v1alpha1

import (
	"k8s.libre.sh/application"
	"k8s.libre.sh/application/settings"
	"k8s.libre.sh/application/settings/parameters"
	"k8s.libre.sh/meta"
	"k8s.libre.sh/status"
)

var (
	wwwDataUserID int64 = 82
)

func (o *Nextcloud) GetOwner() status.ObjectWithStatus { return o }
func (o *Nextcloud) GetName() string                   { return o.Name }
func (o *Nextcloud) GetNamespace() string              { return o.Namespace }
func (o *Nextcloud) GetInstance() string               { return o.Name }
func (o *Nextcloud) SetInstance(s string)              {}
func (o *Nextcloud) GetVersion() string                { return o.Spec.Version }
func (o *Nextcloud) SetVersion(s string)               {}
func (o *Nextcloud) GetComponent() string              { return "instance" }
func (o *Nextcloud) SetComponent(s string)             {}
func (o *Nextcloud) GetPartOf() string                 { return "Nextcloud" }
func (o *Nextcloud) SetPartOf(s string)                {}
func (o *Nextcloud) GetManagedBy() string              { return "Nextcloud-operator" }
func (o *Nextcloud) SetManagedBy(s string)             {}
func (o *Nextcloud) GetApplication() string            { return "Nextcloud" }
func (o *Nextcloud) SetApplication(s string)           {}

func (o *Nextcloud) GetApplicationStatus() status.ApplicationStatus {
	return o.Status.ApplicationStatus
}

func (o *Nextcloud) SetApplicationStatus(appStatus status.ApplicationStatus) {
	o.Status.ApplicationStatus = appStatus
}

func (o *Nextcloud) GetSettings() map[string]settings.Component {

	setts := map[string]settings.Component{
		"app": o.Spec.Settings.AppSettings,
		"web": o.Spec.Settings.Web,
	}

	return setts
}

func (app *Nextcloud) GetComponentsSyncOrder() map[int]string {
	return map[int]string{
		0: "settings",
		1: "install",
		2: "upgrade",
		3: "app",
		4: "web",
		5: "cron",
	}
}
func (app *Nextcloud) GetComponents() map[string]application.Component {

	cpts := map[string]application.Component{
		"app":  app.Spec.App,
		"web":  app.Spec.Web,
		"cron": app.Spec.Cron,
	}

	return cpts
}

func (app *Nextcloud) GetJobs() map[string]application.Component {

	cpts := map[string]application.Component{
		"install": app.Spec.Jobs.Install,
		"upgrade": app.Spec.Jobs.Upgrade,
	}

	return cpts
}

func (app *Nextcloud) Init() {

	if app.Spec.App == nil {
		app.Spec.App = &App{}
	}

	if app.Spec.Web == nil {
		app.Spec.Web = &Web{}
	}

	if app.Spec.Cron == nil {
		app.Spec.Cron = &CronJob{}
	}

	if app.Spec.Jobs == nil {
		app.Spec.Jobs = &Jobs{}
	}

	app.Spec.Jobs.Init()

	for _, c := range app.GetComponents() {
		c.Init()
	}

	if app.Spec.Settings == nil {
		app.Spec.Settings = &Settings{}
	}

	app.Spec.Settings.Init()

}

func (app *Nextcloud) SetDefaults() {
	meta.SetObjectMetaFromInstance(app, app.Spec.App)
	app.Spec.App.SetDefaults()

	meta.SetObjectMetaFromInstance(app, app.Spec.Web)
	app.Spec.Web.SetDefaults()

	meta.SetObjectMetaFromInstance(app, app.Spec.Jobs)
	app.Spec.Jobs.SetDefaults()

	meta.SetObjectMetaFromInstance(app, app.Spec.Cron)
	app.Spec.Cron.SetDefaults()

	meta.SetObjectMetaFromInstance(app, app.Spec.Settings.CreateOptions.CommonMeta)
	app.Spec.Settings.SetDefaults()

	// TODO TOFIX
	if app.Spec.Settings.AppSettings.General.Version == nil {
		app.Spec.Settings.AppSettings.General.Version = &parameters.Parameter{}
	}

	if len(app.Spec.Settings.AppSettings.General.Version.Value) == 0 || len(app.Spec.Settings.AppSettings.General.Version.Ref) == 0 {
		app.Spec.Settings.AppSettings.General.Version.Value = app.Spec.Version
	}
}
