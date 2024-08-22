package parsing

import (
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/types"
)

func CreateJSType(t types.TypeRegistry) func(f models.Field) string {
	return func(f models.Field) string {
		switch t.GetType(f.Type) {
		case types.String:
			return "string"
		case types.Number:
			return "number"
		case types.DateTime,
			types.Date:
			return "ZonedDateTime"
		}
		return f.Type
	}
}
