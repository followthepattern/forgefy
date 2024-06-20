package specification

import (
	"strings"

	"github.com/followthepattern/forgefy/specification/models"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Feature struct {
	FeatureName string         `yaml:"name"`
	Fields      []models.Field `yaml:"fields"`
}

func (f Feature) FeatureNameLowerCamel() string {
	return LowerFirst(f.FeatureName)
}

func (f Feature) IDField() models.Field {
	return f.Fields[0]
}

func (f Feature) ToDirName() string {
	return strings.ToLower(f.FeatureName)
}

func (f Feature) FeatureGraphQLName() string {
	return CapitalizeFirst(f.FeatureName)
}

func (f Feature) FeatureAsDirName() string {
	return strings.ToLower(f.FeatureName)
}

func (f Feature) FeatureURL() string {
	return strings.ToLower(f.FeatureName)
}

func (f Feature) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.FeatureName, validation.Required),
		validation.Field(&f.Fields, validation.Required),
	)
}
