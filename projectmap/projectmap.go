package projectmap

import (
	"fmt"

	"github.com/followthepattern/forgefy/apptemplates"
	"github.com/followthepattern/forgefy/apptemplates/apps/frontend"
)

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

	packageJson := File{
		fileName: "package.json",
		template: frontend.PackageJSON,
	}

	frontend := FolderMap{
		folderName: "frontend",
		files: map[string]File{
			packageJson.fileName: packageJson,
		},
	}

	apps := FolderMap{
		folderName: "app",
		folders: map[string]FolderMap{
			"frontend": frontend,
		},
	}

	frontend.parentFolder = &apps

	return ProjectMap{
		files: map[string]File{
			dockerCompose.fileName: dockerCompose,
		},
		folders: map[string]FolderMap{
			"apps": apps,
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

func (p ProjectMap) Walk(fn func(folderName string, f File)) {
	for _, file := range p.files {
		fn(p.folderName, file)
	}
	for _, dir := range p.folders {
		fmt.Println("starts walking folderMap")
		dir.Walk(fn)
	}
}
