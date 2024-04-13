package appbuilder

import (
	"github.com/followthepattern/forgefy/apptemplates"
	"github.com/followthepattern/forgefy/apptemplates/apps/backend"
	"github.com/followthepattern/forgefy/apptemplates/apps/frontend"
	"github.com/followthepattern/forgefy/featureset"
	"github.com/followthepattern/forgefy/productmap"
)

type Builder struct {
	fs                featureset.FeatureSet
	rootDirectoryName string
}

func NewBuilder(fs featureset.FeatureSet) Builder {
	return Builder{
		fs: fs,
	}
}

func (builder Builder) Build() (productmap.ProductMap, error) {
	fs := builder.fs

	pm := productmap.NewProductMap()

	err := builder.addDefaultFiles(pm)
	if err != nil {
		return pm, err
	}

	err = builder.addlocalDevFiles(pm)
	if err != nil {
		return pm, err
	}

	for _, app := range fs.Apps {
		err = builder.addAppSpecificFiles(pm, app)
		if err != nil {
			return pm, err
		}
	}

	return pm, nil
}

func (builder Builder) addDefaultFiles(_ productmap.ProductMap) error { return nil }

func (builder Builder) addlocalDevFiles(pm productmap.ProductMap) (err error) {
	dir := apptemplates.RootDirectory(builder.rootDirectoryName)
	return pm.Insert(dir, apptemplates.DockerCompose)
}

func (b Builder) addAppSpecificFiles(pm productmap.ProductMap, app featureset.App) (err error) {
	switch app.Type {
	case featureset.Backend:
		err = b.addBackendFiles(pm, app)
	case featureset.Frontend:
		err = b.addFrontendFiles(pm, app)
	}
	return
}

func (b Builder) addBackendFiles(pm productmap.ProductMap, app featureset.App) error {
	dir := backend.Directory(b.rootDirectoryName, app.Name)
	err := pm.Insert(dir, backend.GoMod)

	if err != nil {
		return err
	}

	return nil
}

func (b Builder) addFrontendFiles(pm productmap.ProductMap, app featureset.App) error {
	dir := frontend.Directory(b.rootDirectoryName, app.Name)
	packageJSON := frontend.PackageJSON.WithData(frontend.PackageJsonModel{
		FrontendProjectName: app.Name,
	})

	err := pm.Insert(dir, packageJSON)
	if err != nil {
		return err
	}

	return nil
}