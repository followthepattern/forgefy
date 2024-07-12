package forgefy

import (
	"fmt"

	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/productmap"
	"github.com/followthepattern/forgefy/specification"
)

type Builder struct {
	productSpecification specification.Product
	plugins              []plugins.Plugin
	apps                 map[string][]plugins.App
}

func newBuilder(productSpecification specification.Product) Builder {
	return Builder{
		productSpecification: productSpecification,
		plugins:              make([]plugins.Plugin, 0),
		apps:                 make(map[string][]plugins.App),
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

	for _, plugin := range builder.plugins {
		err = plugin.Build(pm, builder.productSpecification)
		if err != nil {
			return pm, err
		}
	}

	for _, appSpecification := range builder.productSpecification.Apps {
		appSpecification.Product = builder.productSpecification
		appSpecification.Features = append(builder.productSpecification.Features, appSpecification.Features...)
		appSpecification.Init()
		err = builder.buildApps(pm, appSpecification)
		if err != nil {
			return pm, err
		}
	}

	return pm, nil
}

func (builder Builder) addDefaultFiles(_ productmap.ProductMap) error { return nil }

func (b Builder) buildApps(pm productmap.ProductMap, app specification.App) error {
	apps, ok := b.apps[string(app.AppType)]
	if !ok {
		return fmt.Errorf("unknown app definition: %s", app.AppType)
	}

	for _, a := range apps {
		err := a.Build(pm, app)

		if err != nil {
			return err
		}
	}

	return nil
}
