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

func (ReactFrontend) AddDefaultFiles(pm productmap.ProductMap, fs specification.Product) error {
	return nil
}

func (plugin ReactFrontend) Build(pm productmap.ProductMap, product specification.Product, app specification.App) error {
	dir := templates.Files

	return fs.WalkDir(dir, ".", plugin.createWalkFn(pm, product, app))
}

func (plugin ReactFrontend) createWalkFn(pm productmap.ProductMap, product specification.Product, app specification.App) func(filepath string, d fs.DirEntry, err error) error {
	reactApp := App{app}
	reactApp.App.Features = append(product.Features, app.Features...)

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
			).WithData(reactApp)

			return pm.Insert(dirName, file)
		}

		for _, feature := range reactApp.Features() {
			newFilePath := strings.Replace(filepath, "(feature)", feature.ToDirName(), 1)
			newFilePath = strings.Replace(newFilePath, "(feature)", feature.FeatureToFileSuffix(), 1)

			dirName, fileName := path.Split(newFilePath)
			file := productmap.NewFile(
				fileName,
				string(content),
			).WithData(FeatureTemplateModel{
				Feature: feature,
				App:     reactApp,
			})

			err = pm.Insert(dirName, file)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
