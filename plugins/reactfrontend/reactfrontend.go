package reactfrontend

import (
	"github.com/followthepattern/forgefy/featureset"
	"github.com/followthepattern/forgefy/plugins"
	"github.com/followthepattern/forgefy/plugins/reactfrontend/apptemplates"
	"github.com/followthepattern/forgefy/productmap"
)

var _ plugins.App = &ReactFrontend{}

type ReactFrontend struct{}

func (ReactFrontend) Name() string {
	return "React Front-end application"
}

func (ReactFrontend) Type() string {
	return "react-frontend"
}

func (ReactFrontend) AddDefaultFiles(pm productmap.ProductMap, fs featureset.FeatureSet) error {
	return nil
}

func (plugin ReactFrontend) Builder(pm productmap.ProductMap, fs featureset.FeatureSet, app featureset.App) error {
	dir := apptemplates.Directory(app.AppName)

	return pm.Insert(dir,
		apptemplates.PackageJSON.WithData(app))
}
