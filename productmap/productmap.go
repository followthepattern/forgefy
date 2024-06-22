package productmap

import "fmt"

type WalkFn func(directoryName string, f File) error

type ProductMap struct {
	files map[string]File
}

func NewProductMap() ProductMap {
	return ProductMap{
		files: make(map[string]File),
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
