package gobackend

import (
	"strings"

	"github.com/followthepattern/forgefy/specification"
)

type Feature struct {
	specification.Feature
}

func (f Feature) GoTypeName() string {
	return specification.CapitalizeFirst(f.FeatureName)
}

func (f Feature) PackageName() string {
	return strings.ToLower(f.FeatureName)
}

func (f Feature) FeatureTableName() string {
	return strings.ToLower(f.FeatureName)
}

func (f Feature) Fields() []specification.Field {
	fields := f.Feature.Fields
	for i := range fields {
		fields[i].Type = f.overrideFields(fields[i].Type)
	}

	return fields
}

func (f Feature) overrideFields(fieldType string) string {
	switch fieldType {
	case "string":
		return "types.String"
	default:
		return fieldType
	}
}
