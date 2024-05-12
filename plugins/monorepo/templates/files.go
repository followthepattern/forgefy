package templates

import (
	"embed"
)

var (
	//go:embed *
	Files embed.FS
)

func RootDirectory() string {
	return ""
}
