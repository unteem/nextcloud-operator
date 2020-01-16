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

package application

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1beta1"

	common "git.indie.host/nextcloud-operator/components/common"
	interfaces "git.indie.host/nextcloud-operator/interfaces"
)

type Component struct {
	Name string
	*common.Common
	Deployment appsv1.Deployment
	Service    corev1.Service
	Ingress    networking.Ingress
	Secret     corev1.Secret
}

func CreateAndInit(common *common.Common) *Component {
	c := &Component{}
	c.Name = "app"
	c.Common = common

	objects := c.GetObjects()
	labels := c.Labels("app")

	for _, o := range objects {
		o.SetName(c.GetName())
		o.SetNamespace(c.Owner.Namespace)
		o.SetLabels(labels)
	}

	return c
}

func (c *Component) GetName() string {
	return fmt.Sprintf("%s-%s", c.Owner.Name, c.Name)
}

func (c *Component) GetObjects() []interfaces.Object {
	return []interfaces.Object{
		&c.Secret,
		&c.Deployment,
		&c.Service,
		&c.Ingress,
	}
}
