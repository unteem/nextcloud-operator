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

	appsv1beta1 "git.indie.host/nextcloud-operator/api/v1beta1"
	common "git.indie.host/nextcloud-operator/components/common"
)

type App struct {
	Name string
	*common.Common
	Deployment appsv1.Deployment
	Service    corev1.Service
	Ingress    networking.Ingress
	Secret     corev1.Secret
}

func NewApp(nc *appsv1beta1.Nextcloud) *App {
	app := &App{}
	app.Name = "app"
	app.Common = common.NewCommon(nc)
	app.Owner = nc

	app.Service.Name = app.GetName()
	app.Service.Namespace = app.Owner.Namespace

	app.Ingress.SetName(app.GetName())
	app.Ingress.SetNamespace(app.Owner.Namespace)

	app.Deployment.SetName(app.GetName())
	app.Deployment.SetNamespace(app.Owner.Namespace)

	app.Secret.SetName(app.GetName())
	app.Secret.SetNamespace(app.Owner.Namespace)

	return app
}

func (app *App) GetName() string {
	return fmt.Sprintf("%s-%s", app.Owner.Name, app.Name)
}
