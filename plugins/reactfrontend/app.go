package reactfrontend

import (
	"io/fs"
	"path"
	"strings"

	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/plugins/monorepo/templates/apps"
	"github.com/followthepattern/forgefy/plugins/reactfrontend/templates"
	"github.com/followthepattern/forgefy/productmap"
	"github.com/followthepattern/forgefy/specification"
)

var _ plugins.App = &ReactFrontend{}

type ReactFrontend struct{}

func (ReactFrontend) Name() string {
	return "React Front-end application"
}

func (ReactFrontend) Type() string {
	return "react-frontend"
}

func (ReactFrontend) AddDefaultFiles(pm productmap.ProductMap, fs specification.Product) error {
	return nil
}

func (plugin ReactFrontend) Build(pm productmap.ProductMap, product specification.Product, app specification.App) error {
	dir := templates.Files

	return fs.WalkDir(dir, ".", plugin.createWalkFn(pm, product, app))
}

func (plugin ReactFrontend) createWalkFn(pm productmap.ProductMap, product specification.Product, app specification.App) func(filepath string, d fs.DirEntry, err error) error {
	features := append(product.Features, app.Features...)

	return func(filepath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(filepath, ".tmpl") {
			return nil
		}

		content, err := fs.ReadFile(templates.Files, filepath)
		if err != nil {
			return err
		}

		filepath = strings.TrimSuffix(filepath, ".tmpl")

		filepath = path.Join(apps.Directory(), app.AppName, filepath)

		if !strings.Contains(filepath, "(feature)") {
			dirName, fileName := path.Split(filepath)
			file := productmap.NewFile(
				fileName,
				string(content),
			).WithData(app)

			return pm.Insert(dirName, file)
		}

		for _, feature := range features {
			newFilePath := strings.ReplaceAll(filepath, "(feature)", feature.ToDirName())

			dirName, fileName := path.Split(newFilePath)
			file := productmap.NewFile(
				fileName,
				string(content),
			).WithData(FeatureTemplateModel{
				Feature: feature,
				App:     app,
			})

			err = pm.Insert(dirName, file)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
