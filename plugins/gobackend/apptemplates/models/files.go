package models

import (
	"path"

	"github.com/followthepattern/forgefy/plugins/gobackend/apptemplates"
	"github.com/followthepattern/forgefy/productmap"

	_ "embed"
)

func Directory(appName string) string {
	return path.Join(apptemplates.Directory(appName), "models")
}

var (
	//go:embed list.go.tmpl
	list string
	//go:embed response_status.go.tmpl
	responseStatus string
	//go:embed userlog.go.tmpl
	userLog string
)

var Files = []productmap.File{
	productmap.NewFile("list.go", list),
	productmap.NewFile("response_status.go", responseStatus),
	productmap.NewFile("userlog.go", userLog),
}
