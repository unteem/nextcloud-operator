/*

Licensed under the GNU AFFERO GENERAL PUBLIC LICENSE Version 3 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.gnu.org/licenses/agpl-3.0.html

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"k8s.libre.sh/controller-utils/application/settings"
	"k8s.libre.sh/controller-utils/meta"
)

type Settings struct {
	CreateOptions *settings.CreateOptions `json:"createOptions,omitempty"`
	Sources       []*settings.Source      `json:"sources,omitempty"`
	AppSettings   *AppSettings            `json:"app,omitempty"`
	Web           *WebSettings            `json:"web,omitempty"`
}

func (s *Settings) Init() {

	if s.AppSettings == nil {
		s.AppSettings = &AppSettings{}
	}
	if s.Web == nil {
		s.Web = &WebSettings{}
	}

	if s.CreateOptions == nil {
		s.CreateOptions = &settings.CreateOptions{}
	}

	s.CreateOptions.Init()

	s.AppSettings.CreateOptions.Init()

	if s.Web.CreateOptions == nil {
		s.Web.CreateOptions = &settings.CreateOptions{}
	}

	s.Web.CreateOptions.Init()
}

func (s *Settings) SetDefaults() {
	if s.AppSettings == nil {
		s.AppSettings = &AppSettings{}
	}
	if s.Web == nil {
		s.Web = &WebSettings{}
	}

	meta.SetObjectMetaFromInstance(s.CreateOptions.CommonMeta, s.AppSettings.CreateOptions.CommonMeta)
	meta.SetObjectMetaFromInstance(s.CreateOptions.CommonMeta, s.Web.CreateOptions.CommonMeta)

	s.AppSettings.SetDefaults()
	s.Web.SetDefaults()
}
