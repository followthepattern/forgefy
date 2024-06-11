package models

import (
	"strings"

	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/models"
)

type Field struct {
	models.Field
}

func (f Field) SchemaVarName() string {
	return strings.ToLower(f.Name)
}

func (f Field) SchemaType() string {
	switch f.Type {
	case "string":
		return "String"
	case "int":
		return "Int64"
	}
	return specification.CapitalizeFirst(f.Type)
}

func (f Field) DBType() string {
	switch f.Type {
	case "string":
		return "varchar"
	case "bool":
		return "boolean"
	case "int":
		return "integer"
	}
	return ""
}

func (f Field) DBColumnName() string {
	return strings.ToLower(f.Name)
}

func (f Field) DBValue() string {
	return f.Value
}

func (f Field) Nullable() string {
	if f.Field.Nullable || f.Name != models.IDFieldName {
		return ""
	}

	return "NOT NULL"
}

func (f Field) GoType() string {
	switch f.Type {
	case "string":
		return "types.String"
	case "int":
		return "types.Int64"
	}
	return f.Type
}

func (f Field) VariableName() string {
	return strings.ToLower(f.Name)
}

func (f Field) TagName() string {
	return strings.ToLower(f.Name)
}

func (f Field) FieldName() string {
	return specification.CapitalizeFirst(f.Name)
}