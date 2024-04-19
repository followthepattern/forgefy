package appbuilder

import (
	"github.com/followthepattern/forgefy/apptemplates/apps/backend"
	"github.com/followthepattern/forgefy/apptemplates/apps/backend/features/feature"
	"github.com/followthepattern/forgefy/apptemplates/apps/backend/tests/integration/testdata"
	"github.com/followthepattern/forgefy/featureset"
	"github.com/followthepattern/forgefy/productmap"
)

func (b Builder) addBackendFiles(pm productmap.ProductMap, app featureset.App, features []featureset.Feature) error {
	dir := backend.Directory(b.rootDirectoryName, app.AppName)

	err := b.addTests(pm, app, features)
	if err != nil {
		return err
	}

	for _, feat := range features {
		err := b.addFeature(pm, app, feat)
		if err != nil {
			return err
		}
	}

	return pm.Insert(dir,
		backend.GoMod.WithData(app),
		backend.DockerFile.WithData(app))
}

func (b Builder) addTests(pm productmap.ProductMap, app featureset.App, features []featureset.Feature) error {
	dir := testdata.Directory(b.rootDirectoryName, app.AppName)

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

func (b Builder) addFeature(pm productmap.ProductMap, app featureset.App, f featureset.Feature) error {
	dir := feature.Directory(b.rootDirectoryName, app.AppName, f.FeatureName)

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
