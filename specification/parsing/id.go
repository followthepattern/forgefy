package parsing

import (
	"slices"

	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/types"
)

func CreateFindIDFunc(t types.TypeRegistry) func(f specification.Feature) models.Field {
	return func(f specification.Feature) models.Field {
		for _, field := range f.Fields {
			if t.GetType(field.Type) == types.ID {
				return field
			}
		}

		if len(f.Fields) > 0 {
			return f.Fields[0]
		}

		return models.Field{
			Name: "<unknown id field>",
			Type: "unknown",
		}
	}
}

func CreateNoneID(t types.TypeRegistry) func([]models.Field) []models.Field {
	return func(f []models.Field) []models.Field {
		copiedFields := slices.Clone(f)

		return slices.DeleteFunc(copiedFields, func(e models.Field) bool {
			return t.GetType(e.Type) == types.ID
		})
	}
}
