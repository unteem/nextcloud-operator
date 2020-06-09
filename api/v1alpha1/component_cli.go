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
)

func (app *CLI) SetDefaults() {

	if app.SecurityContext == nil {
		app.SecurityContext = &corev1.PodSecurityContext{
			RunAsUser:  &wwwDataUserID,
			RunAsGroup: &wwwDataUserID,
			FSGroup:    &wwwDataUserID,
		}
	}

	if len(app.Image) == 0 {
		app.Image = "libresh/nextcloud:18.0.0"
	}

	if len(app.RestartPolicy) == 0 {
		app.RestartPolicy = corev1.RestartPolicyOnFailure
	}

	app.ObjectMeta.SetComponent("cli")

	// TODO TOFIX
	// meta.SetObjectMeta(app, app.ObjectMeta)

}
