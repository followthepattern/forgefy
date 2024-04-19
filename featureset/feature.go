package featureset

import validation "github.com/go-ozzo/ozzo-validation"

type Feature struct {
	FeatureName string  `yaml:"name"`
	Fields      []Field `yaml:"fields"`
}

func (f Feature) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.FeatureName, validation.Required),
		validation.Field(&f.Fields, validation.Required),
	)
}
