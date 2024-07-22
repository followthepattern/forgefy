package models

type Userlog struct {
	CreationUserID Field
	UpdateUserID   Field
	CreatedAt      Field
	UpdatedAt      Field
}

func (u Userlog) ToFields() []Field {
	return []Field{u.CreationUserID, u.UpdateUserID, u.CreatedAt, u.UpdatedAt}
}
