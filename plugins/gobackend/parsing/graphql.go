package parsing

import (
	"strings"

	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/naming"
	"github.com/followthepattern/forgefy/specification/types"
)

func PackageName(i interface{}) string {
	switch tt := i.(type) {
	case specification.App:
		return appPackageName(tt)
	case specification.Feature:
		return featurePackageName(tt)
	}

	return "unknown"
}

func appPackageName(a specification.App) string {
	return strings.ToLower(a.AppName)
}

func featurePackageName(f specification.Feature) string {
	return strings.ToLower(naming.ToSnakeCase(f.FeatureName))
}

func NameGraphQL(i interface{}) string {
	switch tt := i.(type) {
	case models.Field:
		return fieldGraphQLName(tt)
	case specification.Feature:
		return featureGraphQLName(tt)
	}
	return "unknown"
}

func fieldGraphQLName(f models.Field) string {
	return naming.ToLowerCamelCase(f.Name)
}

func featureGraphQLName(f specification.Feature) string {
	return naming.ToUpperCamelCase(f.FeatureName)
}

func CreateTypeGraphQL(t types.TypeRegistry) func(models.Field) string {
	return func(f models.Field) string {
		switch t.GetType(f.Type) {
		case types.String:
			return "String"
		case types.Number:
			return "Int64"
		case types.Date:
		case types.DateTime:
			return "Time"
		}
		return naming.ToUpperCamelCase(f.Type)
	}
}
