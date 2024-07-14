package models

import (
	"fmt"
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
	return f.FieldNameSnakeCaseLower()
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
	return naming.ToUpperCamelCase(f.Type)
}

func (f Field) FieldNameDB() string {
	return f.FieldNameSnakeCaseLower()
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

func (f Field) ValueDB() string {
	if len(f.Value) == 0 {
		return "NULL"
	}

	return fmt.Sprintf("'%s'", f.Value)
}

func DB(f models.Field) string {
	if len(f.Value) == 0 {
		return "NULL"
	}

	return fmt.Sprintf("'%s'", f.Value)
}

func (f Field) Nullable() string {
	if f.Field.Nullable || strings.ToLower(f.Field.Type) == models.IDFieldType {
		return ""
	}

	return "NOT NULL"
}
