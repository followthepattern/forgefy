package forgeio

type Writer interface {
	Write(folderName string, fileName string, content string) error
}
