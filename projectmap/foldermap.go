package projectmap

import (
	"fmt"
	"path"
)

type FolderMap struct {
	parentFolder *FolderMap
	folderName   string
	files        map[string]File
	folders      map[string]FolderMap
}

func (fm FolderMap) DirName() string {
	if fm.parentFolder == nil {
		return fm.folderName
	}
	return path.Join(fm.parentFolder.DirName(), fm.folderName)
}

func (fm FolderMap) Walk(fn func(folderName string, f File)) {
	for _, file := range fm.files {
		fmt.Println("inside printing files in folderMap")
		fmt.Println(fm.DirName())
		fn(fm.DirName(), file)
	}

	for _, dir := range fm.folders {
		fmt.Println("inside walking folderMap")
		dir.Walk(fn)
	}
}
