package featureset

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gopkg.in/yaml.v3"
)

type FeatureSet struct {
	Version     string `yaml:"version"`
	ProductName string `yaml:"product_name"`
	Apps        []App  `yaml:"apps"`
}

func (fs FeatureSet) Validate() error {
	return validation.ValidateStruct(&fs,
		validation.Field(&fs.ProductName, validation.Required),
		validation.Field(&fs.Apps, validation.Required),
	)
}

func UnmarshalYaml(data []byte) (fs FeatureSet, err error) {
	err = yaml.Unmarshal(data, &fs)
	if err != nil {
		return
	}

	err = fs.Validate()
	return
}
