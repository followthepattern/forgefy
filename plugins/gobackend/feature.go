package gobackend

import (
	"strings"

	"github.com/followthepattern/forgefy/plugins/gobackend/models"
	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/naming"
)

type Feature struct {
	specification.Feature
	Fields []models.Field
}

func (f Feature) FeatureNameGo() string {
	return naming.CapitalizeFirst(f.FeatureName)
}

func (f Feature) FeatureNameGraphQL() string {
	return naming.CapitalizeFirst(f.FeatureName)
}

func (f Feature) FeatureNamePackage() string {
	return strings.ToLower(f.FeatureName)
}

func (f Feature) FeatureNameDBTable() string {
	return strings.ToLower(f.FeatureName)
}

func (f Feature) FeatureNameURL() string {
	return strings.ToLower(f.FeatureName)
}

func (f Feature) IDField() models.Field {
	return f.Fields[0]
}
