package defaults

import (
	"fmt"

	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/naming"
	"github.com/google/uuid"
)

func Roles(admin models.User, features []string) (roles []models.Role) {
	for _, feature := range features {
		role := models.Role{
			ID: models.Field{
				Name:     "ID",
				Value:    uuid.NewString(),
				Nullable: false,
			},
			Code: models.Field{
				Name:     "Code",
				Value:    fmt.Sprintf("%s:editor", naming.LowerFirst(feature)),
				Nullable: false,
			},
			Name: models.Field{
				Name:     "Name",
				Value:    fmt.Sprintf("%s Editor", feature),
				Nullable: false,
			},
			Userlog: DefaultUserlog(admin),
		}
		roles = append(roles, role)
	}
	return
}
