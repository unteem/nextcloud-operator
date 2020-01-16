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

package controllers

import (
	"github.com/presslabs/controller-util/syncer"

	appsv1beta1 "git.indie.host/nextcloud-operator/api/v1beta1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
)

type Cron struct {
	*Common
	CronJob batchv1beta1.CronJob
}

func NewCron(nc *appsv1beta1.Nextcloud) *Cron {
	cron := &Cron{}
	cron.Common = NewCommon(nc)
	cron.Owner = nc
	cron.CronJob.SetName("test")
	cron.CronJob.SetNamespace(cron.Owner.Namespace)
	return cron
}

func (cron *Cron) NewCronJobSyncer(r *NextcloudReconciler) syncer.Interface {
	return syncer.NewObjectSyncer("CronJob", cron.Owner, &cron.CronJob, r.Client, r.Scheme, cron.MutateCronJob)
}

func (cron *Cron) MutateCronJob() error {
	cron.Settings.MutatePod(&cron.CronJob.Spec.JobTemplate.Spec.Template)
	cron.Runtime.MutatePod(&cron.CronJob.Spec.JobTemplate.Spec.Template)

	labels := cron.Labels("cronjob")
	cron.CronJob.SetLabels(labels)

	cron.CronJob.Spec.Schedule = "*/15 * * * *"
	cron.CronJob.Spec.JobTemplate.ObjectMeta = cron.CronJob.ObjectMeta
	cron.CronJob.Spec.JobTemplate.Spec.Template.ObjectMeta = cron.CronJob.ObjectMeta
	cron.CronJob.Spec.JobTemplate.Spec.Template.Spec.RestartPolicy = corev1.RestartPolicyNever

	args := []string{"/usr/local/bin/php", "/var/www/html/cron.php"}
	cron.CronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Args = args

	return nil
}
