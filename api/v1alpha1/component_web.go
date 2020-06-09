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

package v1alpha1

func (app *Web) SetDefaults() {

	app.ObjectMeta.SetComponent("web")

	if len(app.Backend.Paths) == 0 {
		app.Backend.Paths = []string{"/"}
	}

	if &app.Backend.Port == nil || app.Backend.Port.Port == 0 {
		app.Backend.Port.Port = 80
	}
	if len(app.Port.Protocol) == 0 {
		app.Backend.Port.Protocol = "TCP"
	}
	if len(app.Backend.Port.Name) == 0 {
		app.Backend.Port.Name = "http"
	}

	if len(app.Image) == 0 {
		app.Image = "libresh/nextcloud:18.0.0"
	}

	if len(app.Settings) == 0 {
		app.Settings = []string{"web"}
	}

	app.Workload.SetDefaults()
}
