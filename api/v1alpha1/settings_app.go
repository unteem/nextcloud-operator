package v1alpha1

import (
	"fmt"

	"k8s.libre.sh/application/settings"
	"k8s.libre.sh/application/settings/parameters"
	"k8s.libre.sh/interfaces"
	"k8s.libre.sh/meta"
	"k8s.libre.sh/objects"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type AppSettings struct {
	CreateOptions settings.CreateOptions `json:"createOptions,omitempty"`
	Sources       *settings.Sources      `json:"sources,omitempty"`
	Database      *Database              `json:"database,omitempty"`
	SMTP          *SMTP                  `json:"smtp,omitempty"`
	General       *General               `json:"general,omitempty"`
	ObjectStore   *ObjectStore           `json:"objectStore,omitempty"`
	Redis         *Redis                 `json:"redis,omitempty"`
}

func (s *AppSettings) GetMeta() meta.Instance { return s.CreateOptions.CommonMeta }

func (s *AppSettings) SetDefaults() {

	if len(s.CreateOptions.CommonMeta.GetComponent()) == 0 {
		s.CreateOptions.CommonMeta.SetComponent("app")
	}

	meta.SetObjectMeta(s.CreateOptions.CommonMeta, s.CreateOptions.ConfigMeta)
	meta.SetObjectMeta(s.CreateOptions.CommonMeta, s.CreateOptions.SecretMeta)
	fmt.Println(s.CreateOptions.SecretMeta.Name)
	fmt.Println(s.CreateOptions.SecretMeta.Namespace)

	if s.Database == nil {
		s.Database = &Database{}
	}

	if s.SMTP == nil {
		s.SMTP = &SMTP{}
	}

	if s.General == nil {
		s.General = &General{}
	}

	if s.ObjectStore == nil {
		s.ObjectStore = &ObjectStore{}
	}
	if s.Redis == nil {
		s.Redis = &Redis{}
	}

	if s.Sources == nil {
		s.Sources = &settings.Sources{}
	}
	s.General.SetDefaults()
	s.Database.SetDefaults()
	s.SMTP.SetDefaults()
	s.ObjectStore.SetDefaults()
	s.Redis.SetDefaults()
}

func (s *AppSettings) GetConfig() settings.Config {

	settings := &settings.SettingsSpec{
		Parameters: s.GetParameters(),
		Sources:    s.Sources,
	}

	return settings
}

func (s *AppSettings) GetParameters() *parameters.Parameters {

	params := *s.General.GetParameters()
	params = append(params, *s.Database.GetParameters()...)
	params = append(params, *s.SMTP.GetParameters()...)
	params = append(params, *s.ObjectStore.GetParameters()...)
	params = append(params, *s.Redis.GetParameters()...)

	return &params

}

func (s *AppSettings) GetObjects() map[int]objects.Object {

	return nil
}

func (s *AppSettings) GetSources() *settings.Sources {

	return s.Sources
}

func (s *AppSettings) Init(c client.Client, owner interfaces.Object) error {

	err := settings.Init(s, c, owner)
	if err != nil {
		return err
	}

	return nil
}

func (s *AppSettings) GetCreateOptions() *settings.CreateOptions {
	return &s.CreateOptions
}
