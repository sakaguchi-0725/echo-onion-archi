package model

import (
	"errors"

	"github.com/google/uuid"
)

type LoanID string

func NewLoanID(s string) (LoanID, error) {
	if _, err := uuid.Parse(s); err != nil {
		return "", errors.New("invalid loan ID format")
	}
	return LoanID(s), nil
}

func GenerateNewLoanID() LoanID {
	return LoanID(uuid.New().String())
}

func (u LoanID) String() string {
	return string(u)
}
