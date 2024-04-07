package backend

import (
	_ "embed"

	"github.com/followthepattern/forgefy/apptemplates"
)

//go:embed go.mod.tmpl
var goMod string

var GoMod = apptemplates.TemplateFile{
	Name:     "go.mod",
	Template: goMod,
}
