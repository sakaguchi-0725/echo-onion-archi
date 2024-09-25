//go:generate mockgen -source=profile_repository.go -destination=../../mocks/domain/repository/mock_profile_repository.go -package=mocks

package repository

import "github.com/sakaguchi-0725/echo-onion-arch/domain/model"

type ProfileRepository interface {
	Insert(user model.Profile, password string) (model.Profile, error)
	FindAll() ([]model.Profile, error)
	FindByID(id model.UserID) (model.Profile, error)
	FindByEmail(email string) (model.Profile, error)
	DeleteByID(id model.UserID) error
}
