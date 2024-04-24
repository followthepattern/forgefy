package features

import (
	"path"

	"github.com/followthepattern/forgefy/apptemplates/apps/backend"
)

func Directory(appName string) string {
	return path.Join(backend.Directory(appName), "features")
}
