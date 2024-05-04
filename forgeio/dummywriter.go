package forgeio

import (
	"io"
	"path"
)

type DummyFile struct {
	dirName  string
	fileName string
}

var _ Writer = &DummyWriter{}

type DummyWriter struct {
	files map[string]DummyFile
}

// IOWriter implements Writer.
func (d *DummyWriter) Write(dirName string, fileName string, _ func(io.Writer) error) error {
	filePath := path.Join(dirName, fileName)
	d.files[filePath] = DummyFile{
		dirName:  dirName,
		fileName: fileName,
	}
	return nil
}

func NewDummyWriter() DummyWriter {
	return DummyWriter{
		files: make(map[string]DummyFile, 0),
	}
}
