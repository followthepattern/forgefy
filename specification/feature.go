package specification

import validation "github.com/go-ozzo/ozzo-validation"

type Feature struct {
	FeatureName         string  `yaml:"name"`
	Fields              []Field `yaml:"fields"`
	CountOfRandomValues int
}

func (f Feature) TemplateIterator() []int {
	count := 10

	if f.CountOfRandomValues != 0 {
		count = f.CountOfRandomValues
	}

	return make([]int, count)
}

func (f Feature) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.FeatureName, validation.Required),
		validation.Field(&f.Fields, validation.Required),
	)
}
