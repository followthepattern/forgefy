package specification

import (
	"fmt"
	"strings"

	"github.com/followthepattern/forgefy/datagenerator"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

const IDFieldName = "ID"

type Field struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Nullable bool   `yaml:"nullable" default:"true"`
}

func (f Field) RandomValue() string {
	if f.Name == IDFieldName {
		return uuid.NewString()
	}

	switch f.Type {
	case "string":
		return datagenerator.String(10)
	case "uuid":
		return uuid.NewString()
	case "int":
		return fmt.Sprint(datagenerator.RandomInt())
	}
	return ""
}

func (f Field) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.Name, validation.Required),
		validation.Field(&f.Type, validation.Required),
	)
}

func (f Field) FieldNameGraphQL() string {
	return strings.ToLower(f.Name)
}

func (f Field) FieldNameCamelCase() string {
	return strings.ToLower(f.Name)
}

func (f Field) FieldNameHumanReadable() string {
	return f.Name
}
