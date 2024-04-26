package plugins

import (
	"github.com/followthepattern/forgefy/featureset"
	"github.com/followthepattern/forgefy/productmap"
)

type Plugin interface {
	Name() string
	AddFiles(productmap.ProductMap, featureset.FeatureSet) error
	Apps() []App
}

type App interface {
	Name() string
	Type() string
	Builder(productmap.ProductMap, featureset.FeatureSet, featureset.App) error
}
