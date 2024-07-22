package models

type UserRole struct {
	UserID Field
	RoleID Field

	Userlog
}

func (u UserRole) ToFields() []Field {
	return []Field{
		u.UserID,
		u.RoleID,
		u.CreationUserID,
		u.UpdateUserID,
		u.CreatedAt,
		u.UpdatedAt,
	}
}
