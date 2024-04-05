package projectmap

import "github.com/followthepattern/forgefy/apptemplates"

type ProjectMap struct {
	folderName string
	files      map[string]File
	folders    map[string]FolderMap
	plugins    map[string]Plugin
}

func NewProjectMap() ProjectMap {
	dockerCompose := File{
		fileName: "docker-compose.yaml",
		template: apptemplates.DockerCompose,
	}

	return ProjectMap{
		files: map[string]File{
			dockerCompose.fileName: dockerCompose,
		},
	}
}

type Plugin interface {
	Name() string
}

func (p *ProjectMap) WithPlugin(plugin Plugin) *ProjectMap {
	p.plugins[plugin.Name()] = plugin
	return p
}

func (p ProjectMap) Walk(fn func(f File)) {
	for _, file := range p.files {
		fn(file)
	}
}
