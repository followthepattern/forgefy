package productmap

import (
	"github.com/followthepattern/forgefy/apptemplates"
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

func (builder Builder) Create() ProductMap {
	pm := NewProductMap(builder.fs.ProductName)
	pm = builder.addDefaults(pm)
	pm = builder.addlocalDevFiles(pm)

	return pm
}

func (f Builder) addDefaults(pm ProductMap) ProductMap {
	return pm
}

func (f Builder) addlocalDevFiles(pm ProductMap) ProductMap {
	pm.Insert("", NewFileFromTemplate(apptemplates.DockerCompose))
	return pm
}
