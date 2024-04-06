package projectmap

import (
	"path"
)

type Directory struct {
	parentDirectory *Directory
	directoryName   string
	files           map[string]File
	directories     map[string]Directory
}

func (dir Directory) DirName() string {
	if dir.parentDirectory == nil {
		return dir.directoryName
	}
	return path.Join(dir.parentDirectory.DirName(), dir.directoryName)
}

func (dir Directory) Walk(fn func(directoryName string, f File)) {
	for _, file := range dir.files {
		fn(dir.DirName(), file)
	}

	for _, dir := range dir.directories {
		dir.Walk(fn)
	}
}
