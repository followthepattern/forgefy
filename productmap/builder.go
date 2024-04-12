package productmap

import (
	"github.com/followthepattern/forgefy/apptemplates"
	"github.com/followthepattern/forgefy/apptemplates/apps/backend"
	"github.com/followthepattern/forgefy/apptemplates/apps/frontend"
	"github.com/followthepattern/forgefy/featureset"
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

func (builder Builder) Build() (ProductMap, error) {
	fs := builder.fs

	pm := NewProductMap()

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

func (builder Builder) addDefaultFiles(_ ProductMap) error { return nil }

func (builder Builder) addlocalDevFiles(pm ProductMap) (err error) {
	dir := apptemplates.RootDirectory(builder.rootDirectoryName)
	return pm.Insert(dir, NewFileFromTemplate(apptemplates.DockerCompose))
}

func (b Builder) addAppSpecificFiles(pm ProductMap, app featureset.App) (err error) {
	switch app.Type {
	case featureset.Backend:
		err = b.addBackendFiles(pm, app)
	case featureset.Frontend:
		err = b.addFrontendFiles(pm, app)
	}
	return
}

func (b Builder) addBackendFiles(pm ProductMap, app featureset.App) error {
	dir := backend.Directory(b.rootDirectoryName, app.Name)
	err := pm.Insert(dir, NewFileFromTemplate(backend.GoMod))

	if err != nil {
		return err
	}

	return nil
}

func (b Builder) addFrontendFiles(pm ProductMap, app featureset.App) error {
	dir := frontend.Directory(b.rootDirectoryName, app.Name)
	err := pm.Insert(dir, NewFileFromTemplate(frontend.PackageJSON))
	if err != nil {
		return err
	}

	return nil
}
