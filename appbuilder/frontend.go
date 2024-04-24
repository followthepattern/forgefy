package appbuilder

import (
	"github.com/followthepattern/forgefy/apptemplates/apps/frontend"
	"github.com/followthepattern/forgefy/featureset"
	"github.com/followthepattern/forgefy/productmap"
)

func (b Builder) addFrontendFiles(pm productmap.ProductMap, app featureset.App) error {
	dir := frontend.Directory(app.AppName)

	return pm.Insert(dir,
		frontend.PackageJSON.WithData(app))
}
