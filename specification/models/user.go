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
