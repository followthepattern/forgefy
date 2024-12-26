package templates

import (
	"embed"
	_ "embed"
)

var (
	//go:embed * apps/\[(application)\]/.env.development.tmpl apps/\[(application)\]/.env.production.tmpl apps/\[(application)\]/.gitignore.tmpl
	Files embed.FS
)
