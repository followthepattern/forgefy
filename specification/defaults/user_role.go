package defaults

import "github.com/followthepattern/forgefy/specification/models"

func UserRole(admin models.User, roles []models.Role) (userRoles []models.UserRole) {
	for _, role := range roles {
		userRoles = append(userRoles, models.UserRole{
			UserID:  admin.ID,
			RoleID:  role.ID,
			Userlog: DefaultUserlog(admin),
		})
	}

	return
}
