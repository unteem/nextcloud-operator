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
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1beta1"
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
		fmt.Println("container in settings")
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
		fmt.Println("no container in runtime")

		obj.Spec.Containers = containers
	} else {
		fmt.Println("container in settings")
		r.MutateContainer(&obj.Spec.Containers[0])
	}

	return nil
}

func (s *Settings) MutateContainerEnvFrom(obj *corev1.Container) error {
	configMapSources := s.Parameters.From
	secretSources := s.Secrets

	for _, source := range configMapSources {
		if len(source.Key) == 0 {
			envFrom := corev1.EnvFromSource{}
			envFrom.ConfigMapRef.LocalObjectReference = source.LocalObjectReference
			obj.EnvFrom = append(obj.EnvFrom, envFrom)
		}
		if len(source.Key) > 0 {
			envVar := corev1.EnvVar{}
			valueFrom := &corev1.EnvVarSource{
				ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: source.Name,
					},
					Key: source.Key,
				},
			}
			envVar.Name = source.Value
			envVar.ValueFrom = valueFrom
			obj.Env = append(obj.Env, envVar)
		}
	}

	for _, source := range secretSources {
		if len(source.Key) == 0 {
			envFrom := corev1.EnvFromSource{
				SecretRef: &corev1.SecretEnvSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: source.Name,
					},
				},
			}
			obj.EnvFrom = append(obj.EnvFrom, envFrom)
		}
		if len(source.Key) > 0 {
			envVar := corev1.EnvVar{}
			valueFrom := &corev1.EnvVarSource{
				ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: source.Name,
					},
					Key: source.Key,
				},
			}
			envVar.Name = source.Value
			envVar.ValueFrom = valueFrom
			obj.Env = append(obj.Env, envVar)
		}
	}

	for k, v := range s.Parameters.EnvVar {
		envVar := corev1.EnvVar{
			Name:  k,
			Value: v,
		}
		obj.Env = append(obj.Env, envVar)
	}

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
