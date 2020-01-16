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

func CreateAndInit(common *common.Common) *Component {
	c := &Component{}
	c.Name = "cli"
	c.Common = common

	labels := c.Labels("cli")

	objects := c.GetObjects()
	for _, o := range objects {
		o.SetName(c.GetName())
		o.SetNamespace(c.Owner.Namespace)
		o.SetLabels(labels)
	}
	return c
}

func (c *Component) NewJobSyncer(r interfaces.Reconcile) syncer.Interface {
	return syncer.NewObjectSyncer("Job", c.Owner, &c.Job, r.GetClient(), r.GetScheme(), c.MutateJob)
}

func (c *Component) MutateJob() error {
	c.Settings.MutatePod(&c.Job.Spec.Template)
	c.Runtime.MutatePod(&c.Job.Spec.Template)

	//	_ = mergo.Merge(&component.Job.Spec.Template.ObjectMeta, &component.Job.ObjectMeta)
	// component.Job.Spec.Template.ObjectMeta = component.Job.ObjectMeta

	c.Job.Spec.Template.Spec.RestartPolicy = corev1.RestartPolicyNever

	// args := []string{"/usr/local/bin/php", "/var/www/html/cron.php"}
	// component.Job.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Args = args

	return nil
}

func (c *Component) GetName() string {
	return fmt.Sprintf("%s-%s", c.Owner.Name, c.Name)
}

func (c *Component) GetObjects() []interfaces.Object {
	return []interfaces.Object{
		&c.Job,
	}
}
