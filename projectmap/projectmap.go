package projectmap

type ProjectMap struct {
	Directory
	plugins map[string]Plugin
}

func NewProjectMap(projectName string) ProjectMap {
	return ProjectMap{
		Directory: NewDirectory(projectName),
	}
}

type Plugin interface {
	Name() string
}

func (p *ProjectMap) WithPlugin(plugin Plugin) *ProjectMap {
	p.plugins[plugin.Name()] = plugin
	return p
}
