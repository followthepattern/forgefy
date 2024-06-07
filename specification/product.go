package specification

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"gopkg.in/yaml.v3"
)

const (
	defaultEmail = "app@product.io"
)

type Product struct {
	ForgeVersion string    `yaml:"forge_version"`
	ProductName  string    `yaml:"product_name"`
	Email        string    `yaml:"email"`
	Apps         []App     `yaml:"apps"`
	Features     []Feature `yaml:"features"`
}

func (fs Product) Validate() error {
	return validation.ValidateStruct(&fs,
		validation.Field(&fs.ProductName, validation.Required),
	)
}

func (p Product) ProductNameAsLower() string {
	return strings.ToLower(p.ProductName)
}

func UnmarshalYaml(data []byte) (p Product, err error) {
	err = yaml.Unmarshal(data, &p)
	if err != nil {
		return
	}

	p = setDefault(p)

	err = p.Validate()
	return
}

func setDefault(p Product) Product {
	if p.Email == "" {
		p.Email = defaultEmail
	}
	return p
}
