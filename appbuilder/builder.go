package appbuilder

import (
	"github.com/followthepattern/forgefy/apptemplates"
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
		err = builder.addAppSpecificFiles(pm, app, fs.Features)
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

func (b Builder) addAppSpecificFiles(pm productmap.ProductMap, app featureset.App, features []featureset.Feature) (err error) {
	switch app.AppType {
	case featureset.Backend:
		err = b.addBackendFiles(pm, app, features)
	case featureset.Frontend:
		err = b.addFrontendFiles(pm, app)
	}
	return
}
