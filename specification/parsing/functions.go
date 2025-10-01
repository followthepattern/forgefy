package parsing

import (
	"slices"
	"strings"

	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/types"
)

func CreateTableViewFilter(t types.TypeRegistry) func([]models.Field) []models.Field {
	return func(f []models.Field) []models.Field {
		copiedFields := slices.Clone(f)

		result := slices.DeleteFunc(copiedFields, func(e models.Field) bool {
			return t.GetType(e.Type) == types.ID ||
				t.GetType(e.Type) == types.File ||
				t.GetType(e.Type) == types.Text
		})

		if len(result) <= 3 {
			return result
		}

		return result[:3]
	}
}

func SList(a ...string) []string {
	return a
}

func TypeCRUD(features []specification.Feature) []specification.Feature {
	copiedFields := slices.Clone(features)

	result := slices.DeleteFunc(copiedFields, func(e specification.Feature) bool {
		return e.FeatureType == specification.CRUD
	})

	return result
}

func NoneMobile(apps []specification.App) []specification.App {
	copiedFields := slices.Clone(apps)

	result := slices.DeleteFunc(copiedFields, func(a specification.App) bool {
		return strings.Contains(string(a.AppType), "mobile")
	})

	return result
}
