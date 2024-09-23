package model

import (
	"errors"

	"github.com/google/uuid"
)

type BookID string

func NewBookID(s string) (BookID, error) {
	if _, err := uuid.Parse(s); err != nil {
		return "", errors.New("invalid book ID format")
	}
	return BookID(s), nil
}

func GenerateNewBookID() BookID {
	return BookID(uuid.New().String())
}

func (u BookID) String() string {
	return string(u)
}
