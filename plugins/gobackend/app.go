package gobackend

import (
	"io/fs"
	"path"
	"strings"

	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/plugins/gobackend/apptemplates"
	"github.com/followthepattern/forgefy/plugins/gobackend/models"
	"github.com/followthepattern/forgefy/plugins/monorepo/templates/apps"
	"github.com/followthepattern/forgefy/productmap"
	"github.com/followthepattern/forgefy/specification"
)

type App struct {
	specification.App
	DbPort     int
	CerbosPort int
}

func (a App) AppNameCapital() string {
	return strings.ToUpper(a.AppName)
}

func (a App) AppNamePackage() string {
	return strings.ToLower(a.AppName)
}

func (a App) Features() []Feature {
	features := make([]Feature, len(a.App.Features))
	for i, feature := range a.App.Features {

		fields := make([]models.Field, len(feature.Fields))
		for j, field := range feature.Fields {
			fields[j] = models.Field{Field: field}
		}
		features[i] = Feature{Feature: feature, Fields: fields}
	}
	return features
}

var _ plugins.App = &GoBackendPluginApp{}

type GoBackendPluginApp struct {
	port       int
	dbPort     int
	cerbosPort int
}

func (GoBackendPluginApp) Name() string {
	return "Go Back-end application"
}

func (GoBackendPluginApp) Type() string {
	return "go-backend"
}

func (plugin *GoBackendPluginApp) GetNextPortNumber() int {
	port := plugin.port

	plugin.port++

	return port
}

func (plugin *GoBackendPluginApp) GetNextDBPort() int {
	port := plugin.dbPort

	plugin.dbPort++

	return port
}

func (plugin *GoBackendPluginApp) GetNextCerbosPort() int {
	port := plugin.cerbosPort

	plugin.cerbosPort++

	return port
}

func (plugin *GoBackendPluginApp) Build(pm productmap.ProductMap, app specification.App) error {
	dir := apptemplates.EntireDir

	goApp := App{app, plugin.dbPort, plugin.cerbosPort}

	goApp.App.AppPort = plugin.GetNextPortNumber()
	goApp.DbPort = plugin.GetNextDBPort()
	goApp.CerbosPort = plugin.GetNextCerbosPort()

	return fs.WalkDir(dir, ".", plugin.createWalkFn(pm, goApp))
}

func (b GoBackendPluginApp) createWalkFn(pm productmap.ProductMap, goApp App) func(filepath string, d fs.DirEntry, err error) error {
	dir := apptemplates.EntireDir

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

		if strings.Contains(filepath, "(appName)") {
			filepath = strings.ReplaceAll(filepath, "(appName)", goApp.AppName)
		}

		filepath = strings.TrimSuffix(filepath, ".tmpl")
		filepath = path.Join(apps.Directory(), goApp.AppName, filepath)

		if !strings.Contains(filepath, "[feature]") {
			file := productmap.NewFile(
				filepath,
				string(content),
			).WithData(goApp)

			return pm.Insert(file)
		}

		for _, feature := range goApp.Features() {
			newFilePath := strings.ReplaceAll(filepath, "[feature]", feature.FeatureNamePackage())

			file := productmap.NewFile(
				newFilePath,
				string(content),
			).WithData(FeatureTemplateModel{
				Feature: feature,
				App:     goApp,
			})

			err = pm.Insert(file)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
