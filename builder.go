package forgefy

import (
	"fmt"

	"github.com/followthepattern/forgefy/apptemplates"
	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/productmap"
	"github.com/followthepattern/forgefy/specification"
)

type Builder struct {
	fs      specification.Product
	plugins []plugins.Plugin
	apps    map[string][]plugins.App
}

func newBuilder(fs specification.Product) Builder {
	return Builder{
		fs:      fs,
		plugins: make([]plugins.Plugin, 0),
		apps:    make(map[string][]plugins.App),
	}
}

func (builder Builder) withPlugins(plugins ...plugins.Plugin) Builder {
	builder.plugins = plugins
	return builder
}

func (builder Builder) withApps(apps map[string][]plugins.App) Builder {
	builder.apps = apps
	return builder
}

func (builder Builder) Build(plugins ...plugins.Plugin) (productmap.ProductMap, error) {
	pm := productmap.NewProductMap()

	err := builder.addDefaultFiles(pm)
	if err != nil {
		return pm, err
	}

	err = builder.addProjectLevelFiles(pm)
	if err != nil {
		return pm, err
	}

	for _, plugin := range builder.plugins {
		err = plugin.AddFiles(pm, builder.fs)
		if err != nil {
			return pm, err
		}
	}

	for _, app := range builder.fs.Apps {
		err = builder.addAppSpecificFiles(pm, builder.fs, app)
		if err != nil {
			return pm, err
		}
	}

	return pm, nil
}

func (builder Builder) addDefaultFiles(_ productmap.ProductMap) error { return nil }

func (builder Builder) addProjectLevelFiles(pm productmap.ProductMap) error {
	dir := apptemplates.RootDirectory()
	return pm.Insert(dir, apptemplates.DockerCompose)
}

func (b Builder) addAppSpecificFiles(pm productmap.ProductMap, fs specification.Product, app specification.App) error {
	apps, ok := b.apps[string(app.AppType)]
	if !ok {
		return fmt.Errorf("unknown app definition: %s", app.AppType)
	}

	for _, a := range apps {
		err := a.Builder(pm, fs, app)

		if err != nil {
			return err
		}
	}

	return nil
}
