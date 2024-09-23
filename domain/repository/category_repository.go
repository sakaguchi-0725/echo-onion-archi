package repository

import "github.com/sakaguchi-0725/echo-onion-arch/domain/model"

type CategoryRepository interface {
	Insert(category model.Category) (uint, error)
	Update(category model.Category) (uint, error)
	FindAll() ([]model.Category, error)
	DeleteByID(id uint) error
}
