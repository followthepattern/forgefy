package integration

import (
	"path"

	"github.com/followthepattern/forgefy/apptemplates/apps/backend/tests"
)

func Directory(appName string) string {
	return path.Join(tests.Directory(appName), "integration")
}
