package productmap

import (
	"github.com/followthepattern/forgefy/apptemplates"
	"github.com/followthepattern/forgefy/apptemplates/apps/backend"
	"github.com/followthepattern/forgefy/apptemplates/apps/frontend"
	"github.com/followthepattern/forgefy/featureset"
)

type Builder struct {
	fs featureset.FeatureSet
}

func NewBuilder(fs featureset.FeatureSet) Builder {
	return Builder{
		fs: fs,
	}
}

func (builder Builder) Build() ProductMap {
	productName := builder.fs.ProductName
	fs := builder.fs

	pm := NewProductMap(productName)

	builder.addDefaultFiles(pm)
	builder.addlocalDevFiles(pm)

	for _, app := range fs.Apps {
		builder.addAppSpecificFiles(pm, app)
	}

	return pm
}

func (builder Builder) addDefaultFiles(pm ProductMap) {}

func (builder Builder) addlocalDevFiles(pm ProductMap) {
	pm.Insert("", NewFileFromTemplate(apptemplates.DockerCompose))
}

func (b Builder) addAppSpecificFiles(pm ProductMap, app featureset.App) {
	switch app.Type {
	case featureset.Backend:
		b.addBackendFiles(pm, app)
	case featureset.Frontend:
		b.addFrontendFiles(pm, app)
	}
}

func (b Builder) addBackendFiles(pm ProductMap, app featureset.App) {
	pm.Insert(app.Name, NewFileFromTemplate(backend.GoMod))
}

func (b Builder) addFrontendFiles(pm ProductMap, app featureset.App) {
	pm.Insert(app.Name, NewFileFromTemplate(frontend.PackageJSON))
}
