package specification

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gopkg.in/yaml.v3"
)

type Product struct {
	ForgeVersion string    `yaml:"forge_version"`
	ProductName  string    `yaml:"product_name"`
	Apps         []App     `yaml:"apps"`
	Features     []Feature `yaml:"features"`
}

func (fs Product) Validate() error {
	return validation.ValidateStruct(&fs,
		validation.Field(&fs.ProductName, validation.Required),
		validation.Field(&fs.Apps, validation.Required),
	)
}

func UnmarshalYaml(data []byte) (fs Product, err error) {
	err = yaml.Unmarshal(data, &fs)
	if err != nil {
		return
	}

	err = fs.Validate()
	return
}
