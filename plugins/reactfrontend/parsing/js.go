package parsing

import (
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/types"
)

func CreateJSType(t types.TypeRegistry) func(f models.Field) string {
	return func(f models.Field) string {
		switch t.GetType(f.Type) {
		case types.Boolean:
			return "boolean"
		case types.ID,
			types.String,
			types.Text:
			return "string"
		case types.Number:
			return "number"
		case types.Time:
			return "Time"
		case types.Date:
			return "Date"
		case types.File:
			return "string"
		}
		return f.Type
	}
}
