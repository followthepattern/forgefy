package gobackend

import (
	"github.com/followthepattern/forgefy/featureset"
	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/plugins/gobackend/apptemplates"
	"github.com/followthepattern/forgefy/plugins/gobackend/apptemplates/features/feature"
	"github.com/followthepattern/forgefy/plugins/gobackend/apptemplates/tests/integration/testdata"
	"github.com/followthepattern/forgefy/productmap"
)

var _ plugins.App = &GoBackendPluginApp{}

type GoBackendPluginApp struct{}

func (GoBackendPluginApp) Name() string {
	return "Go Back-end application"
}

func (GoBackendPluginApp) Type() string {
	return "go-backend"
}

func (GoBackendPluginApp) AddDefaultFiles(pm productmap.ProductMap, fs featureset.FeatureSet) error {
	return nil
}

func (plugin GoBackendPluginApp) Builder(pm productmap.ProductMap, fs featureset.FeatureSet, app featureset.App) error {
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

func (b GoBackendPluginApp) addTests(pm productmap.ProductMap, app featureset.App, features []featureset.Feature) error {
	dir := testdata.Directory(app.AppName)

	d := struct {
		Features []featureset.Feature
		featureset.App
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

func (b GoBackendPluginApp) addFeature(pm productmap.ProductMap, app featureset.App, f featureset.Feature) error {
	dir := feature.Directory(app.AppName, f.FeatureName)

	d := struct {
		featureset.Feature
		featureset.App
	}{
		Feature: f,
		App:     app,
	}

	for i := 0; i < len(feature.Files); i++ {
		feature.Files[i] = feature.Files[i].WithData(d)
	}

	return pm.Insert(dir, feature.Files...)
}
