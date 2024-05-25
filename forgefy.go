package forgefy

import (
	"github.com/followthepattern/forgefy/forgeio"
	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/productmap"
	"github.com/followthepattern/forgefy/specification"
)

type Forgefy struct {
	plugins []plugins.Plugin
	apps    map[string][]plugins.App
}

func New() Forgefy {
	return Forgefy{
		apps: make(map[string][]plugins.App),
	}
}

func (f Forgefy) Forge(yaml string, fw forgeio.Writer) (string, error) {
	productSpecification, err := specification.UnmarshalYaml([]byte(yaml))
	if err != nil {
		return "", err
	}

	builder := newBuilder(productSpecification).withPlugins(f.plugins...).withApps(f.apps)

	product, err := builder.Build()
	if err != nil {
		return productSpecification.ProductName, err
	}

	err = product.Walk(func(folderName string, file productmap.File) error {
		return fw.Write(folderName, file.FileName(), file.Write)
	})

	return productSpecification.ProductName, err
}

func (f Forgefy) verifyPlugins(_ []plugins.Plugin) error { return nil }

func (f *Forgefy) InstallPlugins(plugins ...plugins.Plugin) error {
	if len(plugins) < 1 {
		return nil
	}

	err := f.verifyPlugins(plugins)

	if err != nil {
		return err
	}

	f.plugins = plugins

	for _, plugin := range plugins {
		for _, appBuilder := range plugin.Apps() {
			appBuilders := f.apps[appBuilder.Type()]
			f.apps[appBuilder.Type()] = append(appBuilders, appBuilder)
		}
	}
	return nil
}
