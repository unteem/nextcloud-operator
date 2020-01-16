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

	"k8s.io/apimachinery/pkg/util/intstr"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	appsv1beta1 "git.indie.host/nextcloud-operator/api/v1beta1"
	common "git.indie.host/nextcloud-operator/components/common"
	interfaces "git.indie.host/nextcloud-operator/interfaces"
)

type App struct {
	*common.Common
	Deployment appsv1.Deployment
	Service    corev1.Service
	Ingress    networking.Ingress
}

func NewApp(nc *appsv1beta1.Nextcloud) *App {
	app := &App{}
	app.Common = common.NewCommon(nc)
	app.Owner = nc
	app.Service.Name = "test"
	app.Service.Namespace = app.Owner.Namespace
	app.Ingress.SetName("test")
	app.Ingress.SetNamespace(app.Owner.Namespace)

	app.Deployment.SetName("test")
	app.Deployment.SetNamespace(app.Owner.Namespace)
	return app
}

func (app *App) NewDeploymentSyncer(r interfaces.Reconcile) syncer.Interface {
	return syncer.NewObjectSyncer("Deployment", app.Owner, &app.Deployment, r.GetClient(), r.GetScheme(), app.MutateDeployment)
}

func (app *App) NewServiceSyncer(r interfaces.Reconcile) syncer.Interface {
	return syncer.NewObjectSyncer("Service", app.Owner, &app.Service, r.GetClient(), r.GetScheme(), app.MutateService)
}

func (app *App) NewIngressSyncer(r interfaces.Reconcile) syncer.Interface {
	return syncer.NewObjectSyncer("Ingress", app.Owner, &app.Ingress, r.GetClient(), r.GetScheme(), app.MutateIngress)
}

func (app *App) MutateService() error {
	labels := app.Labels("app")

	app.Runtime.MutateService(&app.Service)
	app.Service.SetLabels(labels)
	app.Service.Spec.Selector = labels

	return nil
}

func (app *App) MutateDeployment() error {
	app.Settings.MutateDeployment(&app.Deployment)
	app.Runtime.MutateDeployment(&app.Deployment)

	labels := app.Labels("app")

	app.Deployment.SetLabels(labels)

	app.Deployment.Spec.Template.ObjectMeta = app.Deployment.ObjectMeta
	app.Deployment.Spec.Selector = metav1.SetAsLabelSelector(labels)

	return nil
}

func (app *App) MutateIngress() error {
	app.Runtime.MutateIngress(&app.Ingress)

	labels := app.Labels("web")
	app.Ingress.SetLabels(labels)

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

	for _, d := range app.Runtime.Hosts {
		rules = append(rules, networking.IngressRule{
			Host: string(d),
			IngressRuleValue: networking.IngressRuleValue{
				HTTP: &networking.HTTPIngressRuleValue{
					Paths: bkpaths,
				},
			},
		})
	}

	fmt.Println(app.Ingress)
	app.Ingress.Spec.Rules = rules

	return nil
}
