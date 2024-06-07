package gobackend

import (
	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/productmap"
	"github.com/followthepattern/forgefy/specification"
)

var _ plugins.Plugin = Plugin{}

type Plugin struct{}

func (Plugin) Name() string { return "Go Plugin" }

func (Plugin) Apps() []plugins.App {
	return []plugins.App{
		&GoBackendPluginApp{
			port:       8080,
			cerbosPort: 3592,
			dbPort:     5433,
		},
	}
}

func (Plugin) Build(_ productmap.ProductMap, _ specification.Product) error { return nil }
