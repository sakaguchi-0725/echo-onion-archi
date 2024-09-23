package model

import "errors"

type User struct {
	ID       UserID
	Name     string
	Email    string
	Password string
	Role     UserRole
}

func NewUser(id UserID, name, email, password string, role UserRole) (User, error) {
	if name == "" {
		return User{}, errors.New("name is required")
	}

	if len(name) < 3 {
		return User{}, errors.New("name must be 3 or more characters")
	}

	if email == "" {
		return User{}, errors.New("email is required")
	}

	if password == "" {
		return User{}, errors.New("password is required")
	}

	return User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
	}, nil
}

func RecreateUser(id UserID, name, email, password string, role UserRole) User {
	return User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
	}
}
