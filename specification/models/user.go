package models

type User struct {
	ID        Field
	Email     Field
	FirstName Field
	LastName  Field
	Password  Field
	Active    Field

	Userlog
}

func (u User) ToFields() []Field {
	return []Field{
		u.ID,
		u.Email,
		u.FirstName,
		u.LastName,
		u.Password,
		u.CreationUserID,
		u.UpdateUserID,
		u.CreatedAt,
		u.UpdatedAt,
	}
}
