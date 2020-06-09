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
	"k8s.libre.sh/application/components"
	"k8s.libre.sh/application/settings"
	"k8s.libre.sh/interfaces"
	"k8s.libre.sh/meta"
	"k8s.libre.sh/objects/job"
)

var (
	wwwDataUserID int64 = 33
)

func (o *Nextcloud) GetOwner() interfaces.Object { return o }
func (o *Nextcloud) GetName() string             { return o.Name }
func (o *Nextcloud) GetNamespace() string        { return o.Namespace }
func (o *Nextcloud) GetInstance() string         { return o.Name }
func (o *Nextcloud) SetInstance(s string)        {}
func (o *Nextcloud) GetVersion() string          { return o.Spec.Version }
func (o *Nextcloud) SetVersion(s string)         {}
func (o *Nextcloud) GetComponent() string        { return "instance" }
func (o *Nextcloud) SetComponent(s string)       {}
func (o *Nextcloud) GetPartOf() string           { return "Nextcloud" }
func (o *Nextcloud) SetPartOf(s string)          {}
func (o *Nextcloud) GetManagedBy() string        { return "Nextcloud-operator" }
func (o *Nextcloud) SetManagedBy(s string)       {}
func (o *Nextcloud) GetApplication() string      { return "Nextcloud" }
func (o *Nextcloud) SetApplication(s string)     {}

func (o *Nextcloud) GetSettings() map[string]settings.Settings {

	// TODO TO FIX
	setts := map[string]settings.Settings{
		"app": settings.NewSettings(&o.Spec.Settings.AppSettings),
		"web": settings.NewSettings(&o.Spec.Settings.Web),
	}

	if len(o.Status.Settings) > 0 {
		for k, v := range o.Status.Settings {
			setts[k].(*settings.Component).ConfigSpec.Sources = append(setts[k].(*settings.Component).ConfigSpec.Sources, v.Sources...)
		}
	}

	return setts
}

func (app *Nextcloud) GetComponentsSyncOrder() map[int]string {
	return map[int]string{
		0: "settings",
		1: "cli",
		2: "app",
		3: "web",
	}
}
func (app *Nextcloud) GetComponents() map[string]application.Component {

	cpts := map[string]application.Component{
		"app": app.Spec.App,
		"web": app.Spec.Web,
		"cli": app.Spec.CLI,
	}

	return cpts
}

func (app *Nextcloud) Init() {

	if app.Spec.App == nil {
		app.Spec.App = &App{}
		app.Spec.App.InternalWorkload = &components.InternalWorkload{}
	}

	if app.Spec.Web == nil {
		app.Spec.Web = &Web{}
		app.Spec.Web.Workload = &components.Workload{}

	}

	if app.Spec.CLI == nil {
		app.Spec.CLI = &CLI{}
		app.Spec.CLI.CLI = &components.CLI{}

	}

	for _, c := range app.GetComponents() {
		c.Init()
	}

	if app.Spec.CLI == nil {
		app.Spec.CLI = &CLI{}
		app.Spec.CLI.Job = &job.Job{}
	}

	app.Spec.CLI.Init()

	app.Spec.Settings.AppSettings.CreateOptions.Init()
	app.Spec.Settings.Web.CreateOptions.Init()

}

func (app *Nextcloud) SetDefaultMeta() {

	app.Spec.Web.ObjectMeta.SetComponent("web")

	// TODO TO FIX create a func in application package
	for _, c := range app.GetComponents() {

		meta.SetObjectMetaFromInstance(app, c)

		for _, o := range c.GetObjects() {
			meta.SetObjectMeta(c, o)
		}
	}

	// TODO TOFIX
	meta.SetObjectMeta(app, app.Spec.CLI.ObjectMeta)
	app.Spec.CLI.ObjectMeta.SetComponent("cli")

	meta.SetObjectMeta(app, app.Spec.Settings.AppSettings.CreateOptions.CommonMeta)
	app.Spec.Settings.AppSettings.CreateOptions.CommonMeta.Labels["app.kubernetes.io/component"] = "app"

	meta.SetObjectMeta(app, app.Spec.Settings.Web.CreateOptions.CommonMeta)
	app.Spec.Settings.Web.CreateOptions.CommonMeta.Labels["app.kubernetes.io/component"] = "web"

}

func (app *Nextcloud) SetDefaults() {

	app.Spec.App.SetDefaults()
	app.Spec.Web.SetDefaults()
	app.Spec.CLI.SetDefaults()

	app.Spec.Settings.SetDefaults()
	app.Spec.Settings.SetDefaults()

}
