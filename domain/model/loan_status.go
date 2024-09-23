package model

import "errors"

type LoanStatus string

const (
	Borrowed LoanStatus = "borrowed"
	Returned LoanStatus = "returned"
)

func NewLoanStatus(s string) (LoanStatus, error) {
	switch s {
	case string(Borrowed), string(Returned):
		return LoanStatus(s), nil
	default:
		return "", errors.New("invalid loan status")
	}
}

func (l LoanStatus) String() string {
	return string(l)
}
