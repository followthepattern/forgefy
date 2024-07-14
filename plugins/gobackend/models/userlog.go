package models

import "github.com/followthepattern/forgefy/specification/models"

type Userlog models.Userlog

func (u Userlog) ToFields() []Field {
	return []Field{
		{u.CreationUserID},
		{u.UpdateUserID},
		{u.CreatedAt},
		{u.UpdatedAt},
	}
}
