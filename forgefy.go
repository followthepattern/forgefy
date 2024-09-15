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

type ForgeConfging struct {
	exclude string
}

type ForgeConfigOption interface {
	apply(f *ForgeConfging)
}

type ForgeConfigOptionFn func(f *ForgeConfging)

func (fn ForgeConfigOptionFn) apply(f *ForgeConfging) {
	fn(f)
}

func WithExclude(exclude string) ForgeConfigOption {
	return ForgeConfigOptionFn(func(f *ForgeConfging) {
		f.exclude = exclude
	})
}

func (f Forgefy) Forge(yaml string, fw forgeio.Writer, opts ...ForgeConfigOption) (string, error) {
	forgeConfig := &ForgeConfging{}

	for _, opt := range opts {
		opt.apply(forgeConfig)
	}

	productSpecification, err := specification.UnmarshalYaml([]byte(yaml))
	if err != nil {
		return "", err
	}

	builder := newBuilder(productSpecification).withPlugins(f.plugins...).withApps(f.apps)

	product, err := builder.Build()
	if err != nil {
		return productSpecification.ProductName, err
	}

	if len(forgeConfig.exclude) != 0 {
		product.Exclude(forgeConfig.exclude)
	}

	err = product.Walk(func(filePath string, file productmap.File) error {
		return fw.Write(filePath, file.Write)
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
