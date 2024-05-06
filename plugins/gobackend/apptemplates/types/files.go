package types

import (
	"path"

	"github.com/followthepattern/forgefy/plugins/gobackend/apptemplates"
	"github.com/followthepattern/forgefy/productmap"

	_ "embed"
)

func Directory(appName string) string {
	return path.Join(apptemplates.Directory(appName), "types")
}

var (
	//go:embed bool.go.tmpl
	bool string
	//go:embed bytes.go.tmpl
	bytes string
	//go:embed convert.go.tmpl
	convert string
	//go:embed int.go.tmpl
	intFile string
	//go:embed int64.go.tmpl
	int64File string
	//go:embed string.go.tmpl
	stringFile string
	//go:embed time.go.tmpl
	timeFile string
	//go:embed type.go.tmpl
	typeFile string
	//go:embed uint.go.tmpl
	uintFile string
	//go:embed validation.go.tmpl
	validation string
)

var Files = []productmap.File{
	productmap.NewFile("bool.go", bool),
	productmap.NewFile("bytes.go", bytes),
	productmap.NewFile("convert.go", convert),
	productmap.NewFile("int.go", intFile),
	productmap.NewFile("int64.go", int64File),
	productmap.NewFile("string.go", stringFile),
	productmap.NewFile("time.go", timeFile),
	productmap.NewFile("type.go", typeFile),
	productmap.NewFile("uint.go", uintFile),
	productmap.NewFile("validation.go", validation),
}
