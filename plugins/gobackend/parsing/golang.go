package parsing

import (
	"strings"

	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/naming"
	"github.com/followthepattern/forgefy/specification/types"
)

func CreateGoType(t types.TypeRegistry) func(i interface{}) string {
	return func(i interface{}) string {
		switch tt := i.(type) {
		case models.Field:
			return goFieldType(t, tt)
		case specification.Feature:
			return goFeatureType(tt)
		}
		return "unknown"
	}
}

func goFeatureType(f specification.Feature) string {
	return naming.ToUpperCamelCase(f.FeatureName)
}

func goFieldType(t types.TypeRegistry, f models.Field) string {
	switch t.GetType(f.Type) {
	case types.String:
		return "types.String"
	case types.Number:
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
