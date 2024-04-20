package forgefy

import (
	"github.com/followthepattern/forgefy/appbuilder"
	"github.com/followthepattern/forgefy/featureset"
	"github.com/followthepattern/forgefy/io"
	"github.com/followthepattern/forgefy/productmap"
)

type Forgefy struct{}

func New() Forgefy {
	return Forgefy{}
}

func (f Forgefy) Forge(yaml string, fw io.Writer) (string, error) {
	fs, err := featureset.UnmarshalYaml([]byte(yaml))
	if err != nil {
		return "", err
	}

	builder := appbuilder.NewBuilder(fs)

	product, err := builder.Build()
	if err != nil {
		return fs.ProductName, err
	}

	err = product.Walk(func(folderName string, file productmap.File) error {
		content, err := file.Content()
		if err != nil {
			return err
		}
		return fw.Write(folderName, file.FileName(), content)
	})

	return fs.ProductName, err
}
