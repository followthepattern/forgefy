package gobackend

import (
	"io/fs"
	"path"
	"strings"

	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/plugins/gobackend/apptemplates"
	"github.com/followthepattern/forgefy/plugins/gobackend/apptemplates/models"
	"github.com/followthepattern/forgefy/plugins/gobackend/apptemplates/tests/integration/testdata"
	"github.com/followthepattern/forgefy/plugins/gobackend/apptemplates/types"
	"github.com/followthepattern/forgefy/productmap"
	"github.com/followthepattern/forgefy/specification"
)

var _ plugins.App = &GoBackendPluginApp{}

type GoBackendPluginApp struct{}

func (GoBackendPluginApp) Name() string {
	return "Go Back-end application"
}

func (GoBackendPluginApp) Type() string {
	return "go-backend"
}

func (GoBackendPluginApp) AddStaticFiles(pm productmap.ProductMap, fs specification.Product, app specification.App) error {
	if err := pm.Insert(types.Directory(app.AppName),
		types.Files...); err != nil {
		return err
	}

	if err := pm.Insert(models.Directory(app.AppName),
		models.Files...); err != nil {
		return err
	}

	return nil
}

func (plugin GoBackendPluginApp) Builder2(pm productmap.ProductMap, fs specification.Product, app specification.App) error {
	dir := apptemplates.Directory(app.AppName)

	features := append(fs.Features, app.Features...)

	err := plugin.addTests(pm, app, features)
	if err != nil {
		return err
	}

	err = plugin.AddStaticFiles(pm, fs, app)
	if err != nil {
		return err
	}

	return pm.Insert(dir,
		apptemplates.GoMod.WithData(app),
		apptemplates.DockerFile.WithData(app))
}

func (plugin GoBackendPluginApp) Builder(pm productmap.ProductMap, product specification.Product, app specification.App) error {
	dir := apptemplates.EntireDir
	features := append(product.Features, app.Features...)

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

		if strings.Contains(filepath, "[feature]") {
			filepath = path.Join(apptemplates.Directory(app.AppName), filepath)

			for _, feature := range features {
				feat := Feature(feature)

				newFilePath := strings.ReplaceAll(filepath, "[feature]", feat.PackageName())

				dirName, fileName := path.Split(newFilePath)
				file := productmap.NewFile(
					fileName,
					string(content),
				).WithData(FeatureTemplateModel{
					Feature: feat,
					App:     app,
				})

				err = pm.Insert(dirName, file)
				if err != nil {
					return err
				}
			}
		} else {
			filepath = path.Join(apptemplates.Directory(app.AppName), filepath)

			dirName, fileName := path.Split(filepath)
			file := productmap.NewFile(
				fileName,
				string(content),
			).WithData(app)

			return pm.Insert(dirName, file)
		}
		return nil
	})
}

func (b GoBackendPluginApp) addTests(pm productmap.ProductMap, app specification.App, features []specification.Feature) error {
	dir := testdata.Directory(app.AppName)

	d := struct {
		Features []specification.Feature
		specification.App
	}{
		Features: features,
		App:      app,
	}

	for i := 0; i < len(testdata.Files); i++ {
		testdata.Files[i] = testdata.Files[i].WithData(d)
	}

	return pm.Insert(dir,
		testdata.Files...,
	)
}
