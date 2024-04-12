package backend

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
	//go:embed go.mod.tmpl
	goMod string
)

var GoMod = apptemplates.TemplateFile{
	Name:     "go.mod",
	Template: goMod,
}
