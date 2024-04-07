package main

import (
	"log/slog"

	"github.com/followthepattern/forgefy/featureset"
	"github.com/followthepattern/forgefy/io"
	"github.com/followthepattern/forgefy/projectmap"
)

func main() {
	fs := featureset.FeatureSet{
		ProjectName: "project",
	}

	builder := projectmap.NewBuilder(fs)

	project := builder.Create()

	writer := io.NewFileWriter("output")

	project.Walk(func(folderName string, f projectmap.File) {
		err := writer.Write(folderName, f.FileName(), f.Template())
		if err != nil {
			slog.Error(err.Error())
		}
	})
}
