package feature

import (
	"path"

	"github.com/followthepattern/forgefy/apptemplates/apps/backend/features"
)

func Directory(productName, appName string) string {
	return path.Join(features.Directory(productName, appName))
}
