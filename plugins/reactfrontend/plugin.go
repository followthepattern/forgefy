package reactfrontend

import (
	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/productmap"
	"github.com/followthepattern/forgefy/specification"
)

var _ plugins.Plugin = Plugin{}

type Plugin struct{}

func (Plugin) Name() string { return "React Plugin" }

func (Plugin) Apps() []plugins.App {
	return []plugins.App{
		ReactFrontend{},
	}
}

func (Plugin) AddFiles(_ productmap.ProductMap, _ specification.Product) error { return nil }
