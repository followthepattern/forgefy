package parsing

import (
	"fmt"
	"strings"

	"github.com/followthepattern/forgefy/specification/models"
)

func NameDB(f models.Field) string {
	return f.FieldNameSnakeCaseLower()
}

func TypeDB(f models.Field) string {
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
