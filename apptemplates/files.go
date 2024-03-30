package apptemplates

import (
	_ "embed"
)

//go:embed docker-compose.yaml.tmpl
var DockerCompose string
