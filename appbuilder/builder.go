package appbuilder

import (
	"fmt"

	"github.com/followthepattern/forgefy/apptemplates"
	"github.com/followthepattern/forgefy/featureset"
	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/productmap"
)

type Builder struct {
	fs      featureset.FeatureSet
	plugins []plugins.Plugin
	apps    map[string][]plugins.App
}

func NewBuilder(fs featureset.FeatureSet) Builder {
	return Builder{
		fs:      fs,
		plugins: make([]plugins.Plugin, 0),
		apps:    make(map[string][]plugins.App),
	}
}

func (builder Builder) verifyPlugins(_ []plugins.Plugin) error { return nil }

func (builder *Builder) setPlugins(plugins []plugins.Plugin) error {
	err := builder.verifyPlugins(plugins)

	if err != nil {
		return err
	}

	builder.plugins = plugins

	for _, plugin := range plugins {
		for _, appBuilder := range plugin.Apps() {
			appBuilders := builder.apps[appBuilder.Type()]
			builder.apps[appBuilder.Type()] = append(appBuilders, appBuilder)
		}
	}
	return nil
}

func (builder Builder) Build(plugins ...plugins.Plugin) (productmap.ProductMap, error) {
	fs := builder.fs

	err := builder.setPlugins(plugins)
	if err != nil {
		return productmap.ProductMap{}, nil
	}

	pm := productmap.NewProductMap()

	err = builder.addDefaultFiles(pm)
	if err != nil {
		return pm, err
	}

	err = builder.addProjectLevelFiles(pm)
	if err != nil {
		return pm, err
	}

	for _, plugin := range builder.plugins {
		err = plugin.AddFiles(pm, fs)
		if err != nil {
			return pm, err
		}
	}

	for _, app := range fs.Apps {
		err = builder.addAppSpecificFiles(pm, fs, app)
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

func (b Builder) addAppSpecificFiles(pm productmap.ProductMap, fs featureset.FeatureSet, app featureset.App) error {
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
