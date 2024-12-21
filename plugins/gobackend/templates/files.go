package templates

import (
	"embed"
)

var (
	//go:embed * apps/\[(application)\]/policies/.cerbos.yaml.tmpl
	EntireDir embed.FS
)
