package defaults

import (
	"time"

	"github.com/followthepattern/forgefy/specification/models"
)

func DefaultUserlog(user models.User) models.Userlog {
	return models.Userlog{
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
			Value: user.ID.Value,
		},
		UpdateUserID: models.Field{
			Name:     "UpdateUserID",
			Nullable: true,
		},
	}
}
