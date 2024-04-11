package apptemplates

import (
	_ "embed"
)

var (
	//go:embed docker-compose.yaml.tmpl
	dockerComposeContent string
)

func RootDirectory(dirName string) string {
	return dirName
}

var DockerCompose = TemplateFile{
	Name:     "docker-compose.yaml",
	Template: dockerComposeContent,
}
