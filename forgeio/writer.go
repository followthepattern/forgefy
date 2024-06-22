package forgeio

import "io"

type Writer interface {
	Write(filePath string, writerFn func(io.Writer) error) error
}
