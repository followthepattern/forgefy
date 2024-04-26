package reactfrontend

import (
	"github.com/followthepattern/forgefy/featureset"
	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/productmap"
)

var _ plugins.Plugin = Plugin{}

type Plugin struct{}

func (Plugin) Name() string { return "Go Plugin" }

func (Plugin) Apps() []plugins.App {
	return []plugins.App{
		ReactFrontend{},
	}
}

func (Plugin) AddFiles(_ productmap.ProductMap, _ featureset.FeatureSet) error { return nil }
