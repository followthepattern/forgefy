package featureset

import validation "github.com/go-ozzo/ozzo-validation"

type AppType string

const (
	Backend  AppType = "backend"
	Frontend AppType = "frontend"
)

type App struct {
	Name string  `yaml:"name"`
	Type AppType `yaml:"type"`
}

func (a App) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required),
		validation.Field(&a.Type, validation.In(Backend, Frontend)),
	)
}
