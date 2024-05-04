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

func (fw FileWriter) Write(folderName string, fileName string, writerFn func(io.Writer) error) error {
	destDir := path.Join(fw.outputFolder, folderName)

	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		slog.Error(err.Error())
	}

	file, err := os.Create(path.Join(destDir, fileName))
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
