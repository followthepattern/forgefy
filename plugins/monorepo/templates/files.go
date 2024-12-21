package templates

import (
	"embed"
)

var (
	//go:embed * \.dagger\($dagger$)/.gitattributes.tmpl \.dagger\($dagger$)/.gitignore.tmpl
	Files embed.FS
)
