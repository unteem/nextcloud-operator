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
	"k8s.libre.sh/controller-utils/application/components"
)

func (app *CronJob) Init() {
	if app.CronJob == nil {
		app.CronJob = &components.CronJob{}
	}

	app.CronJob.Init()
}

func (app *CronJob) SetDefaults() {

	if len(app.Schedule) == 0 {
		app.Schedule = "*/5 * * * *"
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

	if len(app.RestartPolicy) == 0 {
		app.RestartPolicy = corev1.RestartPolicyOnFailure
	}

	if len(app.Settings) == 0 {
		app.Settings = []string{"app"}
	}

	app.CronJob.SetDefaults()

}
