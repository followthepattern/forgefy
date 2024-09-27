package parsing

import (
	"fmt"
	"strings"

	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/naming"
	"github.com/followthepattern/forgefy/specification/types"
)

func DBName(i interface{}) string {
	switch tt := i.(type) {
	case models.Field:
		return fieldDBName(tt)
	case specification.Feature:
		return featureDBName(tt)
	}

	return "uknown"
}

func fieldDBName(f models.Field) string {
	return f.FieldNameSnakeCaseLower()
}

func featureDBName(f specification.Feature) string {
	return strings.ToLower(naming.ToSnakeCase(f.FeatureName))
}

func CreateDBType(t types.TypeRegistry) func(f models.Field) string {
	return func(f models.Field) string {
		switch t.GetType(f.Type) {
		case types.Boolean:
			return "boolean"
		case types.Number:
			return "integer"
		case types.ID:
			return "varchar"
		case types.String:
			return "varchar"
		case types.Text:
			return "varchar"
		case types.Time,
			types.Date,
			types.DateTime:
			return "timestamp"
		case types.File:
			return "varchar"
		}
		return "unknown"
	}
}

func ValueDB(f models.Field) string {
	if len(f.Value) != 0 {
		return fmt.Sprintf("'%s'", f.Value)
	}

	if f.Nullable {
		return "NULL"
	}

	return `''`
}

func CreateDBNullable(t types.TypeRegistry) func(f models.Field) string {
	return func(f models.Field) string {
		if f.Nullable || t.GetType(f.Type) == types.ID {
			return ""
		}

		return "NOT NULL"
	}
}
