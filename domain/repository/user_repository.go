package repository

import "github.com/sakaguchi-0725/echo-onion-arch/domain/model"

type UserRepository interface {
	Insert(id model.UserID, email string, password string) (model.UserID, error)
	FindByEmail(email string) (model.UserID, error)
}
