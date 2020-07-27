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

type Database struct {
	Database *parameters.Parameter `json:"database,omitempty" env:"DB_NAME"`
	Host     *parameters.Parameter `json:"host,omitempty" env:"DB_HOST"`
	Port     *parameters.Parameter `json:"port,omitempty" env:"DB_PORT"`
	Type     *parameters.Parameter `json:"type,omitempty" env:"DB_TYPE"`
	Username *parameters.Parameter `json:"username,omitempty" env:"DB_USER"`
	Password *parameters.Parameter `json:"password,omitempty" env:"DB_PASSWORD"`
}

func (s *Database) SetDefaults() {
	if s.Database == nil {
		s.Database = &parameters.Parameter{}
	}
	if len(s.Database.Value) == 0 && len(s.Database.ValueFrom.Ref) == 0 {
		s.Database.Value = "nextcloud"
	}

	if s.Port == nil {
		s.Port = &parameters.Parameter{}
	}
	if len(s.Port.Value) == 0 && len(s.Port.ValueFrom.Ref) == 0 {
		s.Port.Value = "5425"
	}

	if s.Type == nil {
		s.Type = &parameters.Parameter{}
	}
	if len(s.Type.Value) == 0 && len(s.Type.ValueFrom.Ref) == 0 {
		s.Type.Value = "pgsql"
	}
}

func (s *Database) GetParameters() *parameters.Parameters {
	params, _ := parameters.Marshal(*s)
	return params
}
