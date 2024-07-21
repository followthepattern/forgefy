package defaults

import (
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/google/uuid"
)

var adminID = uuid.NewString()

func AdminUser() models.User {
	user := models.User{
		ID: models.Field{
			Name:  "Id",
			Type:  "uuid",
			Value: adminID,
		},
		Email: models.Field{
			Name:  "Email",
			Type:  "string",
			Value: "admin@admin.com",
		},
		FirstName: models.Field{
			Name:  "FirstName",
			Type:  "string",
			Value: "John",
		},
		LastName: models.Field{
			Name:  "LastName",
			Type:  "string",
			Value: "Jones",
		},
		Password: models.Field{
			Name:  "Password",
			Type:  "string",
			Value: "$2a$10$1UZWtyK2f6BvSlqp6SBzkeGiTP5pdkiuRgvlt4Gd4MZIyfpVWCkYq",
		},
		Active: models.Field{
			Name:  "Active",
			Type:  "boolean",
			Value: "true",
		},
	}

	user.Userlog = DefaultUserlog(user)

	return user
}
