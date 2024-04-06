package projectmap

import (
	"path"
)

type Directory struct {
	parentFolder  *Directory
	directoryName string
	files         map[string]File
	directories   map[string]Directory
}

func (dir Directory) DirName() string {
	if dir.parentFolder == nil {
		return dir.directoryName
	}
	return path.Join(dir.parentFolder.DirName(), dir.directoryName)
}

func (dir Directory) Walk(fn func(directoryName string, f File)) {
	for _, file := range dir.files {
		fn(dir.DirName(), file)
	}

	for _, dir := range dir.directories {
		dir.Walk(fn)
	}
}
