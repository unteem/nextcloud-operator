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

type General struct {
	AppStore       AppStore `json:"appStore,omitempty"`
	Locales        Locales  `json:"locales,omitempty"`
	GlobalSecrets  `json:",inline"`
	GlobalSettings `json:",inline"`
}

func (s *General) SetDefaults() {
	s.AppStore.SetDefaults()
	s.Locales.SetDefaults()
	s.GlobalSecrets.SetDefaults()
	s.GlobalSettings.SetDefaults()

}

func (s *General) GetParameters() *parameters.Parameters {
	params := append(*s.AppStore.GetParameters(), *s.Locales.GetParameters()...)
	params = append(params, *s.GlobalSecrets.GetParameters()...)
	params = append(params, *s.GlobalSettings.GetParameters()...)
	return &params
}

type GlobalSettings struct {
	Domains           *parameters.Parameter `json:"domains,omitempty" env:"NEXTCLOUD_TRUSTED_DOMAINS"`
	OverwriteCLI      *parameters.Parameter `json:"overwriteCLI,omitempty" env:"OVERWRITE_CLI_URL"`
	OverwriteProtocol *parameters.Parameter `json:"overwriteProtocol,omitempty" env:"OVERWRITE_PROTOCOL"`
	DataDirectory     *parameters.Parameter `json:"dataDirectory,omitempty" env:"DATA_DIRECTORY"`
	Debug             *parameters.Parameter `json:"debug,omitempty" env:"DEBUG"`
	ReadOnly          *parameters.Parameter `json:"readOnly,omitempty" env:"CONFIG_READONLY"`
	UpdateChecker     *parameters.Parameter `json:"updateChecker,omitempty" env:"UPDATE_CHECKER"`
	UpdateURL         *parameters.Parameter `json:"udpateURL,omitempty" env:"UPDATE_URL"`
	UpdateChannel     *parameters.Parameter `json:"updateChannel,omitempty" env:"UPDATE_CHANNEL"`
	UpdateDisable     *parameters.Parameter `json:"updateDisable,omitempty" env:"UPDATE_DISABLE_WEB"`
	BruteForce        *parameters.Parameter `json:"bruteforce,omitempty" env:"BRUTEFORCE"`
	Version           *parameters.Parameter `json:"version,omitempty" env:"VERSION"`
}

func (s *GlobalSettings) GetParameters() *parameters.Parameters {
	params, _ := parameters.Marshal(*s)
	return params
}

func (s *GlobalSettings) SetDefaults() {
	if s.ReadOnly == nil {
		s.ReadOnly = &parameters.Parameter{}
	}
	if len(s.ReadOnly.Value) == 0 && len(s.ReadOnly.ValueFrom.Ref) == 0 {
		s.ReadOnly.Value = "true"
	}

	if s.DataDirectory == nil {
		s.DataDirectory = &parameters.Parameter{}
	}
	if len(s.DataDirectory.Value) == 0 && len(s.DataDirectory.ValueFrom.Ref) == 0 {
		s.DataDirectory.Value = "/var/www/html/data"
	}

	if s.UpdateChecker == nil {
		s.UpdateChecker = &parameters.Parameter{}
	}
	if len(s.UpdateChecker.Value) == 0 && len(s.UpdateChecker.ValueFrom.Ref) == 0 {
		s.UpdateChecker.Value = "false"
	}

	if s.UpdateDisable == nil {
		s.UpdateDisable = &parameters.Parameter{}
	}
	if len(s.UpdateDisable.Value) == 0 && len(s.UpdateDisable.ValueFrom.Ref) == 0 {
		s.UpdateDisable.Value = "true"
	}

	if s.Domains == nil {
		s.Domains = &parameters.Parameter{}
	}
	if len(s.Domains.Value) == 0 && len(s.Domains.ValueFrom.Ref) == 0 {
		/* 		s.Domains.Value = "{{ .components.web.network.hostname }}"
		   		s.Domains.Generate = parameters.GenerateTemplate
		   		s.Domains.Type = parameters.SecretParameter */
	}
}

type GlobalSecrets struct {
	InstanceID    *parameters.Parameter `json:"instanceID,omitempty" env:"INSTANCE_ID"`
	PasswordSalt  *parameters.Parameter `json:"passwordSalt,omitempty" env:"PASSWORD_SALT"`
	Secret        *parameters.Parameter `json:"secret,omitempty" env:"SECRET"`
	AdminPassword *parameters.Parameter `json:"adminPassword,omitempty" env:"NEXTCLOUD_ADMIN_PASSWORD"`
	AdminUsername *parameters.Parameter `json:"adminUsername,omitempty" env:"NEXTCLOUD_ADMIN_USER"`
}

func (s *GlobalSecrets) SetDefaults() {
	// TODO return warning if value is defined and ignore it. Or use secretParameter type to enforce no values
	if s.InstanceID == nil {
		s.InstanceID = &parameters.Parameter{}
	}
	if len(s.InstanceID.Value) == 0 && len(s.InstanceID.ValueFrom.Ref) == 0 {
		s.InstanceID.Generate = parameters.GenerateRand12
	}
	if s.PasswordSalt == nil {
		s.PasswordSalt = &parameters.Parameter{}
	}
	if len(s.PasswordSalt.Value) == 0 && len(s.PasswordSalt.ValueFrom.Ref) == 0 {
		s.PasswordSalt.Generate = parameters.GenerateRand12
	}

	if s.Secret == nil {
		s.Secret = &parameters.Parameter{}
	}
	if len(s.Secret.Value) == 0 && len(s.Secret.ValueFrom.Ref) == 0 {
		s.Secret.Generate = parameters.GenerateRand12
	}
	if s.AdminPassword == nil {
		s.AdminPassword = &parameters.Parameter{}
	}

	if len(s.AdminPassword.Value) == 0 && len(s.AdminPassword.ValueFrom.Ref) == 0 {
		s.AdminPassword.Generate = parameters.GenerateRand12
	}

	if s.AdminUsername == nil {
		s.AdminUsername = &parameters.Parameter{}
	}
	if len(s.AdminUsername.Value) == 0 && len(s.AdminUsername.ValueFrom.Ref) == 0 {
		s.AdminUsername.Value = "admin"
	}
}

func (s *GlobalSecrets) GetParameters() *parameters.Parameters {
	params, _ := parameters.Marshal(*s)
	return params
}

type AppStore struct {
	// StoreEnabled defines if the app store is enabled
	StoreEnabled *parameters.Parameter `json:"storeEnabled,omitempty" env:"APPS_STORE_ENABLE"`
	// Default defines the default app
	Default *parameters.Parameter `json:"default,omitempty" env:"APPS_DEFAULT"`
	// StoreUrl defines the URL for the app store
	StoreURL *parameters.Parameter `json:"storeURL,omitempty" env:"APPS_STORE_URL"`
}

func (s *AppStore) SetDefaults() {
	if s.StoreEnabled == nil {
		s.StoreEnabled = &parameters.Parameter{}
	}
	if len(s.StoreEnabled.Value) == 0 || len(s.StoreEnabled.ValueFrom.Ref) == 0 {
		s.StoreEnabled.Value = "false"
	}
}

func (s *AppStore) GetParameters() *parameters.Parameters {
	params, _ := parameters.Marshal(*s)
	return params
}

type Locales struct {
	// Default defines the default language
	Default       *parameters.Parameter `json:"default,omitempty" env:"DEFAULT_LANGUAGE"`
	Locale        *parameters.Parameter `json:"locale,omitempty" env:"DEFAULT_LOCALE"`
	ForceLanguage *parameters.Parameter `json:"forceLanguage,omitempty" env:"FORCE_LANGUAGE"`
	ForceLocale   *parameters.Parameter `json:"forceLocale,omitempty" env:"DEFAULT_LOCALE"`
}

func (s *Locales) SetDefaults() {}

func (s *Locales) GetParameters() *parameters.Parameters {
	params, _ := parameters.Marshal(*s)
	return params
}
