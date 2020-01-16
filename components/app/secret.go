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

	"github.com/presslabs/controller-util/rand"

	interfaces "git.indie.host/nextcloud-operator/interfaces"
)

var (
	generatedSalts = map[string]int{
		"instanceID":    10,
		"adminPassword": 12,
		"passwordSalt":  20,
		"secret":        20,
	}
)

func (app *App) NewSecretSyncer(r interfaces.Reconcile) syncer.Interface {
	return syncer.NewObjectSyncer("Secret", app.Owner, &app.Secret, r.GetClient(), r.GetScheme(), app.MutateSecret)
}

func (app *App) MutateSecret() error {
	labels := app.Labels("app")

	app.Runtime.MutateService(&app.Service)
	app.Secret.SetLabels(labels)
	data, err := app.GenSecretData()
	if err != nil {
		return err
	}
	app.Secret.Data = data

	return nil
}

func (app *App) GenSecretData() (map[string][]byte, error) {
	data := make(map[string][]byte)

	for name, size := range generatedSalts {
		random, err := rand.AlphaNumericString(size)
		if err != nil {
			return data, err
		}
		data[name] = []byte(random)
	}

	return data, nil
}
