package apps

import (
	"path"

	"github.com/followthepattern/forgefy/apptemplates"
)

func Directory() string {
	return path.Join(apptemplates.RootDirectory(), "apps")
}
