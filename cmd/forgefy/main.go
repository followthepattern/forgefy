package main

import (
	"log/slog"

	"github.com/followthepattern/forgefy/io"
	"github.com/followthepattern/forgefy/projectmap"
)

func main() {
	project := projectmap.NewProjectMap()

	writer := io.NewFileWriter("output")

	project.Walk(func(f projectmap.File) {
		err := writer.Write(f.FileName(), f.Template())
		if err != nil {
			slog.Error(err.Error())
		}
	})
}
