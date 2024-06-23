package specification

import (
	"fmt"
	"time"

	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/naming"
	"github.com/google/uuid"
)

type DefaultValues struct {
	Roles []models.Role
}

func (d *DefaultValues) InitDefaultTypes(features []Feature) {
	d.generateRoles(features)
}

func (d *DefaultValues) generateRoles(features []Feature) {
	for _, feature := range features {
		role := models.Role{
			ID: models.Field{
				Name:     "ID",
				Value:    uuid.NewString(),
				Nullable: false,
			},
			Code: models.Field{
				Name:     "Code",
				Value:    fmt.Sprintf("%s:editor", naming.LowerFirst(feature.FeatureName)),
				Nullable: false,
			},
			Name: models.Field{
				Name:     "Name",
				Value:    fmt.Sprintf("%s Editor", feature.FeatureName),
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
		d.Roles = append(d.Roles, role)
	}
}
