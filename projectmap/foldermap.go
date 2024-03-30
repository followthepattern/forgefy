package projectmap

type FolderMap struct {
	folderName string
	files      map[string]File
	folders    map[string]FolderMap
}
