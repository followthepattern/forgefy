package forgeio

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

func (fw FileWriter) Write(folderName string, fileName string, content string) error {
	destDir := path.Join(fw.outputFolder, folderName)

	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		slog.Error(err.Error())
	}

	file, err := os.Create(path.Join(destDir, fileName))
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}
