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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.libre.sh/application/components"
	"k8s.libre.sh/application/settings"
)

// NextcloudStatus defines the observed state of Nextcloud
type NextcloudStatus struct {
	// Version defines the installed version
	Version  string                    `json:"version,omitempty"`
	Settings map[string]SettingsStatus `json:"settings,omitempty"`
	Phase    Phase                     `json:"phase,omitempty"`
}

type SettingsStatus struct {
	Sources []settings.Source `json:"sources,omitempty"`
}

// NextcloudSpec defines the desired state of Nextcloud
type NextcloudSpec struct {
	Version  string   `json:"version,omitempty"`
	Settings Settings `json:"settings,omitempty"`
	App      *App     `json:"app,omitempty"`
	Web      *Web     `json:"web,omitempty"`
	CLI      *CLI     `json:"cli,omitempty"`
	//	Cron        Component           `json:"cron,omitempty"`
}

type App struct {
	*components.InternalWorkload `json:",inline"`
}

type Web struct {
	*components.Workload `json:",inline"`
}

type CLI struct {
	*components.CLI `json:",inline"`
}

// Phase is the current status of a App as a whole.
type Phase string

const (
	PhaseNone       Phase = ""
	PhasePlanning   Phase = "Planning"
	PhaseRunning    Phase = "Running"
	PhaseCreating   Phase = "Creating"
	PhaseInstalling Phase = "Installing"
	PhaseUpgrading  Phase = "Upgrading"
	PhaseComplete   Phase = "Complete"
	PhaseFailed     Phase = "Failed"
)

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
