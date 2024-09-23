package model

import "errors"

type BookStatus string

const (
	Available BookStatus = "available"
	Loaned    BookStatus = "loaned"
	Reserved  BookStatus = "reserved"
)

func NewBookStatus(s string) (BookStatus, error) {
	switch s {
	case string(Available), string(Loaned), string(Reserved):
		return BookStatus(s), nil
	default:
		return "", errors.New("invalid book status")
	}
}

func (b BookStatus) String() string {
	return string(b)
}
