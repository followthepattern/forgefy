package forgeio

import (
	"io"
	"log/slog"
	"os"
	"path"
)

var _ Writer = FileWriter{}

type FileWriter struct {
	outputFolder string
}

func (fw FileWriter) Write(filePath string, writerFn func(io.Writer) error) error {
	dir, _ := path.Split(filePath)
	destDir := path.Join(fw.outputFolder, dir)

	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		slog.Error(err.Error())
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	return writerFn(file)
}

func NewFileWriter(outputFolder string) FileWriter {
	return FileWriter{
		outputFolder: outputFolder,
	}
}
