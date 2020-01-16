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

package common

import (
	"github.com/imdario/mergo"
	"k8s.io/apimachinery/pkg/labels"

	appsv1beta1 "git.indie.host/nextcloud-operator/api/v1beta1"
)

type Common struct {
	Owner    *appsv1beta1.Nextcloud
	Settings *appsv1beta1.Settings
	Runtime  *appsv1beta1.Runtime
}

func CreateAndInit(app *appsv1beta1.Nextcloud) *Common {
	// app.SetDefaults()

	runtime := &app.Spec.App.Runtime
	settings := &app.Spec.App.Settings

	if app.Spec.Database.Enabled {
		settings.Secrets = append(settings.Secrets, app.Spec.Database.Settings.Secrets...)
		err := mergo.Merge(settings.Parameters, app.Spec.Database.Settings.Parameters)
		if err != nil {
			// do something
		}
	}
	if app.Spec.Redis.Enabled {
		settings.Secrets = append(settings.Secrets, app.Spec.Redis.Settings.Secrets...)
		err := mergo.Merge(settings.Parameters, app.Spec.Redis.Settings.Parameters)
		if err != nil {
			// do something
		}
	}
	if app.Spec.ObjectStore.Enabled {
		settings.Secrets = append(settings.Secrets, app.Spec.ObjectStore.Settings.Secrets...)
		err := mergo.Merge(settings.Parameters, app.Spec.ObjectStore.Settings.Parameters)
		if err != nil {
			// do something
		}
	}
	if app.Spec.SMTP.Enabled {
		err := mergo.Merge(settings, app.Spec.SMTP.Settings)
		if err != nil {
			// do something
		}
	}
	return &Common{
		Owner:    app,
		Settings: settings,
		Runtime:  runtime,
	}
}

func (c *Common) Labels(component string) labels.Set {
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

//func (c *Common) SetDefaults() {
//}

//func (c *Common) GetDeployment() obj *appsv1.deployment {
//	c.Settings.MutateDeployment(obj)
//	return obj
//}

//func (c *Common) GetService() obj *corev1.service {
//	return obj
//}

//func (c *Common) GetConfigMap() obj *corev1.ConfigMap {
//	return obj
//}

//func (c *Common) GetSecret() obj *corev1.Secret {
//return obj
//}
