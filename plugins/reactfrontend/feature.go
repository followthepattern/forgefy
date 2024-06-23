package reactfrontend

import (
	"strings"

	"github.com/followthepattern/forgefy/plugins/reactfrontend/models"
	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/naming"
)

type Feature struct {
	specification.Feature
	Fields []models.Field
}

func (f Feature) FeatureNameGraphQL() string {
	return naming.CapitalizeFirst(f.FeatureName)
}

func (f Feature) FeatureToFileSuffix() string {
	return naming.CapitalizeFirst(f.FeatureName)
}

func (f Feature) FeatureNameURL() string {
	return strings.ToLower(f.FeatureName)
}
