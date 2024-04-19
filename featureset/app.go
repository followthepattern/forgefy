package featureset

import validation "github.com/go-ozzo/ozzo-validation"

type AppType string

const (
	Backend  AppType = "backend"
	Frontend AppType = "frontend"
)

type App struct {
	AppName  string    `yaml:"name"`
	AppType  AppType   `yaml:"type"`
	Features []Feature `yaml:"features"`
}

func (a App) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.AppName, validation.Required),
		validation.Field(&a.AppType, validation.In(Backend, Frontend)),
	)
}
