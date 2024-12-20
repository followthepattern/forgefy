package templates

import (
	"embed"
)

var (
	//go:embed * dagger_dagger_/.gitattributes.tmpl dagger_dagger_/.gitignore.tmpl
	Files embed.FS
)

func RootDirectory() string {
	return ""
}
