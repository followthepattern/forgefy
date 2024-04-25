package integration

import (
	"path"

	"github.com/followthepattern/forgefy/plugins/gobackend/apptemplates/tests"
)

func Directory(appName string) string {
	return path.Join(tests.Directory(appName), "integration")
}
