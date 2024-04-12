package productmap

import (
	"strings"
	"text/template"

	"github.com/followthepattern/forgefy/apptemplates"
)

type File struct {
	fileName string
	template string
	data     any
}

func NewFile(fileName string, template string, data any) File {
	return File{
		fileName: fileName,
		template: template,
		data:     data,
	}
}

func NewFileFromTemplate(tpl apptemplates.TemplateFile) File {
	return File{
		fileName: tpl.Name,
		template: tpl.Template,
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
	if f.data == nil {
		return f.template, nil
	}

	tmpl, err := template.New(f.fileName).Parse(f.template)
	if err != nil {
		return "", err
	}
	s := &strings.Builder{}

	tmpl.Execute(s, f.data)

	return s.String(), nil
}

func (f File) Data() any {
	return f.data
}
