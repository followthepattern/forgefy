package projectmap

import "github.com/followthepattern/forgefy/apptemplates"

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

func (f File) FileName() string {
	return f.fileName
}

func (f File) Template() string {
	return f.template
}

func (f File) Data() any {
	return f.data
}
