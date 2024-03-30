package backend

import (
	_ "embed"
)

//go:embed go.mod.tmpl
var GoMod string
