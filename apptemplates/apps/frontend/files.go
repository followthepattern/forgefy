package frontend

import (
	_ "embed"
)

//go:embed package.json.tmpl
var PackageJSON string
