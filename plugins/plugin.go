package plugins

import (
	"github.com/followthepattern/forgefy/productmap"
	"github.com/followthepattern/forgefy/specification"
)

type Plugin interface {
	Name() string
	AddFiles(productmap.ProductMap, specification.Product) error
	Apps() []App
}

type App interface {
	Name() string
	Type() string
	Builder(productmap.ProductMap, specification.Product, specification.App) error
}
