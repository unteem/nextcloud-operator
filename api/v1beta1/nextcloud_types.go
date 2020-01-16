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

package v1beta1

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NextcloudStatus defines the observed state of Nextcloud
type NextcloudStatus struct {
	// Version defines the installed version
	Version string `json:"version,omitempty"`
	Phase   Phase  `json:"Phase,omitempty"`
}

// NextcloudSpec defines the desired state of Nextcloud
type NextcloudSpec struct {
	//	Replicas *int32 `json:"replicas,omitempty"`
	//	Hosts    []Host `json:"hosts,omitempty"`
	//	Storage  storage.StorageSpec `json:"storage,omitempty"`
	Version     string     `json:"version,omitempty"`
	App         Component  `json:"app,omitempty"`
	Web         Component  `json:"web,omitempty"`
	CLI         Component  `json:"cli,omitempty"`
	Cron        Component  `json:"cron,omitempty"`
	Database    Dependency `json:"database,omitempty"`
	SMTP        Dependency `json:"smtp,omitempty"`
	ObjectStore Dependency `json:"objectStore,omitempty"`
	Redis       Dependency `json:"redis,omitempty"`
}

type From struct {
	// The ConfigMap to select from.
	corev1.LocalObjectReference `json:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	Value                       string `json:"value,omitempty"`
	// The key to select.
	Key string `json:"key,omitempty" protobuf:"bytes,2,opt,name=key"`
	// An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
	// +optional
	Prefix string `json:"prefix,omitempty" protobuf:"bytes,1,opt,name=prefix"`
}

//type Foo struct {
//	Name string        `json:"name,omitempty"`
//	Keys []ParametersKey `json:"keys,omitempty"`
//}

//type ParametersKey struct {
//	Value string `json:"value,omitempty"`
//	Key string `json:"key" protobuf:"bytes,2,opt,name=key"`
// An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
// +optional
//	Prefix string `json:"prefix,omitempty" protobuf:"bytes,1,opt,name=prefix"`
//}

type Parameters struct {
	From   []From            `json:"from,omitempty"`
	EnvVar map[string]string `json:"envVar,omitempty"`
}

type Settings struct {
	Secrets    []From      `json:"secrets,omitempty"`
	Parameters *Parameters `json:"parameters,omitempty"`
}

type Component struct {
	Enabled  bool     `json:"enabled,omitempty"`
	Name     string   `json:"name,omitempty"`
	Runtime  Runtime  `json:"runtime,omitempty"`
	Settings Settings `json:"settings,omitempty"`
}

type Dependency struct {
	Enabled  bool     `json:"enabled,omitempty"`
	Name     string   `json:"name,omitempty"`
	Settings Settings `json:"settings,omitempty"`
}

// SecretRef represents a reference to a Secret
type SecretRef string

// Host represents a valid hostname
type Host string

type Runtime struct {
	Image              string                        `json:"image,omitempty"`
	ImagePullPolicy    corev1.PullPolicy             `json:"imagePullPolicy,omitempty"`
	ImagePullSecrets   []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
	ServiceType        corev1.ServiceType            `json:"serviceType,omitempty"`
	IngressAnnotations map[string]string             `json:"ingressAnnotations,omitempty"`
	Hosts              []Host                        `json:"hosts,omitempty"`
	// Number of desired pods. This is a pointer to distinguish between explicit
	// zero and not specified. Defaults to 1.
	// +optional
	Replicas        *int32                      `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas"`
	TLSSecretRef    SecretRef                   `json:"tlsSecretRef,omitempty"`
	Ports           []corev1.ContainerPort      `json:"ports,omitempty"`
	Resources       corev1.ResourceRequirements `json:"resources,omitempty" protobuf:"bytes,8,opt,name=resources"`
	SecurityContext *corev1.PodSecurityContext  `json:"securityContext,omitempty"`
	ReadinessProbe  *corev1.Probe               `json:"readinessProbe,omitempty" protobuf:"bytes,11,opt,name=readinessProbe"`
	LivenessProbe   *corev1.Probe               `json:"livenessProbe,omitempty" protobuf:"bytes,10,opt,name=livenessProbe"`
	// The deployment strategy to use to replace existing pods with new ones.
	// +optional
	// +patchStrategy=retainKeys
	Strategy appsv1.DeploymentStrategy `json:"strategy,omitempty" patchStrategy:"retainKeys" protobuf:"bytes,4,opt,name=strategy"`
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
