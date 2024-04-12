package main

import (
	"log/slog"
	"os"

	"github.com/followthepattern/forgefy/featureset"
	"github.com/followthepattern/forgefy/io"
	"github.com/followthepattern/forgefy/productmap"
)

func main() {
	yaml := `
version: 0
product_name: "test product"
apps:
  - name: backend1
    type: backend
  - name: frontend1
    type: frontend
`
	fs, err := featureset.UnmarshalYaml([]byte(yaml))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	builder := productmap.NewBuilder(fs)

	product, err := builder.Build()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	writer := io.NewFileWriter("output")

	product.Walk(func(folderName string, f productmap.File) {
		err := writer.Write(folderName, f.FileName(), f.Template())
		if err != nil {
			slog.Error(err.Error())
		}
	})
}
