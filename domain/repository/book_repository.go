package repository

import "github.com/sakaguchi-0725/echo-onion-arch/domain/model"

type BookFilter struct {
	Title    *string
	Author   *string
	Category *uint
	Status   *model.BookStatus
}

type BookRepository interface {
	Insert(book model.Book) (model.BookID, error)
	Update(book model.Book) (model.BookID, error)
	FindAll(filter BookFilter) ([]model.Book, error)
	FindByFilter(title, author *string)
}
