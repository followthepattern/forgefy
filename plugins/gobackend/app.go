package gobackend

import (
	"fmt"

	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/plugins/gobackend/apptemplates"
	"github.com/followthepattern/forgefy/plugins/gobackend/apptemplates/features/feature"
	"github.com/followthepattern/forgefy/plugins/gobackend/apptemplates/tests/integration/testdata"
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

func (GoBackendPluginApp) DockerComposeInfos(appName, appPath string) []plugins.DockerComposeInfo {
	return []plugins.DockerComposeInfo{
		{
			ServiceName:           "db",
			DockerComposeFilePath: fmt.Sprintf("%s/docker-compose.yml", appPath),
		},
		{
			ServiceName:           "cerbos",
			DockerComposeFilePath: fmt.Sprintf("%s/docker-compose.yml", appPath),
		},
		{
			ServiceName:           appName,
			DockerComposeFilePath: fmt.Sprintf("%s/docker-compose.yml", appPath),
		},
	}
}

func (GoBackendPluginApp) AddDefaultFiles(pm productmap.ProductMap, fs specification.Product) error {
	return nil
}

func (plugin GoBackendPluginApp) Builder(pm productmap.ProductMap, fs specification.Product, app specification.App) error {
	dir := apptemplates.Directory(app.AppName)

	features := append(fs.Features, app.Features...)

	err := plugin.addTests(pm, app, features)
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

	d := struct {
		specification.Feature
		specification.App
	}{
		Feature: f,
		App:     app,
	}

	for i := 0; i < len(feature.Files); i++ {
		feature.Files[i] = feature.Files[i].WithData(d)
	}

	return pm.Insert(dir, feature.Files...)
}
