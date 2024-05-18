package gobackend

import (
	"io/fs"
	"path"
	"strings"

	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/plugins/gobackend/apptemplates"
	"github.com/followthepattern/forgefy/plugins/monorepo/templates/apps"
	"github.com/followthepattern/forgefy/productmap"
	"github.com/followthepattern/forgefy/specification"
)

type App struct {
	specification.App
}

func (a App) AppNameCapital() string {
	return strings.ToUpper(a.AppName)
}

func (a App) AppNameToPackageName() string {
	return strings.ToLower(a.AppName)
}

func (a App) Features() []Feature {
	features := make([]Feature, len(a.App.Features))
	for i, feature := range a.App.Features {

		fields := make([]Field, len(feature.Fields))
		for j, field := range feature.Fields {
			fields[j] = Field{field}
		}
		features[i] = Feature{Feature: feature, Fields: fields}
	}
	return features
}

var _ plugins.App = &GoBackendPluginApp{}

type GoBackendPluginApp struct{}

func (GoBackendPluginApp) Name() string {
	return "Go Back-end application"
}

func (GoBackendPluginApp) Type() string {
	return "go-backend"
}

func (plugin GoBackendPluginApp) Build(pm productmap.ProductMap, product specification.Product, app specification.App) error {
	dir := apptemplates.EntireDir

	return fs.WalkDir(dir, ".", plugin.createWalkFn(pm, product, app))
}

func (b GoBackendPluginApp) createWalkFn(pm productmap.ProductMap, product specification.Product, app specification.App) func(filepath string, d fs.DirEntry, err error) error {
	dir := apptemplates.EntireDir
	goApp := App{app}
	goApp.App.Features = append(product.Features, app.Features...)

	return func(filepath string, d fs.DirEntry, err error) error {
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

		filepath = path.Join(apps.Directory(), goApp.AppName, filepath)

		if !strings.Contains(filepath, "[feature]") {
			dirName, fileName := path.Split(filepath)
			file := productmap.NewFile(
				fileName,
				string(content),
			).WithData(goApp)

			return pm.Insert(dirName, file)
		}

		for _, feature := range goApp.Features() {

			newFilePath := strings.ReplaceAll(filepath, "[feature]", feature.PackageName())

			dirName, fileName := path.Split(newFilePath)
			file := productmap.NewFile(
				fileName,
				string(content),
			).WithData(FeatureTemplateModel{
				Feature: feature,
				App:     goApp,
			})

			err = pm.Insert(dirName, file)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
