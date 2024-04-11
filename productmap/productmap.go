package productmap

const rootDirectory = ""

type ProductMap struct {
	Directory
	plugins map[string]Plugin
}

func NewProductMap() ProductMap {
	return ProductMap{
		Directory: NewDirectory(rootDirectory),
	}
}

type Plugin interface {
	Name() string
}

func (p *ProductMap) WithPlugin(plugin Plugin) *ProductMap {
	p.plugins[plugin.Name()] = plugin
	return p
}
