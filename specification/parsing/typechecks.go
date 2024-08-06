package parsing

import (
	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/types"
)

func CreateIsNumber(t types.TypeRegistry) func(models.Field) bool {
	return func(f models.Field) bool {
		return t.GetType(f.Type) == types.Number
	}
}

func CreateHasNumber(t types.TypeRegistry) func(specification.Feature) bool {
	return func(f specification.Feature) bool {
		for _, field := range f.Fields {
			if t.GetType(field.Type) == types.Number {
				return true
			}
		}
		return false
	}
}
