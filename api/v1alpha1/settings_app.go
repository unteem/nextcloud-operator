package v1alpha1

import (
	"k8s.libre.sh/application/settings"
	"k8s.libre.sh/meta"
	"k8s.libre.sh/objects"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type AppSettings struct {
	CreateOptions settings.CreateOptions `json:"createOptions,omitempty"`
	Sources       []settings.Source      `json:"sources,omitempty"`
	Database      Database               `json:"database,omitempty"`
	SMTP          SMTP                   `json:"smtp,omitempty"`
	General       General                `json:"general,omitempty"`
	ObjectStore   ObjectStore            `json:"objectStore,omitempty"`
	Redis         Redis                  `json:"redis,omitempty"`
}

func (s *AppSettings) GetMeta() meta.Instance { return s.CreateOptions.CommonMeta }

func (s *AppSettings) SetDefaults() {

	//	s.CreateOptions.Init()
	//	s.CreateOptions.CommonMeta.Labels["app.kubernetes.io/component"] = "app"

	s.General.SetDefaults()
	s.Database.SetDefaults()
	s.SMTP.SetDefaults()
	s.ObjectStore.SetDefaults()
	s.Redis.SetDefaults()
}

func (s *AppSettings) GetConfig() settings.Config {

	params := *s.General.GetParameters()
	params = append(params, *s.Database.GetParameters()...)
	params = append(params, *s.SMTP.GetParameters()...)
	params = append(params, *s.ObjectStore.GetParameters()...)
	params = append(params, *s.Redis.GetParameters()...)

	settings := &settings.ConfigSpec{
		Parameters: &params,
		Sources:    s.Sources,
	}

	return settings
}

func (s *AppSettings) GetObjects() map[int]objects.Object {

	return nil
}

func (s *AppSettings) Init(c client.Client) error {

	return nil
}

func (s *AppSettings) GetCreateOptions() *settings.CreateOptions {
	return &s.CreateOptions
}
