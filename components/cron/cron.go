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

package cron

import (
	"fmt"

	"github.com/presslabs/controller-util/syncer"

	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"

	common "git.indie.host/nextcloud-operator/components/common"
	interfaces "git.indie.host/nextcloud-operator/interfaces"
)

type Component struct {
	Name string
	*common.Common
	CronJob batchv1beta1.CronJob
}

func CreateAndInit(common *common.Common) *Component {
	c := &Component{}
	c.Name = "cron"
	c.Common = common

	labels := c.Labels("cronjob")
	c.CronJob.SetLabels(labels)

	objects := c.GetObjects()
	for _, o := range objects {
		o.SetName(c.GetName())
		o.SetNamespace(c.Owner.Namespace)
		o.SetLabels(labels)
	}

	return c
}

func (c *Component) NewCronJobSyncer(r interfaces.Reconcile) syncer.Interface {
	return syncer.NewObjectSyncer("CronJob", c.Owner, &c.CronJob, r.GetClient(), r.GetScheme(), c.MutateCronJob)
}

func (c *Component) MutateCronJob() error {
	c.Settings.MutatePod(&c.CronJob.Spec.JobTemplate.Spec.Template)
	c.Runtime.MutatePod(&c.CronJob.Spec.JobTemplate.Spec.Template)

	c.CronJob.Spec.Schedule = "*/15 * * * *"
	c.CronJob.Spec.JobTemplate.ObjectMeta = c.CronJob.ObjectMeta
	c.CronJob.Spec.JobTemplate.Spec.Template.ObjectMeta = c.CronJob.ObjectMeta
	c.CronJob.Spec.JobTemplate.Spec.Template.Spec.RestartPolicy = corev1.RestartPolicyNever

	args := []string{"/usr/local/bin/php", "/var/www/html/cron.php"}
	c.CronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Args = args

	return nil
}

func (c *Component) GetName() string {
	return fmt.Sprintf("%s-%s", c.Owner.Name, c.Name)
}

func (c *Component) GetObjects() []interfaces.Object {
	return []interfaces.Object{
		&c.CronJob,
	}
}
