//go:generate mockgen -source=auth_usecase.go -destination=../../mocks/application/usecase/mock_auth_usecase.go -package=mocks
package usecase

import (
	"github.com/sakaguchi-0725/echo-onion-arch/domain/apperr"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	SignIn(email, password string) (model.UserID, error)
	SignUpForAdmin(name, email, password string) (model.UserID, error)
	SignUpForGeneral(name, email, password string) (model.UserID, error)
}

type authUsecase struct {
	userRepo    repository.UserRepository
	profileRepo repository.ProfileRepository
}

func NewAuthUsecase(userRepo repository.UserRepository, profileRepo repository.ProfileRepository) AuthUsecase {
	return &authUsecase{userRepo, profileRepo}
}

func (a *authUsecase) SignIn(email string, password string) (model.UserID, error) {
	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return model.UserID(""), err
	}

	if err := a.compareHashPassword(user.Password, password); err != nil {
		return model.UserID(""), apperr.NewApplicationError(apperr.ErrUnauthorized, "Authentication failed. Please check your email and password.", err)
	}

	return user.ID, nil
}

func (a *authUsecase) SignUpForAdmin(name, email, password string) (model.UserID, error) {
	userID := model.GenerateNewUserID()

	hashedPassword, err := a.hashPassword(password)
	if err != nil {
		return model.UserID(""), apperr.NewApplicationError(apperr.ErrInternalError, "Failed to process the pasword. Please try again.", err)
	}

	createdUserID, err := a.userRepo.Insert(userID, email, hashedPassword)
	if err != nil {
		return model.UserID(""), err
	}

	profile, err := model.NewProfile(createdUserID, name, model.Admin)
	if err != nil {
		return model.UserID(""), apperr.NewApplicationError(apperr.ErrBadReqeust, "Invalid profile information provided. Please check your input.", err)
	}

	_, err = a.profileRepo.Insert(profile)
	if err != nil {
		return model.UserID(""), err
	}

	return createdUserID, nil
}

func (a *authUsecase) SignUpForGeneral(name, email, password string) (model.UserID, error) {
	userID := model.GenerateNewUserID()

	hashedPassword, err := a.hashPassword(password)
	if err != nil {
		return model.UserID(""), apperr.NewApplicationError(apperr.ErrInternalError, "Failed to process the pasword. Please try again.", err)
	}

	createdUserID, err := a.userRepo.Insert(userID, email, hashedPassword)
	if err != nil {
		return model.UserID(""), err
	}

	profile, err := model.NewProfile(createdUserID, name, model.General)
	if err != nil {
		return model.UserID(""), apperr.NewApplicationError(apperr.ErrBadReqeust, "Invalid profile information provided. Please check your input.", err)
	}

	_, err = a.profileRepo.Insert(profile)
	if err != nil {
		return model.UserID(""), err
	}

	return createdUserID, nil
}

func (a *authUsecase) hashPassword(s string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (a *authUsecase) compareHashPassword(hashedPassword, requestPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(requestPassword))
	if err != nil {
		return err
	}

	return nil
}
