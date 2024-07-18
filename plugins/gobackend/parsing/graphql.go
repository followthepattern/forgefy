package parsing

import (
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/naming"
)

func NameGraphQL(f models.Field) string {
	return naming.ToLowerCamelCase(f.Name)
}

func TypeGraphQL(f models.Field) string {
	switch f.Type {
	case "string":
		return "String"
	case "int":
		return "Int64"
	}
	return naming.ToUpperCamelCase(f.Type)
}
