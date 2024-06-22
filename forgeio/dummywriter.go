package forgeio

import (
	"io"
)

type DummyFile struct {
	filePath string
}

var _ Writer = &DummyWriter{}

type DummyWriter struct {
	files map[string]DummyFile
}

// IOWriter implements Writer.
func (d *DummyWriter) Write(filePath string, _ func(io.Writer) error) error {
	d.files[filePath] = DummyFile{
		filePath: filePath,
	}
	return nil
}

func NewDummyWriter() DummyWriter {
	return DummyWriter{
		files: make(map[string]DummyFile, 0),
	}
}
