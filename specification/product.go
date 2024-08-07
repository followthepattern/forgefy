package specification

import (
	"fmt"
	"strings"

	"github.com/followthepattern/forgefy/specification/naming"
	validation "github.com/go-ozzo/ozzo-validation"
	"gopkg.in/yaml.v3"
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

func (product *Product) setDefaults() {
	if product.Email == "" {
		product.Email = fmt.Sprintf("info@%s.com", strings.ToLower(product.ProductName))
	}
}

func (p Product) ProductNameCamelCase() string {
	return naming.ToLowerCamelCase(p.ProductName)
}

func UnmarshalYaml(data []byte) (p Product, err error) {
	err = yaml.Unmarshal(data, &p)
	if err != nil {
		return
	}

	err = p.Validate()

	p.setDefaults()

	p = constructProduct(p)

	return
}

func constructProduct(p Product) Product {
	for i := 0; i < len(p.Apps); i++ {
		p.Apps[i].Product = p
		p.Apps[i].Features = append(p.Features, p.Apps[i].Features...)

		p.Apps[i].setDefaults()

		for j := 0; j < len(p.Apps[i].Features); j++ {
			p.Apps[i].Features[j].App = p.Apps[i]
		}
	}

	return p
}
