package models

// User holds user information
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User // slice holds pointers to users
	nextID = 1
)
