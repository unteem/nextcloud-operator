package v1alpha1

import (
	"k8s.libre.sh/application/settings"
	"k8s.libre.sh/meta"
	"k8s.libre.sh/objects"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Settings struct {
	CreateOptions settings.CreateOptions `json:"createOptions,omitempty"`
	Sources       []settings.Source      `json:"sources,omitempty"`
	Database      Database               `json:"database,omitempty"`
	SMTP          SMTP                   `json:"smtp,omitempty"`
	General       General                `json:"general,omitempty"`
	ObjectStore   ObjectStore            `json:"storage,omitempty"`
	Redis         Redis                  `json:"cache,omitempty"`
}

func (s *Settings) GetMeta() meta.Instance { return s.CreateOptions.CommonMeta }

func (s *Settings) SetDefaults() {
	if s.CreateOptions.CommonMeta == nil {
		s.CreateOptions.CommonMeta = new(meta.ObjectMeta)
	}

	if len(s.CreateOptions.CommonMeta.Labels) == 0 {
		s.CreateOptions.CommonMeta.Labels = make(map[string]string)
	}

	s.CreateOptions.CommonMeta.Labels["app.kubernetes.io/component"] = "settings"

	//	meta.SetObjectMeta(i, s.ObjectMeta)

	s.General.SetDefaults()
	s.Database.SetDefaults()
	s.SMTP.SetDefaults()
	s.ObjectStore.SetDefaults()
	s.Redis.SetDefaults()
}

func (s *Settings) GetConfig() settings.Config {

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

func (s *Settings) GetSecretMeta() meta.Instance {
	/* 	if s.CreateOptions.SecretMeta == nil {
		s.CreateOptions.SecretMeta = new(meta.ObjectMeta)
	} */
	return s.CreateOptions.SecretMeta
}
func (s *Settings) GetConfigMapMeta() meta.Instance {
	/* 	if s.CreateOptions.ConfigMeta == nil {
		s.CreateOptions.ConfigMeta = new(meta.ObjectMeta)
	} */
	return s.CreateOptions.ConfigMeta
}

func (s *Settings) GetObjects() map[int]objects.Object {

	return nil
}

func (s *Settings) Init(c client.Client) error {

	return nil
}

func (s *Settings) GetCreateOptions() *settings.CreateOptions {
	return &s.CreateOptions
}
