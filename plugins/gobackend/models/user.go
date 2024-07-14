package models

import "github.com/followthepattern/forgefy/specification/models"

type User models.User

func (u User) ToFields() []Field {
	return []Field{
		{u.ID},
		{u.Email},
		{u.FirstName},
		{u.LastName},
		{u.Password},
		{u.CreationUserID},
		{u.UpdateUserID},
		{u.CreatedAt},
		{u.UpdatedAt},
	}
}
