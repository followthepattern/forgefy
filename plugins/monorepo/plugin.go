package monorepo

import (
	"io/fs"
	"path"

	"github.com/followthepattern/forgefy/forgeio"
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

func (builder MonoRepo) Build(pm productmap.ProductMap, productSpec specification.Product) (err error) {
	dir := templates.Files

	return fs.WalkDir(dir, ".", func(filepath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if forgeio.ExcludeTemplate(filepath, productSpec.ExcludeDagger) {
			return nil
		}

		if !forgeio.IsForgeTemplate(filepath) {
			return nil
		}

		content, err := fs.ReadFile(dir, filepath)
		if err != nil {
			return err
		}

		filepath = forgeio.CleanFilepath(filepath, forgeio.DAGGER_FILE_TOKEN)
		filepath = path.Join(templates.RootDirectory(), filepath)

		file := productmap.NewFile(
			filepath,
			string(content),
		).WithData(productSpec)

		return pm.Insert(file)
	})
}
