package features

import (
	"path"

	"github.com/followthepattern/forgefy/plugins/gobackend/apptemplates"
)

func Directory(appName string) string {
	return path.Join(apptemplates.Directory(appName), "features")
}
