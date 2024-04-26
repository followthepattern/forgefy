package apps

import (
	"path"

	"github.com/followthepattern/forgefy/devtemplates"
)

func Directory() string {
	return path.Join(devtemplates.RootDirectory(), "apps")
}
