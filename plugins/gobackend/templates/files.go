package templates

import (
	"embed"
	"path"

	"github.com/followthepattern/forgefy/plugins/monorepo/templates/apps"
)

func Directory(appName string) string {
	return path.Join(apps.Directory(), appName)
}

var (
	//go:embed * apps/\[(application)\]/policies/.cerbos.yaml.tmpl
	EntireDir embed.FS
)
