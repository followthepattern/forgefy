package apps

import (
	"path"

	"github.com/followthepattern/forgefy/plugins/monorepo/templates"
)

func Directory() string {
	return path.Join(templates.RootDirectory(), "apps")
}
