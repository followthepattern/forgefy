package specification

import (
	"strings"

	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/naming"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Feature struct {
	FeatureName string         `yaml:"name"`
	Fields      []models.Field `yaml:"fields"`
}

func (f Feature) FeatureNameCamelCase() string {
	return naming.LowerFirst(f.FeatureName)
}

func (f Feature) FeatureTypeName() string {
	return naming.CapitalizeFirst(f.FeatureName)
}

func (f Feature) FeatureHumanReadableName() string {
	return naming.CapitalizeFirst(f.FeatureName)
}

func (f Feature) FeatureNameCamelCaseUpper() string {
	return naming.CapitalizeFirst(f.FeatureName)
}

func (f Feature) IDField() models.Field {
	return f.Fields[0]
}

func (f Feature) FeatureNameDir() string {
	return strings.ToLower(f.FeatureName)
}

func (f Feature) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.FeatureName, validation.Required),
		validation.Field(&f.Fields, validation.Required),
	)
}
