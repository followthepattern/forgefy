package projectmap

import "github.com/followthepattern/forgefy/featureset"

type Builder struct {
}

func NewBuilder() Builder {
	return Builder{}
}

func (f Builder) Create(fs featureset.FeatureSet) ProjectMap {
	return ProjectMap{}
}
