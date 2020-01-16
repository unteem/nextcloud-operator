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
	"github.com/presslabs/controller-util/syncer"

	interfaces "git.indie.host/nextcloud-operator/interfaces"
)

func (app *App) NewServiceSyncer(r interfaces.Reconcile) syncer.Interface {
	return syncer.NewObjectSyncer("Service", app.Owner, &app.Service, r.GetClient(), r.GetScheme(), app.MutateService)
}

func (app *App) MutateService() error {
	labels := app.Labels("app")

	app.Runtime.MutateService(&app.Service)
	app.Service.SetLabels(labels)
	app.Service.Spec.Selector = labels

	return nil
}
