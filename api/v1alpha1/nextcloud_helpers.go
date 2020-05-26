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
	"k8s.libre.sh/interfaces"
	"k8s.libre.sh/meta"
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

func (o *Nextcloud) GetSettings() settings.Settings {
	s := settings.NewSettings(&o.Spec.Settings)
	//	s.Generate = o.Spec.Settings.CreateOptions.Generate
	return s
}

func (app *Nextcloud) GetComponents() map[int]application.Component {
	components := map[int]application.Component{
		//	0: app.Spec.App,
	}

	return components
}

/* func (c *App) SetDefaults() {

	if c.Workload.Backend == nil {
		c.Workload.Backend = &components.Backend{}
	}
	if &c.Workload.Backend.Port == nil || c.Workload.Backend.Port.Port == 0 {
		c.Workload.Backend.Port.Port = 9000
	}
	if len(c.Workload.Backend.Port.Protocol) == 0 {
		c.Workload.Backend.Port.Protocol = "TCP"
	}
	if len(c.Workload.Backend.Port.Name) == 0 {
		c.Workload.Backend.Port.Name = "api"
	}

	if len(c.Workload.Backend.Paths) == 0 {
		c.Workload.Backend.Paths = []string{"/"}
	}

	if c.Workload.SecurityContext == nil {
		c.Workload.SecurityContext = &corev1.PodSecurityContext{
			RunAsUser:  &wwwDataUserID,
			RunAsGroup: &wwwDataUserID,
			FSGroup:    &wwwDataUserID,
		}
	}
	c.Workload.SetDefaults()

}
*/
func (app *Nextcloud) SetDefaults() {

	if app.Spec.Settings.CreateOptions.CommonMeta == nil {
		app.Spec.Settings.CreateOptions.CommonMeta = new(meta.ObjectMeta)
	}

	if app.Spec.Settings.CreateOptions.CommonMeta.Labels == nil {
		app.Spec.Settings.CreateOptions.CommonMeta.Labels = make(map[string]string)
	}

	app.Spec.Settings.CreateOptions.CommonMeta.SetComponent("settings")

	meta.SetObjectMeta(app, app.Spec.Settings.CreateOptions.CommonMeta)
	app.Spec.Settings.SetDefaults()
	//	app.Spec.App.SetDefaults()
}
