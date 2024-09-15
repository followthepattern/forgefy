package parsing

import (
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
