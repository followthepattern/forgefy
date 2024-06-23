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

func (f Feature) FieldTypeGoName() string {
	return naming.CapitalizeFirst(f.FeatureName)
}

func (f Feature) PackageName() string {
	return strings.ToLower(f.FeatureName)
}

func (f Feature) FeatureTableName() string {
	return strings.ToLower(f.FeatureName)
}

func (f Feature) IDField() models.Field {
	return f.Fields[0]
}
