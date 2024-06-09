package models

type User struct {
	ID        Field
	Email     Field
	FirstName Field
	LastName  Field
	Password  Field

	Userlog
}
