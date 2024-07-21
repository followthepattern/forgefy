package defaults

import (
	"github.com/followthepattern/forgefy/specification/models"
)

type Defaults struct {
	Users     []models.User
	Roles     []models.Role
	UserRoles []models.UserRole
}
