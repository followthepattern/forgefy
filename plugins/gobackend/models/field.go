package models

import (
	"strings"

	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/naming"
)

type Field struct {
	models.Field
}

func (f Field) FieldTypeGo() string {
	switch f.Type {
	case "string":
		return "types.String"
	case "int":
		return "types.Int64"
	}
	return f.Type
}

func (f Field) FieldNameAsTag() string {
	return strings.ToLower(f.Name)
}

func (f Field) FieldNameGraphQL() string {
	return naming.ToLowerCamelCase(f.Name)
}

func (f Field) FieldTypeGraphQL() string {
	switch f.Type {
	case "string":
		return "String"
	case "int":
		return "Int64"
	}
	return naming.CapitalizeFirst(f.Type)
}

func (f Field) FieldNameDB() string {
	return naming.ToSnakeCase(f.Name)
}

func (f Field) FieldTypeDB() string {
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

func (f Field) FieldValueDB() string {
	return f.Value
}

func (f Field) Nullable() string {
	if f.Field.Nullable || f.Name != models.IDFieldName {
		return ""
	}

	return "NOT NULL"
}
