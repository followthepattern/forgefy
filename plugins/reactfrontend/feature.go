package reactfrontend

import (
	"strings"

	"github.com/followthepattern/forgefy/plugins/reactfrontend/models"
	"github.com/followthepattern/forgefy/specification"
)

type Feature struct {
	specification.Feature
	Fields []models.Field
}

func (f Feature) FeatureTypeName() string {
	return specification.CapitalizeFirst(f.FeatureName)
}

func (f Feature) FeatureHumanReadableName() string {
	return specification.CapitalizeFirst(f.FeatureName)
}

func (f Feature) FeatureURL() string {
	return strings.ToLower(f.FeatureName)
}

func (f Feature) FeatureToFileSuffix() string {
	return specification.CapitalizeFirst(f.FeatureName)
}
