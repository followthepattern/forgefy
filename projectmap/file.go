package projectmap

import (
	"log/slog"
)

type File struct {
	fileName string
	template string
	data     any
}

func (f File) Write() {
	slog.Info(f.fileName, f.template, f.data)
}
