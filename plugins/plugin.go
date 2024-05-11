package plugins

import (
	"github.com/followthepattern/forgefy/productmap"
	"github.com/followthepattern/forgefy/specification"
)

type Plugin interface {
	Name() string
	Build(productmap.ProductMap, specification.Product) error
	Apps() []App
}

type App interface {
	Name() string
	Type() string
	Build(productmap.ProductMap, specification.Product, specification.App) error
}
