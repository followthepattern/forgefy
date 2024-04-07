package projectmap

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

func (builder Builder) Create() ProjectMap {
	pm := NewProjectMap(builder.fs.ProjectName)
	pm = builder.addDefaults(pm)
	pm = builder.addlocalDevFiles(pm)

	return pm
}

func (f Builder) addDefaults(pm ProjectMap) ProjectMap {
	return pm
}

func (f Builder) addlocalDevFiles(pm ProjectMap) ProjectMap {
	pm.Insert("", NewFileFromTemplate(apptemplates.DockerCompose))
	return pm
}
