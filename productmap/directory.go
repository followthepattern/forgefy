package productmap

import (
	"fmt"
	"path"
	"strings"
)

type WalkFn func(directoryName string, f File) error

type directory struct {
	parentDirectory *directory
	directoryName   string
	files           map[string]File
	directories     map[string]directory
}

func newDirectory(directoryName string, parentDir *directory) directory {
	return directory{
		parentDirectory: parentDir,
		directoryName:   directoryName,
		files:           make(map[string]File),
		directories:     make(map[string]directory),
	}
}

func (dir directory) Exists(filePath string) bool {
	dirPath, fileName := path.Split(filePath)
	return dir.exists(dirPath, fileName)
}

func (dir directory) exists(dirPath string, fileName string) bool {
	if len(dirPath) == 0 {
		_, ok := dir.files[fileName]
		return ok
	}

	dirs := strings.Split(dirPath, "/")
	d, ok := dir.directories[dirs[0]]
	if !ok {
		return false
	}

	return d.exists(strings.Join(dirs[1:], "/"), fileName)
}

func (dir directory) DirName() string {
	if dir.parentDirectory == nil {
		return dir.directoryName
	}
	return path.Join(dir.parentDirectory.DirName(), dir.directoryName)
}

func (dir *directory) AddFile(file File) error {
	_, ok := dir.files[file.FileName()]

	if ok {
		return fmt.Errorf("file %s already exists at %s", file.FileName(), dir.DirName())
	}

	dir.files[file.FileName()] = file

	return nil
}

func (dir *directory) AddDirectory(directory directory) error {
	_, exists := dir.directories[directory.directoryName]
	if exists {
		return fmt.Errorf("directory %s already exists at %s", dir.directoryName, dir.DirName())
	}

	dir.directories[directory.directoryName] = directory

	return nil
}

func (dir *directory) AddChild(dirName string) (*directory, error) {
	_, exists := dir.directories[dirName]
	if exists {
		return nil, fmt.Errorf("directory %s already exists at %s", dir.directoryName, dir.DirName())
	}

	childDir := newDirectory(dirName, dir)
	dir.directories[dirName] = childDir

	return &childDir, nil
}

func (dir *directory) Insert(directoryPath string, files ...File) error {
	directoryPath, _ = strings.CutPrefix(directoryPath, "/")

	for _, file := range files {
		if err := dir.insert(directoryPath, file); err != nil {
			return err
		}
	}

	return nil
}

func (dir directory) FindDirectory(dirName string) *directory {
	if len(dirName) == 0 {
		return nil
	}

	dirs := strings.Split(dirName, "/")
	if len(dirs) == 1 {
		d, ok := dir.directories[dirs[0]]
		if ok {
			return &d
		}
		return nil
	}

	d, ok := dir.directories[dirs[0]]
	if !ok {
		return nil
	}

	return d.FindDirectory(strings.Join(dirs[1:], "/"))

}

func (dir *directory) insert(directoryPath string, file File) error {
	if exists := dir.exists(directoryPath, file.FileName()); exists {
		return fmt.Errorf("file %s already exists at %s", file.FileName(), dir.DirName())
	}

	if len(directoryPath) == 0 {
		return dir.AddFile(file)
	}

	dirs := strings.Split(directoryPath, "/")
	childDirName := dirs[0]

	directoryPath = strings.Join(dirs[1:], "/")

	if childDir, ok := dir.directories[childDirName]; ok {
		return childDir.insert(directoryPath, file)
	}

	newChildDir, err := dir.AddChild(childDirName)
	if err != nil {
		return err
	}

	return newChildDir.insert(directoryPath, file)
}

func (dir directory) Walk(fn WalkFn) error {
	for _, file := range dir.files {
		err := fn(dir.DirName(), file)
		if err != nil {
			return err
		}
	}

	for _, d := range dir.directories {
		err := d.Walk(fn)
		if err != nil {
			return err
		}
	}
	return nil
}
