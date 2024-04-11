package frontend

import (
	_ "embed"
	"path"

	"github.com/followthepattern/forgefy/apptemplates"
	"github.com/followthepattern/forgefy/apptemplates/apps"
)

func Directory(productName, appName string) string {
	return path.Join(apps.Directory(productName), appName)
}

var (
	//go:embed package.json.tmpl
	packageJSON string
)

var PackageJSON = apptemplates.TemplateFile{
	Name:     "package.json",
	Template: packageJSON,
}
