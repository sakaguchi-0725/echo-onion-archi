//go:generate mockgen -source=user_repository.go -destination=../../mocks/domain/repository/mock_user_repository.go -package=mocks

package repository

import "github.com/sakaguchi-0725/echo-onion-arch/domain/model"

type UserRepository interface {
	Insert(user model.User, password string) (model.UserID, error)
	FindAll() ([]model.User, error)
	FindByID(id model.UserID) (model.User, error)
	FindByEmail(email string) (model.User, error)
	DeleteByID(id model.UserID) error
}
