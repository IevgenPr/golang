package models

import (
	"errors"
	"fmt"
)

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

// GetUsers returns all users
func GetUsers() []*User {
	return users
}

// AddUser Adds user to a slice
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New user must not include ID" +
			"or it must be set to zero.")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// GetUserByID returns user by ID
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	// return User{}, errors.New("No user found by id:" + string(id))
	return User{}, fmt.Errorf("No user found by id: %d", id)
}

// RemoveUserByID returns user by ID
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	// return User{}, errors.New("No user found by id:" + string(id))
	return fmt.Errorf("No user found by id: %d", id)
}

// UpdateUser updates user with ID
func UpdateUser(us User) error {
	for _, u := range users {
		if u.ID == us.ID {
			u.FirstName = us.FirstName
			u.LastName = us.LastName
			return nil
		}
	}
	// return User{}, errors.New("No user found by id:" + string(id))
	return fmt.Errorf("No user found")
}
