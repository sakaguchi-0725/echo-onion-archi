package model

import "errors"

type Profile struct {
	UserID UserID
	Name   string
	Role   UserRole
}

func NewProfile(user_id UserID, name string, role UserRole) (Profile, error) {
	if name == "" {
		return Profile{}, errors.New("name is required")
	}
	if len(name) < 3 {
		return Profile{}, errors.New("name must be 3 or more characters")
	}

	return Profile{
		UserID: user_id,
		Name:   name,
		Role:   role,
	}, nil
}

func RecreateProfile(user_id UserID, name string, role UserRole) Profile {
	return Profile{
		UserID: user_id,
		Name:   name,
		Role:   role,
	}
}
