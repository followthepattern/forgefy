package apptemplates

import (
	_ "embed"

	"github.com/followthepattern/forgefy/productmap"
)

var (
	//go:embed docker-compose.yaml.tmpl
	dockerComposeContent string
)

func RootDirectory(dirName string) string {
	return dirName
}

var DockerCompose = productmap.NewFile("docker-compose.yaml", dockerComposeContent)
