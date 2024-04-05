package projectmap

type ProjectMap struct {
	folderName string
	files      map[string]File
	folders    map[string]FolderMap
	plugins    map[string]Plugin
}

type Plugin interface {
	Name() string
}

func (p *ProjectMap) WithPlugin(plugin Plugin) *ProjectMap {
	p.plugins[plugin.Name()] = plugin
	return p
}
