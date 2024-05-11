package apptemplates

import (
	_ "embed"
	"path"

	"github.com/followthepattern/forgefy/plugins/monorepo/templates/apps"
	"github.com/followthepattern/forgefy/productmap"
)

func Directory(appName string) string {
	return path.Join(apps.Directory(), appName)
}

var (
	//go:embed package.json.tmpl
	packageJSON string
)

var PackageJSON = productmap.NewFile("package.json", packageJSON)
