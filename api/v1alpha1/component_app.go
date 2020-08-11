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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.libre.sh/controller-utils/application/settings/parameters"
)

func (app *App) SetDefaults() {
	if &app.Service.Port == nil || app.Service.Port.Port == 0 {
		app.Service.Port.Port = 9000
	}

	if len(app.Service.Port.Protocol) == 0 {
		app.Service.Port.Protocol = "TCP"
	}
	if len(app.Service.Port.Name) == 0 {
		app.Service.Port.Name = "api"
	}

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

	if len(app.Settings) == 0 {
		app.Settings = []string{"app"}
	}

	app.InternalWorkload.SetDefaults()

	installParam := &parameters.Parameter{
		Value:     "true",
		Key:       "INSTALLED",
		MountType: parameters.MountLiteral,
	}

	if app.Deployment.Parameters == nil {
		app.Deployment.Parameters = &parameters.Parameters{}
	}

	*app.Deployment.Parameters = append(*app.Deployment.Parameters, installParam)

}
