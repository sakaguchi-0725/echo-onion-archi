//go:generate mockgen -source=user_factory.go -destination=../../mocks/domain/factory/mock_user_factory.go -package=mocks

package factory

import (
	"errors"

	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/service"
)

type UserFactory interface {
	CreateNewUser(name, email, role string) (model.User, error)
}

type userFactory struct {
	userService service.UserService
}

func NewUserFactory(userService service.UserService) UserFactory {
	return &userFactory{userService}
}

func (f *userFactory) CreateNewUser(name, email, role string) (model.User, error) {
	id := model.GenerateNewUserID()

	isUnique, err := f.userService.IsEmailUnique(email)
	if err != nil {
		return model.User{}, err
	}

	if !isUnique {
		return model.User{}, errors.New("email already exists")
	}

	userRole, err := model.NewUserRole(role)
	if err != nil {
		return model.User{}, err
	}

	user, err := model.NewUser(id, name, email, userRole)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
