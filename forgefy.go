package forgefy

import (
	"github.com/followthepattern/forgefy/featureset"
	"github.com/followthepattern/forgefy/io"
	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/productmap"
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

func (f Forgefy) Forge(yaml string, fw io.Writer) (string, error) {
	fs, err := featureset.UnmarshalYaml([]byte(yaml))
	if err != nil {
		return "", err
	}

	builder := NewBuilder(fs).withPlugins(f.plugins...).withApps(f.apps)

	product, err := builder.Build()
	if err != nil {
		return fs.ProductName, err
	}

	err = product.Walk(func(folderName string, file productmap.File) error {
		content, err := file.Content()
		if err != nil {
			return err
		}
		return fw.Write(folderName, file.FileName(), content)
	})

	return fs.ProductName, err
}

func (f Forgefy) verifyPlugins(_ []plugins.Plugin) error { return nil }

func (f *Forgefy) SetPlugins(plugins ...plugins.Plugin) error {
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
