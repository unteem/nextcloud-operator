package v1alpha1

import (
	"k8s.libre.sh/application/settings"
	"k8s.libre.sh/meta"
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
