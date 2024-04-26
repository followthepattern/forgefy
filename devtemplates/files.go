package devtemplates

import (
	_ "embed"

	"github.com/followthepattern/forgefy/productmap"
)

var (
	//go:embed docker-compose.yaml.tmpl
	dockerComposeContent string
)

func RootDirectory() string {
	return ""
}

var DockerCompose = productmap.NewFile("docker-compose.yaml", dockerComposeContent)
