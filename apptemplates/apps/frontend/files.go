package frontend

import (
	_ "embed"

	"github.com/followthepattern/forgefy/apptemplates"
)

//go:embed package.json.tmpl
var packageJSON string

var PackageJSON = apptemplates.TemplateFile{
	Name:     "package.json",
	Template: packageJSON,
}
