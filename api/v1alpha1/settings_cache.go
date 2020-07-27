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

import (
	"k8s.libre.sh/application/settings/parameters"
)

type Redis struct {
	Username *parameters.Parameter `json:"username,omitempty" env:"REDIS_USERNAME"`
	Password *parameters.Parameter `json:"password,omitempty" env:"REDIS_PASSWORD"`
	Host     *parameters.Parameter `json:"host,omitempty" env:"REDIS_HOST"`
	Port     *parameters.Parameter `json:"port,omitempty" env:"REDIS_HOST_PORT"`
}

func (d *Redis) SetDefaults() {
}

func (s *Redis) GetParameters() *parameters.Parameters {
	params, _ := parameters.Marshal(*s)

	return params
}
