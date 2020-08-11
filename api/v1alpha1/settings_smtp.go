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
	"k8s.libre.sh/controller-utils/application/settings/parameters"
)

type SMTP struct {
	Username      *parameters.Parameter `json:"username,omitempty" env:"SMTP_USERNAME"`
	Password      *parameters.Parameter `json:"password,omitempty" env:"SMTP_PASSWORD"`
	FromAdress    *parameters.Parameter `json:"fromAddress,omitempty" env:"MAIL_FROM_ADDRESS"`
	Domain        *parameters.Parameter `json:"domain,omitempty" env:"MAIL_DOMAIN"`
	Secure        *parameters.Parameter `json:"secure,omitempty" env:"SMTP_SECURE"`
	AuthType      *parameters.Parameter `json:"authType,omitempty" env:"SMTP_AUTHTYPE"`
	Debug         *parameters.Parameter `json:"debug,omitempty" env:"SMTP_DEBUG"`
	Host          *parameters.Parameter `json:"host,omitempty" env:"SMTP_HOST"`
	Port          *parameters.Parameter `json:"port,omitempty" env:"SMTP_PORT"`
	TemplateClass *parameters.Parameter `json:"templateClass,omitempty" env:"SMTP_TEMPLATE_CLASS"`
	PlainTextOnly *parameters.Parameter `json:"plainTextOnly,omitempty" env:"SMTP_SEND_PLAINTEXT_ONLY"`
}

func (d *SMTP) SetDefaults() {

}

func (s *SMTP) GetParameters() *parameters.Parameters {
	params, _ := parameters.Marshal(*s)
	return params
}
