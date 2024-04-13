package forgefy

import (
	"log/slog"

	"github.com/followthepattern/forgefy/appbuilder"
	"github.com/followthepattern/forgefy/featureset"
	"github.com/followthepattern/forgefy/io"
	"github.com/followthepattern/forgefy/productmap"
)

type Forgefy struct{}

func New() Forgefy {
	return Forgefy{}
}

func (f Forgefy) Forge(yaml string, fw io.Writer) error {
	fs, err := featureset.UnmarshalYaml([]byte(yaml))
	if err != nil {
		return err
	}

	builder := appbuilder.NewBuilder(fs)

	product, err := builder.Build()
	if err != nil {
		return err
	}

	product.Walk(func(folderName string, f productmap.File) {
		content, err := f.Content()
		if err != nil {
			slog.Error(err.Error())
		}

		err = fw.Write(folderName, f.FileName(), content)
		if err != nil {
			slog.Error(err.Error())
		}
	})

	return nil
}
