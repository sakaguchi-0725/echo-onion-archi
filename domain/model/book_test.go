package model_test

import (
	"testing"

	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestNewBook_Success(t *testing.T) {
	bookID := model.GenerateNewBookID()
	title := "The Go Programming Language"
	author := "Alan A. A. Donovan"
	categoryID := uint(1)

	book, err := model.NewBook(bookID, title, author, categoryID, model.Available)

	assert.NoError(t, err)
	assert.Equal(t, bookID, book.ID)
	assert.Equal(t, title, book.Title)
	assert.Equal(t, author, book.Author)
	assert.Equal(t, categoryID, book.CategoryID)
	assert.Equal(t, model.Available, book.Status)
}

func TestNewBook_Failure(t *testing.T) {
	bookID := model.GenerateNewBookID()
	title := ""
	author := "Alan A. A. Donovan"
	categoryID := uint(1)

	book, err := model.NewBook(bookID, title, author, categoryID, model.Available)

	assert.Error(t, err)
	assert.EqualError(t, err, "title is required")
	assert.Equal(t, book, model.Book{})
}

func TestRecreateBook(t *testing.T) {
	bookID := model.GenerateNewBookID()
	title := "The Go Programming Language"
	author := "Alan A. A. Donovan"
	categoryID := uint(1)

	book := model.RecreateBook(bookID, title, author, categoryID, model.Available)

	assert.Equal(t, bookID, book.ID)
	assert.Equal(t, title, book.Title)
	assert.Equal(t, author, book.Author)
	assert.Equal(t, categoryID, book.CategoryID)
	assert.Equal(t, model.Available, book.Status)
}
