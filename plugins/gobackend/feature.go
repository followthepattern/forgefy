package gobackend

import (
	"strings"

	"github.com/followthepattern/forgefy/specification"
)

type Feature specification.Feature

func (f Feature) GoTypeName() string {
	return specification.CapitalizeFirst(f.FeatureName)
}

func (f Feature) PackageName() string {
	return strings.ToLower(f.FeatureName)
}
