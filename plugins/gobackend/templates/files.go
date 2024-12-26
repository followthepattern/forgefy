package templates

import (
	"embed"
)

var (
	//go:embed * apps/\[(application)\]/policies/.cerbos.yaml.tmpl apps/\[(application)\]/.gitignore.tmpl
	EntireDir embed.FS
)
