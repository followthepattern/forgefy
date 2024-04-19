package tests

import (
	"path"

	"github.com/followthepattern/forgefy/apptemplates/apps/backend"
)

func Directory(productName, appName string) string {
	return path.Join(backend.Directory(productName, appName), "tests")
}
