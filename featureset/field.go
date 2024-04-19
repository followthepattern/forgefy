package featureset

import validation "github.com/go-ozzo/ozzo-validation"

type Field struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

func (f Field) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.Name, validation.Required),
		validation.Field(&f.Type, validation.Required),
	)
}
