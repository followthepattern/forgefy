package monorepo

import (
	"io/fs"
	"path"
	"strings"

	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/plugins/monorepo/templates"
	"github.com/followthepattern/forgefy/productmap"
	"github.com/followthepattern/forgefy/specification"
)

var _ plugins.Plugin = MonoRepo{}

type MonoRepo struct{}

func (MonoRepo) Name() string { return "Mono Repo Plugin" }

func (MonoRepo) Apps() []plugins.App {
	return nil
}

func (builder MonoRepo) Build(pm productmap.ProductMap, productSpec specification.Product) error {
	dir := templates.Files

	return fs.WalkDir(dir, ".", func(filepath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(filepath, ".tmpl") {
			return nil
		}

		content, err := fs.ReadFile(dir, filepath)
		if err != nil {
			return err
		}

		filepath = strings.TrimSuffix(filepath, ".tmpl")

		filepath = path.Join(templates.RootDirectory(), filepath)

		dirName, fileName := path.Split(filepath)
		file := productmap.NewFile(
			fileName,
			string(content),
		).WithData(productSpec)

		return pm.Insert(dirName, file)
	})
}
