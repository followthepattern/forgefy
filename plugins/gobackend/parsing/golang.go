package parsing

import (
	"strings"

	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/naming"
)

func GoType(i interface{}) string {
	switch tt := i.(type) {
	case models.Field:
		return goFieldType(tt)
	case specification.Feature:
		return goFeatureType(tt)
	}
	return "unknown"
}

func goFeatureType(f specification.Feature) string {
	return naming.ToUpperCamelCase(f.FeatureName)
}

func goFieldType(f models.Field) string {
	switch f.Type {
	case "string":
		return "types.String"
	case "int":
		return "types.Int64"
	}
	return f.Type
}

func AsTag(f models.Field) string {
	return f.FieldNameSnakeCaseLower()
}

func AsURL(f specification.Feature) string {
	return strings.ToLower(naming.ToSnakeCase(f.FeatureName))
}
