package projectmap

import (
	"github.com/followthepattern/forgefy/apptemplates"
	"github.com/followthepattern/forgefy/apptemplates/apps/backend"
	"github.com/followthepattern/forgefy/apptemplates/apps/frontend"
)

type ProjectMap struct {
	Directory
	plugins map[string]Plugin
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

	goMod := File{
		fileName: "go.mod",
		template: backend.GoMod,
	}

	backend := Directory{
		directoryName: "backend",
		files: map[string]File{
			goMod.fileName: goMod,
		},
	}

	frontend := Directory{
		directoryName: "frontend",
		files: map[string]File{
			packageJson.fileName: packageJson,
		},
	}

	projectmap := ProjectMap{
		Directory: Directory{
			files: map[string]File{
				dockerCompose.fileName: dockerCompose,
			},
		},
	}

	apps := Directory{
		parentDirectory: &projectmap.Directory,
		directoryName:   "apps",
		directories:     map[string]Directory{},
	}

	backend.parentDirectory = &apps
	frontend.parentDirectory = &apps

	apps.directories[frontend.directoryName] = frontend
	apps.directories[backend.directoryName] = backend

	projectmap.directories = map[string]Directory{
		apps.directoryName: apps,
	}

	return projectmap
}

type Plugin interface {
	Name() string
}

func (p *ProjectMap) WithPlugin(plugin Plugin) *ProjectMap {
	p.plugins[plugin.Name()] = plugin
	return p
}
