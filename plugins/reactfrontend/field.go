package reactfrontend

import (
	"strings"

	"github.com/followthepattern/forgefy/specification"
)

type Field struct {
	specification.Field
}

func (f Field) VariableName() string {
	return strings.ToLower(f.Name)
}

func (f Field) FieldName() string {
	return strings.ToLower(f.Name)
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
