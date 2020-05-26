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

type AppSettings struct {
	OverwriteCLI      parameters.Parameter `json:"overwriteCLI,omitempty" env:"OVERWRITE_CLI_URL"`
	OverwriteProtocol parameters.Parameter `json:"overwriteProtocol,omitempty" env:"OVERWRITE_PROTOCOL"`
	DataDirectory     parameters.Parameter `json:"dataDirectory,omitempty" env:"DATA_DIRECTORY"`
	Debug             parameters.Parameter `json:"debug,omitempty" env:"DEBUG"`
	ReadOnly          parameters.Parameter `json:"readOnly,omitempty" env:"CONFIG_READONLY"`
	UpdateChecker     parameters.Parameter `json:"updateChecker,romitempty" env:"UPDATE_CHECKER"`
	UpdateURL         parameters.Parameter `json:"udpateURL,omitempty" env:"OVERWRITECLI"`
	UpdateChannel     parameters.Parameter `json:"updateChannel,omitempty" env:"UPDATE_URL"`
	UpdateDisable     parameters.Parameter `json:"updateDisable,omitempty" env:"UPDATE_CHANNEL"`
	BruteForce        parameters.Parameter `json:"bruteforce,omitempty" env:"UPDATE_DISABLE_WEB"`
}

func (s *AppSettings) GetParameters() *parameters.Parameters {
	params, _ := parameters.Marshal(*s)
	return &params
}

type AppSecrets struct {
	InstanceID    parameters.Parameter `json:"instanceID,omitempty" env:"INSTANCE_ID"`
	PasswordSalt  parameters.Parameter `json:"passwordSalt,omitempty" env:"PASSWORD_SALT"`
	Secret        parameters.Parameter `json:"secret,omiempty" env:"SECRET"`
	AdminPassword parameters.Parameter `json:"adminPassword,omitempty" env:"ADMIN_PASSWORD"`
	AdminUsername parameters.Parameter `json:"adminUsername,omitempty" env:"ADMIN_USERNAME"`
}

func (s *AppSecrets) GetParameters() *parameters.Parameters {
	params, _ := parameters.Marshal(*s)
	return &params
}

type AppStore struct {
	// StoreEnabled defines if the app store is enabled
	StoreEnabled parameters.Parameter `json:"storeEnabled,omitempty" env:"APPS_STORE_ENABLE"`
	// Default defines the default app
	Default parameters.Parameter `json:"default,omitempty" env:"APPS_DEFAULT"`
	// StoreUrl defines the URL for the app store
	StoreURL parameters.Parameter `json:"storeURL,omitempty" env:"APPS_STORE_URL"`
}

func (s *AppStore) GetParameters() *parameters.Parameters {
	params, _ := parameters.Marshal(*s)
	return &params
}

type Locales struct {
	// Default defines the default language
	Default       parameters.Parameter `json:"default,omitempty" env:"DEFAULT_LANGUAGE"`
	Locale        parameters.Parameter `json:"locale,omitempty" env:"DEFAULT_LOCALE"`
	ForceLanguage parameters.Parameter `json:"forceLanguage,omitempty" env:"FORCE_LANGUAGE"`
	ForceLocale   parameters.Parameter `json:"forceLocale,omitempty" env:"DEFAULT_LOCALE"`
}

func (s *Locales) GetParameters() *parameters.Parameters {
	params, _ := parameters.Marshal(*s)
	return &params
}

type General struct {
	AppStore    AppStore `json:"appStore,omitempty"`
	Locales     Locales  `json:"locales,omitempty"`
	AppSecrets  `json:",inline"`
	AppSettings `json:",inline"`
}

func (s *General) SetDefaults() *parameters.Parameters {
	return &parameters.Parameters{}
}

func (s *General) GetParameters() *parameters.Parameters {
	params := append(*s.AppStore.GetParameters(), *s.Locales.GetParameters()...)
	params = append(params, *s.AppSecrets.GetParameters()...)
	params = append(params, *s.AppSettings.GetParameters()...)
	return &params
}
