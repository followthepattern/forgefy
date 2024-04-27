package forgeio

import (
	"path"
)

type DummyFile struct {
	dirName  string
	fileName string
	content  string
}

var _ Writer = &DummyWriter{}

type DummyWriter struct {
	files map[string]DummyFile
}

func NewDummyWriter() DummyWriter {
	return DummyWriter{
		files: make(map[string]DummyFile, 0),
	}
}

func (d *DummyWriter) Write(dirName, fileName, content string) error {
	filePath := path.Join(dirName, fileName)
	d.files[filePath] = DummyFile{
		dirName:  dirName,
		fileName: fileName,
		content:  content,
	}
	return nil
}
