package io

import (
	"log/slog"
	"os"
	"path"
)

type FileWriter struct {
	outputFolder string
}

func NewFileWriter(outputFolder string) FileWriter {
	return FileWriter{
		outputFolder: outputFolder,
	}
}

func (fw FileWriter) Write(fileName string, content string) error {
	err := os.MkdirAll(fw.outputFolder, 0755)
	if err != nil {
		slog.Error(err.Error())
	}

	file, err := os.Create(path.Join(fw.outputFolder, fileName))
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}
