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

	"github.com/presslabs/controller-util/syncer"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	interfaces "git.indie.host/nextcloud-operator/interfaces"
)

func (app *App) NewDeploymentSyncer(r interfaces.Reconcile) syncer.Interface {
	return syncer.NewObjectSyncer("Deployment", app.Owner, &app.Deployment, r.GetClient(), r.GetScheme(), app.MutateDeployment)
}

func (app *App) MutateDeployment() error {
	app.Settings.MutateDeployment(&app.Deployment)
	app.Runtime.MutateDeployment(&app.Deployment)

	labels := app.Labels("app")

	app.Deployment.SetLabels(labels)

	app.Deployment.Spec.Template.ObjectMeta = app.Deployment.ObjectMeta
	app.Deployment.Spec.Selector = metav1.SetAsLabelSelector(labels)

	fmt.Println(app.Deployment.Spec.Template.Spec.Containers[0])

	return nil
}
