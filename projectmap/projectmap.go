package projectmap

import (
	"fmt"

	"github.com/followthepattern/forgefy/apptemplates"
	"github.com/followthepattern/forgefy/apptemplates/apps/frontend"
)

type ProjectMap struct {
	folderName string
	files      map[string]File
	folders    map[string]Directory
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

	frontend := Directory{
		directoryName: "frontend",
		files: map[string]File{
			packageJson.fileName: packageJson,
		},
	}

	apps := Directory{
		directoryName: "app",
		directories: map[string]Directory{
			"frontend": frontend,
		},
	}

	frontend.parentFolder = &apps

	return ProjectMap{
		files: map[string]File{
			dockerCompose.fileName: dockerCompose,
		},
		folders: map[string]Directory{
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
