package appbuilder

import (
	"github.com/followthepattern/forgefy/apptemplates/apps/backend"
	"github.com/followthepattern/forgefy/featureset"
	"github.com/followthepattern/forgefy/productmap"
)

func (b Builder) addBackendFiles(pm productmap.ProductMap, app featureset.App) error {
	dir := backend.Directory(b.rootDirectoryName, app.AppName)

	return pm.Insert(dir,
		backend.GoMod.WithData(app),
		backend.DockerFile.WithData(app))
}
