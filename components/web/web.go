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
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1beta1"

	appsv1beta1 "git.indie.host/nextcloud-operator/api/v1beta1"
	common "git.indie.host/nextcloud-operator/components/common"

	"k8s.io/apimachinery/pkg/labels"
)

type Component struct {
	Name       string
	Owner      *appsv1beta1.Nextcloud
	Settings   *appsv1beta1.Settings
	Runtime    *appsv1beta1.Runtime
	Deployment appsv1.Deployment
	Service    corev1.Service
	Ingress    networking.Ingress
	ConfigMap  corev1.ConfigMap
}

func CreateAndInit(common *common.Common) *Component {
	component := &Component{}
	component.Name = "web"
	component.Owner = common.Owner

	component.Runtime = &component.Owner.Spec.Web.Runtime
	component.Settings = &component.Owner.Spec.Web.Settings

	component.Service.Name = component.GetName()
	component.Service.Namespace = component.Owner.Namespace

	component.Ingress.SetName(component.GetName())
	component.Ingress.SetNamespace(component.Owner.Namespace)

	component.Deployment.SetName(component.GetName())
	component.Deployment.SetNamespace(component.Owner.Namespace)

	component.ConfigMap.SetName(component.GetName())
	component.ConfigMap.SetNamespace(component.Owner.Namespace)

	return component
}

func (component *Component) GetName() string {
	return fmt.Sprintf("%s-%s", component.Owner.Name, component.Name)
}

func (component *Component) SetDefaults() {
	// if component.Settings
}

func (c *Component) Labels(component string) labels.Set {
	partOf := "nextcloud"
	//	if o.ObjectMeta.Labels != nil && len(o.ObjectMeta.Labels["app.kubernetes.io/part-of"]) > 0 {
	//		partOf = o.ObjectMeta.Labels["app.kubernetes.io/part-of"]
	//	}

	labels := labels.Set{
		"app.kubernetes.io/name":     "nextcloud",
		"app.kubernetes.io/part-of":  partOf,
		"app.kubernetes.io/instance": c.Owner.ObjectMeta.Name,
		//	"app.kubernetes.io/version":    c.Owner.Spec.AppVersion,
		"app.kubernetes.io/component":  component,
		"app.kubernetes.io/managed-by": "nextcloud-operator.libre.sh",
	}

	return labels
}
