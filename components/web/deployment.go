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

package web

import (
	"github.com/presslabs/controller-util/syncer"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	interfaces "git.indie.host/nextcloud-operator/interfaces"
)

func (component *Component) NewDeploymentSyncer(r interfaces.Reconcile) syncer.Interface {
	return syncer.NewObjectSyncer("Deployment", component.Owner, &component.Deployment, r.GetClient(), r.GetScheme(), component.MutateDeployment)
}

func (component *Component) MutateDeployment() error {
	// component.Settings.MutateDeployment(&component.Deployment)
	component.Runtime.MutateDeployment(&component.Deployment)

	labels := component.Labels("web")

	component.Deployment.Spec.Template.ObjectMeta = component.Deployment.ObjectMeta
	component.Deployment.Spec.Selector = metav1.SetAsLabelSelector(labels)

	component.Deployment.Spec.Template.Spec.Containers[0].VolumeMounts = component.GenContainerVolumeMounts()
	component.Deployment.Spec.Template.Spec.Volumes = component.GenContainerVolumes()

	return nil
}

func (c *Component) GenContainerVolumeMounts() []corev1.VolumeMount {
	var mounts []corev1.VolumeMount

	mount := corev1.VolumeMount{
		Name:      "nginx",
		MountPath: "/etc/nginx/nginx.conf",
		SubPath:   "nginx.conf",
	}

	mounts = append(mounts, mount)

	return mounts
}

func (c *Component) GenContainerVolumes() []corev1.Volume {
	var volumes []corev1.Volume

	volume := corev1.Volume{
		Name: "nginx",
		VolumeSource: corev1.VolumeSource{
			ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: c.GetName(),
				},
			},
		},
	}
	volumes = append(volumes, volume)
	return volumes
}
