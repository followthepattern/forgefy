package parsing

import (
	"strings"

	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/naming"
	"github.com/followthepattern/forgefy/specification/types"
)

func CreateHTMLType(t types.TypeRegistry) func(models.Field) string {
	return func(f models.Field) string {
		switch t.GetType(f.Type) {
		case types.Boolean:
			return "checkbox"
		case types.String,
			types.Text:
			return "input"
		case types.Number:
			return "text"
		}
		return "input"
	}
}

func URL(f specification.Feature) string {
	return strings.ToLower(naming.ToSnakeCase(f.FeatureName))
}
