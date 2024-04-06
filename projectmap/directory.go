package projectmap

import (
	"fmt"
	"path"
	"strings"
)

type Directory struct {
	parentDirectory *Directory
	directoryName   string
	files           map[string]File
	directories     map[string]Directory
}

func NewDirectory(directoryName string) Directory {
	return Directory{
		directoryName: directoryName,
		files:         make(map[string]File),
		directories:   make(map[string]Directory),
	}
}

func (dir Directory) Exists(filePath string) bool {
	dirPath, fileName := path.Split(filePath)
	return dir.exists(dirPath, fileName)
}

func (dir Directory) exists(dirPath string, fileName string) bool {
	_, ok := dir.files[fileName]
	if ok {
		return true
	}

	if len(dirPath) == 0 {
		return false
	}

	dirs := strings.Split(dirPath, "/")
	d, ok := dir.directories[dirs[0]]
	if !ok {
		return false
	}

	return d.exists(strings.Join(dirs[1:], "/"), fileName)
}

func (dir Directory) DirName() string {
	if dir.parentDirectory == nil {
		return dir.directoryName
	}
	return path.Join(dir.parentDirectory.DirName(), dir.directoryName)
}

func (dir *Directory) AddFile(file File) error {
	_, ok := dir.files[file.FileName()]

	if ok {
		return fmt.Errorf("file %s already exists at %s", file.FileName(), dir.DirName())
	}

	dir.files[file.FileName()] = file

	return nil
}

func (dir *Directory) AddDirectory(directory Directory) error {
	_, ok := dir.directories[directory.directoryName]
	if ok {
		return fmt.Errorf("directory %s already exists at %s", dir.directoryName, dir.DirName())
	}

	dir.directories[directory.directoryName] = directory

	return nil
}

func (dir *Directory) Insert(directoryPath string, file File) error {
	return dir.insert(directoryPath, file)
}

func (dir *Directory) insert(directoryPath string, file File) error {
	if exists := dir.exists(directoryPath, file.FileName()); exists {
		return fmt.Errorf("file %s already exists at %s", file.FileName(), dir.DirName())
	}

	if len(directoryPath) == 0 || directoryPath == dir.directoryName {
		dir.files[file.FileName()] = file
		return nil
	}

	d, ok := dir.directories[directoryPath]
	if ok {
		d.AddFile(file)
	}

	dirs := strings.Split(directoryPath, "/")
	newDir := NewDirectory(dirs[0])

	directoryPath = strings.Join(dirs[1:], "/")
	dir.AddDirectory(newDir)

	return newDir.insert(directoryPath, file)
}

func (dir Directory) Walk(fn func(directoryName string, f File)) {
	for _, file := range dir.files {
		fn(dir.DirName(), file)
	}

	for _, dir := range dir.directories {
		dir.Walk(fn)
	}
}
