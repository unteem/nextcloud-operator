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

package cli

import (
	"fmt"

	"github.com/presslabs/controller-util/syncer"

	appsv1beta1 "git.indie.host/nextcloud-operator/api/v1beta1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"

	common "git.indie.host/nextcloud-operator/components/common"
	interfaces "git.indie.host/nextcloud-operator/interfaces"
)

type Component struct {
	Name string
	*common.Common
	Job batchv1.Job
}

func CreateAndInit(nc *appsv1beta1.Nextcloud) *Component {
	component := &Component{}
	component.Name = "cli"
	component.Common = common.NewCommon(nc)
	component.Owner = nc

	component.Job.SetName(component.GetName())
	component.Job.SetNamespace(component.Owner.Namespace)
	return component
}

func (component *Component) NewJobSyncer(r interfaces.Reconcile) syncer.Interface {
	return syncer.NewObjectSyncer("Job", component.Owner, &component.Job, r.GetClient(), r.GetScheme(), component.MutateJob)
}

func (component *Component) MutateJob() error {
	component.Settings.MutatePod(&component.Job.Spec.Template)
	component.Runtime.MutatePod(&component.Job.Spec.Template)

	labels := component.Labels("cli")
	component.Job.SetLabels(labels)

	component.Job.Spec.Template.ObjectMeta = component.Job.ObjectMeta
	component.Job.Spec.Template.Spec.RestartPolicy = corev1.RestartPolicyNever

	// args := []string{"/usr/local/bin/php", "/var/www/html/cron.php"}
	// component.Job.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Args = args

	return nil
}

func (component *Component) GetName() string {
	return fmt.Sprintf("%s-%s", component.Owner.Name, component.Name)
}
