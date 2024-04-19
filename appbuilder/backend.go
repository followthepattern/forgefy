package appbuilder

import (
	"github.com/followthepattern/forgefy/apptemplates/apps/backend"
	"github.com/followthepattern/forgefy/apptemplates/apps/backend/features/feature"
	"github.com/followthepattern/forgefy/featureset"
	"github.com/followthepattern/forgefy/productmap"
)

func (b Builder) addBackendFiles(pm productmap.ProductMap, app featureset.App, features []featureset.Feature) error {
	dir := backend.Directory(b.rootDirectoryName, app.AppName)

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

func (b Builder) addFeature(pm productmap.ProductMap, app featureset.App, f featureset.Feature) error {
	dir := feature.Directory(b.rootDirectoryName, app.AppName, f.FeatureName)

	return pm.Insert(dir, feature.Files...)
}
