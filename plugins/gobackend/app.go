package gobackend

import (
	"io/fs"
	"path"
	"strings"

	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/plugins/gobackend/apptemplates"
	"github.com/followthepattern/forgefy/plugins/gobackend/apptemplates/features/feature"
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

	for _, feat := range features {
		err := plugin.addFeature(pm, app, feat)
		if err != nil {
			return err
		}
	}

	return pm.Insert(dir,
		apptemplates.GoMod.WithData(app),
		apptemplates.DockerFile.WithData(app))
}

func (plugin GoBackendPluginApp) Builder(pm productmap.ProductMap, product specification.Product, app specification.App) error {
	dir := apptemplates.EntireDir
	// features := append(product.Features, app.Features...)

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

		filepath = path.Join(apptemplates.Directory(app.AppName), filepath)

		dirName, fileName := path.Split(filepath)

		file := productmap.NewFile(
			fileName,
			string(content),
		).WithData(app)

		pm.Insert(dirName, file)

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

func (b GoBackendPluginApp) addFeature(pm productmap.ProductMap, app specification.App, f specification.Feature) error {
	dir := feature.Directory(app.AppName, f.FeatureName)

	d := FeatureTemplateModel{
		Feature: Feature(f),
		App:     app,
	}

	for i := 0; i < len(feature.Files); i++ {
		feature.Files[i] = feature.Files[i].WithData(d)
	}

	return pm.Insert(dir, feature.Files...)
}
