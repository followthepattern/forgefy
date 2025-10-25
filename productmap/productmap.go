package productmap

import (
	"fmt"
	"maps"
	"strings"
	"text/template"
)

const ROOT_DIRECTORY = ""

type WalkFn func(directoryName string, f File) error

type ProductMap struct {
	files   map[string]File
	funcMap template.FuncMap
}

func NewProductMap() ProductMap {
	return ProductMap{
		files:   make(map[string]File),
		funcMap: template.FuncMap{},
	}
}

func (pm *ProductMap) WithFuncMap(fm template.FuncMap) {
	for k, v := range fm {
		pm.funcMap[k] = v
	}
}

func (dir ProductMap) Exists(filePath string) bool {
	_, ok := dir.files[filePath]

	return ok
}

func (pm *ProductMap) Insert(file File) error {
	_, ok := pm.files[file.FilePath()]

	if ok {
		return fmt.Errorf("file %s already exists", file.FilePath())
	}

	file = file.WithFuncMap(pm.funcMap)
	pm.files[file.FilePath()] = file

	return nil
}

func (pm *ProductMap) Walk(fn WalkFn) error {
	for filePath, file := range pm.files {
		err := fn(filePath, file)
		if err != nil {
			return err
		}
	}

	return nil
}

func (pm *ProductMap) Include(includeFiles map[string]struct{}) {
	maps.DeleteFunc(pm.files, func(key string, _ File) bool {
		for includeKey := range includeFiles {
			if strings.HasPrefix(key, includeKey) {
				return false
			}

		}

		return true
	})
}

func (pm *ProductMap) Exclude(excludedFiles map[string]struct{}) {
	maps.DeleteFunc(pm.files, func(key string, _ File) bool {
		_, ok := excludedFiles[key]
		return ok
	})
}
