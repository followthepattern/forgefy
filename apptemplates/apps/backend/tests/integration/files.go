package integration

import (
	"path"

	"github.com/followthepattern/forgefy/apptemplates/apps/backend/tests"
)

func Directory(productName, appName string) string {
	return path.Join(tests.Directory(productName, appName), "integration")
}
