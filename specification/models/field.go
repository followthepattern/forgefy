package models

import (
	"fmt"
	"strings"

	"github.com/followthepattern/forgefy/datagenerator"
	"github.com/followthepattern/forgefy/specification/naming"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

const IDFieldName = "ID"

type Field struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Value    string `yaml:"value"`
	Nullable bool   `yaml:"nullable" default:"true"`
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
	return naming.LowerFirst(f.Name)
}

func (f Field) FieldNameCamelCaseUpper() string {
	return naming.CapitalizeFirst(f.Name)
}

func (f Field) FieldNameSnakeCase() string {
	return naming.ToSnakeCase(f.Name)
}

func (f Field) FieldNameHumanReadable() string {
	return naming.ToHumanReadable(f.Name)
}

func (f Field) FieldNameVarName() string {
	return strings.ToLower(f.Name)
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
