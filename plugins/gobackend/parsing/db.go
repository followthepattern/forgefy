package parsing

import (
	"fmt"
	"strings"

	gomodels "github.com/followthepattern/forgefy/plugins/gobackend/models"
	"github.com/followthepattern/forgefy/specification/models"
)

func NameDB(f interface{}) string {
	switch v := f.(type) {
	case models.Field:
		return v.FieldNameSnakeCaseLower()
	case gomodels.Field:
		return v.FieldNameSnakeCaseLower()
	}

	return ""
}

func TypeDB(f interface{}) string {
	t := models.Field{}

	switch v := f.(type) {
	case models.Field:
		t = v
	case gomodels.Field:
		t = v.Field
	}

	switch t.Type {
	case "string":
		return "varchar"
	case "bool":
		return "boolean"
	case "int":
		return "integer"
	}
	return ""
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
