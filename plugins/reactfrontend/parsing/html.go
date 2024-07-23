package parsing

import "github.com/followthepattern/forgefy/specification/models"

func HTMLType(f models.Field) string {
	switch f.Type {
	case "string":
		return "text"
	}
	return "input"
}
