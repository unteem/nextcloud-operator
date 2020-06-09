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

type ObjectStore struct {
	Bucket          parameters.Parameter `json:"bucket,omitempty" env:"OBJECTSTORE_S3_BUCKET"`
	Host            parameters.Parameter `json:"host,omitempty" env:"OBJECTSTORE_S3_HOST"`
	Port            parameters.Parameter `json:"port,omitempty" env:"OBJECTSTORE_S3_PORT"`
	AutoCreate      parameters.Parameter `json:"autocreate,omitempty" env:"OBJECTSTORE_S3_AUTOCREATE"`
	SSL             parameters.Parameter `json:"ssl,omitempty" env:"OBJECTSTORE_S3_SSL"`
	Region          parameters.Parameter `json:"region,omitempty" env:"OBJECTSTORE_S3_REGION"`
	PathStyle       parameters.Parameter `json:"pathStyle,omitempty" env:"OBJECTSTORE_S3_USEPATH_STYLE"`
	AccessKeyID     parameters.Parameter `json:"accessKeyID,omitempty" env:"OBJECTSTORE_S3_KEY"`
	SecretAccessKey parameters.Parameter `json:"secretAccessKey,omitempty" env:"OBJECTSTORE_S3_SECRET"`
}

func (s *ObjectStore) SetDefaults() {
	/* 	if len(d.Bucket) == 0 {
	   		d.Bucket = "nextcloud"
	   	}
	   	if len(d.Port) == 0 {
	   		d.Port = "443"
	   	}
	   	if len(d.Region) == 0 {
	   		d.Region = "default"
	   	} */
}

func (s *ObjectStore) GetParameters() *parameters.Parameters {
	params, _ := parameters.Marshal(*s)
	return &params
}
