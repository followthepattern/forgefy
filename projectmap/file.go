package projectmap

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

func (f File) FileName() string {
	return f.fileName
}

func (f File) Template() string {
	return f.template
}

func (f File) Data() any {
	return f.data
}
