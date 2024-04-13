package frontend

import (
	_ "embed"
	"path"

	"github.com/followthepattern/forgefy/apptemplates/apps"
	"github.com/followthepattern/forgefy/productmap"
)

func Directory(productName, appName string) string {
	return path.Join(apps.Directory(productName), appName)
}

var (
	//go:embed package.json.tmpl
	packageJSON string
)

var PackageJSON = productmap.NewFile("package.json", packageJSON)
