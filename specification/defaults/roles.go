package defaults

import (
	"fmt"
	"time"

	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/naming"
	"github.com/google/uuid"
)

func Roles(features []string) (roles []models.Role) {
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
			Userlog: models.Userlog{
				CreatedAt: models.Field{
					Name:     "CreatedAt",
					Value:    time.Now().Format("2006-01-02 15:04:05"),
					Nullable: false,
				},
				UpdatedAt: models.Field{
					Name:     "UpdatedAt",
					Nullable: true,
				},
				CreationUserID: models.Field{
					Name:  "CreationUserID",
					Value: "613254df-c779-479c-9d76-b8036e342979",
				},
			},
		}
		roles = append(roles, role)
	}
	return
}
