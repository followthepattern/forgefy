package parsing

import "github.com/followthepattern/forgefy/specification/models"

func JSType(f models.Field) string {
	switch f.Type {
	case "string":
		return "string"
	case "int":
		return "number"
	}
	return f.Type
}
