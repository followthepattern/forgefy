package parsing

import (
	"strings"

	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/naming"
)

func HTMLType(f models.Field) string {
	switch f.Type {
	case "string":
		return "text"
	}
	return "input"
}

func URL(f specification.Feature) string {
	return strings.ToLower(naming.ToSnakeCase(f.FeatureName))
}
