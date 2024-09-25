package repository

import "github.com/sakaguchi-0725/echo-onion-arch/domain/model"

type UserRepository interface {
	SaveCredentials(email string, password string) (model.UserID, error)
	FindByEmail(email string) (model.UserID, error)
}
