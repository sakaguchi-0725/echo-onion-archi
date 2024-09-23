package model

import (
	"errors"

	"github.com/google/uuid"
)

type UserID string

func NewUserID(s string) (UserID, error) {
	if _, err := uuid.Parse(s); err != nil {
		return "", errors.New("invalid user ID format")
	}
	return UserID(s), nil
}

func GenerateNewUserID() UserID {
	return UserID(uuid.New().String())
}

func (u UserID) String() string {
	return string(u)
}
