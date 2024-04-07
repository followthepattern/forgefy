package apptemplates

import (
	_ "embed"
)

//go:embed docker-compose.yaml.tmpl
var dockerComposeContent string

var DockerCompose = TemplateFile{
	Name:     "docker-compose.yaml",
	Template: dockerComposeContent,
}
