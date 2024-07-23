package parsing

import (
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/naming"
)

func GraphQLName(f models.Field) string {
	return naming.ToLowerCamelCase(f.Name)
}
