package apptemplates

import (
	_ "embed"
	"path"

	"github.com/followthepattern/forgefy/apptemplates/apps"
	"github.com/followthepattern/forgefy/productmap"
)

func Directory(appName string) string {
	return path.Join(apps.Directory(), appName)
}

var (
	//go:embed go.mod.tmpl
	goMod string
	//go:embed Dockerfile.tmpl
	dockerFile string
)

var GoMod = productmap.NewFile("go.mod", goMod)
var DockerFile = productmap.NewFile("Dockerfile", dockerFile)
