package service

import "github.com/sakaguchi-0725/echo-onion-arch/domain/repository"

type UserService interface {
	IsEmailUnique(email string) (bool, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) IsEmailUnique(email string) (bool, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.Email == email {
		return false, nil
	}

	return true, nil
}
