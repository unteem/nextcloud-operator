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
	corev1 "k8s.io/api/core/v1"
	"k8s.libre.sh/controller-utils/application/components"
)

func (app *Jobs) Init() {
	/* 	if app.Install == nil {
	   		app.Install = &components.Job{}
	   	}
	   	if app.Upgrade == nil {
	   		app.Upgrade = &components.Job{}
	   	} */
	if app.Jobs == nil {
		app.Jobs = &components.Jobs{}
	}
	app.Jobs.Init()
}

func (app *Jobs) SetDefaults() {

	if app.Install.SecurityContext == nil {
		app.Install.SecurityContext = &corev1.PodSecurityContext{
			RunAsUser:  &wwwDataUserID,
			RunAsGroup: &wwwDataUserID,
			FSGroup:    &wwwDataUserID,
		}
	}

	if len(app.Install.Image) == 0 {
		app.Install.Image = "libresh/nextcloud:18.0.0"
	}

	if len(app.Install.RestartPolicy) == 0 {
		app.Install.RestartPolicy = corev1.RestartPolicyOnFailure
	}

	if len(app.Install.Settings) == 0 {
		app.Install.Settings = []string{"app"}
	}

	app.Install.ObjectMeta.SetComponent("job")

	if app.Upgrade.SecurityContext == nil {
		app.Upgrade.SecurityContext = &corev1.PodSecurityContext{
			RunAsUser:  &wwwDataUserID,
			RunAsGroup: &wwwDataUserID,
			FSGroup:    &wwwDataUserID,
		}
	}

	if len(app.Upgrade.Image) == 0 {
		app.Install.Image = "libresh/nextcloud:18.0.0"
	}

	if len(app.Upgrade.RestartPolicy) == 0 {
		app.Install.RestartPolicy = corev1.RestartPolicyOnFailure
	}

	if len(app.Upgrade.Settings) == 0 {
		app.Install.Settings = []string{"app"}
	}

	app.Upgrade.ObjectMeta.SetComponent("job")

}
