package models

import (
	"strings"

	"github.com/followthepattern/forgefy/specification/naming"
	validation "github.com/go-ozzo/ozzo-validation"
)

const IDFieldType = "id"

type Field struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Value    string `yaml:"-"`
	Nullable bool   `yaml:"nullable"`
}

func (f Field) FieldName() string {
	return f.Name
}

func (f Field) FieldType() string {
	return f.Type
}

func (f Field) FieldValue() string {
	return f.Value
}

func (f Field) FieldNameCamelCase() string {
	return naming.ToLowerCamelCase(f.Name)
}

func (f Field) FieldNameCamelCaseUpper() string {
	return naming.ToUpperCamelCase(f.Name)
}

func (f Field) FieldNameSnakeCase() string {
	return naming.ToSnakeCase(f.Name)
}

func (f Field) FieldNameSnakeCaseLower() string {
	return strings.ToLower(naming.ToSnakeCase(f.Name))
}

func (f Field) FieldNameHumanReadable() string {
	return naming.ToHumanReadable(f.Name)
}

func (f Field) FieldNameVarName() string {
	return strings.ToLower(f.Name)
}

func (f Field) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.Name, validation.Required),
		validation.Field(&f.Type, validation.Required),
	)
}
