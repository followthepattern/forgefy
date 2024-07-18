package gobackend

import (
	"strings"

	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/naming"
)

type Feature struct {
	specification.Feature
}

func (f Feature) FeatureNameGo() string {
	return naming.ToUpperCamelCase(f.FeatureName)
}

func (f Feature) FeatureNameGraphQL() string {
	return naming.ToUpperCamelCase(f.FeatureName)
}

func (f Feature) FeatureNamePackage() string {
	return strings.ToLower(naming.ToSnakeCase(f.FeatureName))
}

func (f Feature) FeatureNameDBTable() string {
	return strings.ToLower(naming.ToSnakeCase(f.FeatureName))
}

func (f Feature) FeatureNameURL() string {
	return strings.ToLower(naming.ToSnakeCase(f.FeatureName))
}
