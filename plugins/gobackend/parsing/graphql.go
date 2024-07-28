package parsing

import (
	"strings"

	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/naming"
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

func TypeGraphQL(f models.Field) string {
	switch f.Type {
	case "string":
		return "String"
	case "int":
		return "Int64"
	}
	return naming.ToUpperCamelCase(f.Type)
}
