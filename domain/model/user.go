package model

import "errors"

type User struct {
	ID    UserID
	Name  string
	Email string
	Role  UserRole
}

func NewUser(id UserID, name, email string, role UserRole) (User, error) {
	if name == "" {
		return User{}, errors.New("name is required")
	}

	if len(name) < 3 {
		return User{}, errors.New("name must be 3 or more characters")
	}

	if email == "" {
		return User{}, errors.New("email is required")
	}

	return User{
		ID:    id,
		Name:  name,
		Email: email,
		Role:  role,
	}, nil
}

func RecreateUser(id UserID, name, email string, role UserRole) User {
	return User{
		ID:    id,
		Name:  name,
		Email: email,
		Role:  role,
	}
}
