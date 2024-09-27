package parsing

import (
	"slices"

	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/types"
)

func CreateTableView(t types.TypeRegistry) func([]models.Field) []models.Field {
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
