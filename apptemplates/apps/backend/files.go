package backend

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
	//go:embed go.mod.tmpl
	goMod string
)

var GoMod = productmap.NewFile("go.mod", goMod)
