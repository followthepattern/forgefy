package forgeio

import "io"

type Writer interface {
	Write(dirName, fileName string, writerFn func(io.Writer) error) error
}
