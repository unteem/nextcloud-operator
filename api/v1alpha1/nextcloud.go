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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.libre.sh/controller-utils/application/components"
	"k8s.libre.sh/controller-utils/status"
)

// NextcloudStatus defines the observed state of Nextcloud
type NextcloudStatus struct {
	status.ApplicationStatus `json:",inline"`
}

// NextcloudSpec defines the desired state of Nextcloud
type NextcloudSpec struct {
	Version  string    `json:"version,omitempty"`
	Settings *Settings `json:"settings,omitempty"`
	App      *App      `json:"app,omitempty"`
	Web      *Web      `json:"web,omitempty"`
	Jobs     *Jobs     `json:"jobs,omitempty"`
	Cron     *CronJob  `json:"cron,omitempty"`
}

type App struct {
	*components.InternalWorkload `json:",inline"`
}

type Web struct {
	*components.Workload `json:",inline"`
}

type Jobs struct {
	*components.Jobs `json:",inline"`
}

type CronJob struct {
	*components.CronJob `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// Nextcloud is the Schema for the nextclouds API
type Nextcloud struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NextcloudSpec   `json:"spec,omitempty"`
	Status NextcloudStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NextcloudList contains a list of Nextcloud
type NextcloudList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Nextcloud `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Nextcloud{}, &NextcloudList{})
}
