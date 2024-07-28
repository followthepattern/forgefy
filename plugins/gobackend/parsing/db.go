package parsing

import (
	"fmt"
	"strings"

	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/naming"
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

func DBType(f models.Field) string {
	switch f.Type {
	case "string":
		return "varchar"
	case "bool":
		return "boolean"
	case "int":
		return "integer"
	}
	return "unknown"
}

func ValueDB(f models.Field) string {
	if len(f.Value) == 0 {
		return "NULL"
	}

	return fmt.Sprintf("'%s'", f.Value)
}

func NullableDB(f models.Field) string {
	if f.Nullable || strings.ToLower(f.Type) == models.IDFieldType {
		return ""
	}

	return "NOT NULL"
}
