package gobackend

import (
	"strings"

	"github.com/followthepattern/forgefy/specification"
)

type Field struct {
	specification.Field
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
