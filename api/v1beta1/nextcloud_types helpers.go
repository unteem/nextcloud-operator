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
	"sort"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1beta1"

	interfaces "git.indie.host/nextcloud-operator/interfaces"
)

const (
	// InternalHTTPPort represents the internal port used by the runtime container
	HTTPPort = 8080
)

var (
	wwwDataUserID int64 = 33
)

func (r *Runtime) MutateContainer(obj *corev1.Container) error {
	r.SetDefaults()
	obj.Name = "test"
	obj.Image = r.Image
	obj.ImagePullPolicy = r.ImagePullPolicy
	obj.Ports = r.Ports
	obj.LivenessProbe = r.LivenessProbe
	obj.ReadinessProbe = r.ReadinessProbe
	return nil
}
func (s *Settings) MutateDeployment(obj *appsv1.Deployment) error {
	s.MutatePod(&obj.Spec.Template)
	return nil
}

func (r *Runtime) MutateDeployment(obj *appsv1.Deployment) error {
	obj.Spec.Replicas = r.Replicas
	r.MutatePod(&obj.Spec.Template)
	return nil
}

func (s *Settings) MutatePod(obj *corev1.PodTemplateSpec) error {
	container := &corev1.Container{}
	containers := []corev1.Container{}
	container.Name = "test"

	s.MutateContainerEnvFrom(container)

	containers = append(containers, *container)

	if len(obj.Spec.Containers) == 0 {
		obj.Spec.Containers = containers
	} else {
		s.MutateContainerEnvFrom(&obj.Spec.Containers[0])
	}
	return nil
}

func (r *Runtime) MutatePod(obj *corev1.PodTemplateSpec) error {
	container := &corev1.Container{}
	containers := []corev1.Container{}

	obj.Spec.SecurityContext = r.SecurityContext

	r.MutateContainer(container)
	containers = append(containers, *container)

	if len(obj.Spec.Containers) == 0 {

		obj.Spec.Containers = containers
	} else {
		r.MutateContainer(&obj.Spec.Containers[0])
	}

	return nil
}

func (f *From) GetLocalObjectReference() corev1.LocalObjectReference {
	return f.LocalObjectReference
}

func (f *From) GetValue() string {
	return f.Value
}

func (f *From) GetKey() string {
	return f.Key
}

func GenEnv(e interfaces.EnvSource, object string) (corev1.EnvFromSource, corev1.EnvVar) {

	envFrom := corev1.EnvFromSource{}
	envVar := corev1.EnvVar{}

	if len(e.GetKey()) == 0 && len(e.GetLocalObjectReference().Name) > 0 && len(e.GetValue()) == 0 {
		if object == "configmap" {
			ref := &corev1.ConfigMapEnvSource{
				LocalObjectReference: e.GetLocalObjectReference(),
			}
			envFrom.ConfigMapRef = ref
		} else {
			ref := &corev1.SecretEnvSource{
				LocalObjectReference: e.GetLocalObjectReference(),
			}
			envFrom.SecretRef = ref
		}
	} else if len(e.GetLocalObjectReference().Name) > 0 && len(e.GetValue()) > 0 {
		envVar.Name = e.GetValue()
		if object == "configmap" {
			valueFrom := &corev1.EnvVarSource{
				ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
					LocalObjectReference: e.GetLocalObjectReference(),
					Key:                  e.GetKey(),
				},
			}
			envVar.ValueFrom = valueFrom
		} else {
			valueFrom := &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: e.GetLocalObjectReference(),
					Key:                  e.GetKey(),
				},
			}
			envVar.ValueFrom = valueFrom
		}
	}
	return envFrom, envVar
}

func (s *Settings) MutateContainerEnvFrom(obj *corev1.Container) error {

	configMapSources := s.Parameters.From
	secretSources := s.Secrets

	envVars := []corev1.EnvVar{}
	envFroms := []corev1.EnvFromSource{}

	for _, source := range configMapSources {
		envFrom, envVar := GenEnv(&source, "configmap")
		if len(envVar.Name) > 0 {
			envVars = append(envVars, envVar)
		}
		if envFrom.ConfigMapRef != nil {
			envFroms = append(envFroms, envFrom)
		}
	}
	if &secretSources != nil {
		for _, source := range secretSources {
			if &source != nil {
				envFrom, envVar := GenEnv(&source, "secret")
				if len(envVar.Name) > 0 {
					envVars = append(envVars, envVar)
				}
				if envFrom.SecretRef != nil {
					envFroms = append(envFroms, envFrom)
				}
			}
		}
	}

	for k, v := range s.Parameters.EnvVar {
		envVar := corev1.EnvVar{
			Name:  k,
			Value: v,
		}
		envVars = append(envVars, envVar)
	}

	// Sort var to avoid update of the object if var are not in the same order?
	sort.SliceStable(envVars, func(i, j int) bool {
		return envVars[i].Name < envVars[j].Name
	})

	obj.EnvFrom = envFroms
	obj.Env = envVars

	return nil
}

func (r *Runtime) MutateService(obj *corev1.Service) error {
	ports := []corev1.ServicePort{
		{
			Port: 9000,
			Name: "http",
		},
	}
	obj.Spec.Type = r.ServiceType
	obj.Spec.Ports = ports
	return nil
}

func (r *Runtime) MutateIngress(obj *networking.Ingress) error {
	obj.ObjectMeta.Annotations = r.IngressAnnotations

	if len(r.TLSSecretRef) > 0 {
		tls := networking.IngressTLS{
			SecretName: string(r.TLSSecretRef),
		}
		for _, d := range r.Hosts {
			tls.Hosts = append(tls.Hosts, string(d))
		}
		obj.Spec.TLS = []networking.IngressTLS{tls}
	} else {
		obj.Spec.TLS = nil
	}

	return nil
}

func (r *Runtime) SetDefaults() {
	if len(r.Image) == 0 {
		r.Image = "indiehosters/nextcloud"
	}
	if len(r.ServiceType) == 0 {
		r.ServiceType = "ClusterIP"
	}
	if r.Ports == nil {
		ports := []corev1.ContainerPort{
			{
				Name:          "http",
				ContainerPort: int32(HTTPPort),
				Protocol:      "TCP",
			},
		}
		r.Ports = ports
	}
	//	if r.Strategy == nil {
	//		r.Strategy = "recreate"
	//	}
	if r.SecurityContext == nil {
		r.SecurityContext = &corev1.PodSecurityContext{
			RunAsUser:  &wwwDataUserID,
			RunAsGroup: &wwwDataUserID,
			FSGroup:    &wwwDataUserID,
		}
	}
}
