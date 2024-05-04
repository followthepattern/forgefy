package productmap

import (
	"errors"
	"io"
	"strings"
	"text/template"
)

type File struct {
	fileName string
	template string
	data     any
}

func NewFile(fileName string, template string) File {
	return File{
		fileName: fileName,
		template: template,
	}
}

func (f File) WithData(data any) File {
	f.data = data
	return f
}

func (f File) FileName() string {
	return f.fileName
}

func (f File) Template() string {
	return f.template
}

func (f File) Content() (string, error) {
	s := &strings.Builder{}

	err := f.Write(s)
	if err != nil {
		return "", err
	}

	return s.String(), nil
}

func (f File) Write(w io.Writer) error {
	if w == nil {
		return errors.New("the given writer is nil")
	}

	if f.data == nil {
		_, err := w.Write([]byte(f.template))
		return err
	}

	tmpl, err := template.New(f.fileName).Parse(f.template)
	if err != nil {
		return err
	}

	return tmpl.Execute(w, f.data)
}

func (f File) Data() any {
	return f.data
}
