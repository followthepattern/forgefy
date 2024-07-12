package specification

import (
	"strings"

	"github.com/followthepattern/forgefy/specification/naming"
	validation "github.com/go-ozzo/ozzo-validation"
)

type AppType string

type App struct {
	Product
	AppName             string    `yaml:"name"`
	AppType             AppType   `yaml:"type"`
	Features            []Feature `yaml:"features"`
	AppPort             int       `yaml:"port"`
	CountOfRandomValues int       `yaml:""`
}

func (a App) AppNameDir() string {
	return strings.ToLower(naming.ToSnakeCase(a.AppName))
}

func (a App) AppNameCamelCase() string {
	return naming.ToLowerCamelCase(a.AppName)
}

func (a App) LoopLimit() []int {
	count := 10

	if a.CountOfRandomValues != 0 {
		count = a.CountOfRandomValues
	}

	return make([]int, count)
}

func (a App) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.AppName, validation.Required),
		validation.Field(&a.AppType, validation.Required),
	)
}
