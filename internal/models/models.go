package models

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	Active    bool
}
