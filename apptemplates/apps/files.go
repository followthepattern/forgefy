package apps

import (
	"path"

	"github.com/followthepattern/forgefy/apptemplates"
)

func Directory(productName string) string {
	return path.Join(apptemplates.RootDirectory(productName), "apps")
}
