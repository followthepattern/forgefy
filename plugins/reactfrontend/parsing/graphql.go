package parsing

import (
	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/naming"
)

func GraphQLName(i interface{}) string {
	switch tt := i.(type) {
	case specification.Feature:
		return featureGrapQLName(tt)
	case models.Field:
		return fieldGraphQLName(tt)
	}
	return "<unknown>"
}

func featureGrapQLName(f specification.Feature) string {
	return naming.ToUpperCamelCase(f.FeatureName)
}

func fieldGraphQLName(f models.Field) string {
	return naming.ToLowerCamelCase(f.Name)
}
