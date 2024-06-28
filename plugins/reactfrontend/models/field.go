package models

import (
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/naming"
)

type Field struct {
	models.Field
}

func (f Field) FieldNameGraphQL() string {
	return naming.ToLowerCamelCase(f.Name)
}

func (f Field) FieldTypeName() string {
	switch f.Type {
	case "string":
		return "string"
	case "int":
		return "number"
	}
	return f.Type
}

func (f Field) FieldTypeHTML() string {
	switch f.Type {
	case "string":
		return "text"
	}
	return "input"
}
