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

	interfaces "git.indie.host/nextcloud-operator/interfaces"
)

func (component *Component) NewServiceSyncer(r interfaces.Reconcile) syncer.Interface {
	return syncer.NewObjectSyncer("Service", component.Owner, &component.Service, r.GetClient(), r.GetScheme(), component.MutateService)
}

func (component *Component) MutateService() error {
	labels := component.Labels("web")

	component.Runtime.MutateService(&component.Service)
	component.Service.SetLabels(labels)
	component.Service.Spec.Selector = labels

	return nil
}
