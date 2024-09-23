package repository

import "github.com/sakaguchi-0725/echo-onion-arch/domain/model"

type UserRepository interface {
	Insert(user model.User) (model.UserID, error)
	FindAll() ([]model.User, error)
	FindByID(id model.UserID) (model.User, error)
	FindByEmail(email string) (model.User, error)
	DeleteByID(id model.User) error
}
