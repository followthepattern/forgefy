package productmap

type ProductMap struct {
	Directory
	plugins map[string]Plugin
}

func NewProductMap(productName string) ProductMap {
	return ProductMap{
		Directory: NewDirectory(""),
	}
}

type Plugin interface {
	Name() string
}

func (p *ProductMap) WithPlugin(plugin Plugin) *ProductMap {
	p.plugins[plugin.Name()] = plugin
	return p
}
