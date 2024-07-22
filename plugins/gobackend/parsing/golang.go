package parsing

import "github.com/followthepattern/forgefy/specification/models"

func GoType(f models.Field) string {
	switch f.Type {
	case "string":
		return "types.String"
	case "int":
		return "types.Int64"
	}
	return f.Type
}

func AsTag(f models.Field) string {
	return f.FieldNameSnakeCaseLower()
}
