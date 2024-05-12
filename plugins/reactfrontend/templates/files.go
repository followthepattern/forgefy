package templates

import (
	"embed"
	_ "embed"
)

var (
	//go:embed *
	Files embed.FS
)
