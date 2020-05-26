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

import "k8s.libre.sh/application/settings/parameters"

type Database struct {
	Database parameters.Parameter `json:"database,omitempty" env:"DB_NAME"`
	Host     parameters.Parameter `json:"host,omitempty" env:"DB_HOST"`
	Port     parameters.Parameter `json:"port,omitempty" env:"DB_PORT"`
	Type     parameters.Parameter `json:"type,omitempty" env:"DB_TYPE"`
	Username parameters.Parameter `json:"username,omitempty" env:"DB_USERNAME"`
	Password parameters.Parameter `json:"password,omitempty" env:"DB_PASSWORD"`
}

func (s *Database) SetDefaults() {
	/* 	if len(d.Database) == 0 {
	   		d.Database = "nextcloud"
	   	}
	   	if len(d.Host) == 0 {
	   		d.Host = "nextcloud-database"
	   	}
	   	if len(d.Port) == 0 {
	   		d.Port = "5425"
	   	}
	   	if len(d.Type) == 0 {
	   		d.Type = "pgsql"
	   	} */
}

func (s *Database) GetParameters() *parameters.Parameters {
	s.SetDefaults()
	params, _ := parameters.Marshal(*s)
	return &params
}

/* func (s *Database) GetParameters() parameters.Parameters {

	s.SetDefaults()

	params, _ := parameters.Marshal(s.DatabaseConfig)
	secretParams, _ := parameters.Marshal(s.DatabaseSecrets)

	for _, p := range secretParams {
		// TODO Enforce secret type ?
		if len(p.Type) == 0 {
			p.Type = parameters.SecretParameter
		}
		if len(p.ValueFrom.Ref) == 0 && len(p.Generate) == 0 && len(p.Value) == 0 {
			p.Generate = parameters.GenerateRand12
		}
		// TODO TOFIX
		if len(p.MountType) == 0 {
			p.MountType = parameters.MountEnvFile
		}
	}

	params = append(params, secretParams...)

	return params
}
*/
