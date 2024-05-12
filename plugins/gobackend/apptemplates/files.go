package apptemplates

import (
	"embed"
	"path"

	"github.com/followthepattern/forgefy/plugins/monorepo/templates/apps"
)

func Directory(appName string) string {
	return path.Join(apps.Directory(), appName)
}

var (
	//go:embed *
	EntireDir embed.FS
)
