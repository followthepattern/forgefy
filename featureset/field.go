package featureset

import (
	"github.com/followthepattern/forgefy/datagenerator"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type Field struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

func (f Field) RandomValue() string {
	switch f.Type {
	case "string":
		return datagenerator.String(10)
	case "uuid":
		return uuid.NewString()
	}
	return ""
}

func (f Field) DBType() string {
	switch f.Type {
	case "string":
		return "VARCHAR"
	}
	return ""
}

func (f Field) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.Name, validation.Required),
		validation.Field(&f.Type, validation.Required),
	)
}
