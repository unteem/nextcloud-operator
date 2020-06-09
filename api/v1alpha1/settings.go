package v1alpha1

import (
	"k8s.libre.sh/application/settings"
)

type Settings struct {
	CreateOptions settings.CreateOptions `json:"createOptions,omitempty"`
	Sources       []settings.Source      `json:"sources,omitempty"`
	AppSettings   AppSettings            `json:"app,omitempty"`
	Web           WebSettings            `json:"web,omitempty"`
}

func (s *Settings) SetDefaults() {
	s.AppSettings.SetDefaults()
	s.Web.SetDefaults()
}
