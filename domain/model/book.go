package model

import "errors"

type Book struct {
	ID         BookID
	Title      string
	Author     string
	CategoryID uint
	Status     BookStatus
}

func NewBook(id BookID, title, author string, categoryID uint, status BookStatus) (Book, error) {
	if title == "" {
		return Book{}, errors.New("title is required")
	}

	if author == "" {
		return Book{}, errors.New("author is required")
	}

	return Book{
		ID:         id,
		Title:      title,
		Author:     author,
		CategoryID: categoryID,
		Status:     status,
	}, nil
}

func RecreateBook(id BookID, title, author string, categoryID uint, status BookStatus) Book {
	return Book{
		ID:         id,
		Title:      title,
		Author:     author,
		CategoryID: categoryID,
		Status:     status,
	}
}
