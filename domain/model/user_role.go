package model

import "errors"

type UserRole string

const (
	Admin   UserRole = "admin"
	General UserRole = "general"
)

func NewUserRole(s string) (UserRole, error) {
	switch s {
	case string(Admin), string(General):
		return UserRole(s), nil
	default:
		return "", errors.New("invalid user role")
	}
}

func (u UserRole) String() string {
	return string(u)
}
