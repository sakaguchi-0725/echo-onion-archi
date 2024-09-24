//go:generate mockgen -source=book_repository.go -destination=../../mocks/domain/repository/mock_book_repository.go -package=mocks

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
