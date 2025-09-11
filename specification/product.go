package specification

import (
	"fmt"
	"strings"

	"github.com/followthepattern/forgefy/specification/naming"
	validation "github.com/go-ozzo/ozzo-validation"
	"gopkg.in/yaml.v3"
)

type Product struct {
	ForgeVersion  string        `yaml:"forge_version"`
	Name          string        `yaml:"product_name"`
	Email         string        `yaml:"email"`
	Apps          []App         `yaml:"apps"`
	Features      []Feature     `yaml:"features"`
	ExcludeDagger bool          `yaml:"exclude_dagger"`
	Monitoring    bool          `yaml:"monitoring"`
	VSCode        bool          `yaml:"vscode"`
	PluginsConfig PluginsConfig `yaml:"plugins"`
}

type PluginsConfig []map[string]any

func (pc PluginsConfig) GetConfig(pluginName string) map[string]any {
	for _, config := range pc {
		if _, ok := config[pluginName]; ok {
			return config
		}
	}
	return nil
}

func (fs Product) Validate() error {
	return validation.ValidateStruct(&fs,
		validation.Field(&fs.Name, validation.Required),
	)
}

func (product *Product) setDefaults() {
	if product.Email == "" {
		product.Email = fmt.Sprintf("info@%s.com", strings.ToLower(product.Name))
	}
}

func (p Product) ProductNameCamelCase() string {
	return naming.ToLowerCamelCase(p.Name)
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
	commonFeatures := p.Features
	for i := 0; i < len(p.Apps); i++ {
		p.Apps[i].Product = p
		p.Apps[i].setDefaults()

		appSpecificFeatures := p.Apps[i].Features

		sum := append(commonFeatures, appSpecificFeatures...)
		features := make([]Feature, len(sum))

		for index, feature := range sum {
			feature.App = p.Apps[i]
			features[index] = feature
		}

		p.Apps[i].Features = features
	}

	return p
}
