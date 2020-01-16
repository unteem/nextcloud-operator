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

	"github.com/presslabs/controller-util/syncer"

	"k8s.io/apimachinery/pkg/util/intstr"

	networking "k8s.io/api/networking/v1beta1"

	interfaces "git.indie.host/nextcloud-operator/interfaces"
)

func (component *Component) NewIngressSyncer(r interfaces.Reconcile) syncer.Interface {
	return syncer.NewObjectSyncer("Ingress", component.Owner, &component.Ingress, r.GetClient(), r.GetScheme(), component.MutateIngress)
}

func (component *Component) MutateIngress() error {
	component.Runtime.MutateIngress(&component.Ingress)

	labels := component.Labels("web")
	component.Ingress.SetLabels(labels)

	bk := networking.IngressBackend{
		ServiceName: "test",
		ServicePort: intstr.FromString("http"),
	}

	bkpaths := []networking.HTTPIngressPath{
		{
			Path:    "/",
			Backend: bk,
		},
	}

	rules := []networking.IngressRule{}

	for _, d := range component.Runtime.Hosts {
		rules = append(rules, networking.IngressRule{
			Host: string(d),
			IngressRuleValue: networking.IngressRuleValue{
				HTTP: &networking.HTTPIngressRuleValue{
					Paths: bkpaths,
				},
			},
		})
	}

	fmt.Println(component.Ingress)
	component.Ingress.Spec.Rules = rules

	return nil
}
