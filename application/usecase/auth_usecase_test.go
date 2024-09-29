package usecase_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sakaguchi-0725/echo-onion-arch/application/usecase"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/apperr"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	mocks "github.com/sakaguchi-0725/echo-onion-arch/mocks/domain/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func setupAuthUsecase(t *testing.T) (usecase.AuthUsecase, *mocks.MockUserRepository, *mocks.MockProfileRepository) {
	ctrl := gomock.NewController(t)

	userRepo := mocks.NewMockUserRepository(ctrl)
	profileRepo := mocks.NewMockProfileRepository(ctrl)

	authUsecase := usecase.NewAuthUsecase(userRepo, profileRepo)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	return authUsecase, userRepo, profileRepo
}

func TestAuthUsecase_SignIn_Success(t *testing.T) {
	authUsecase, userRepo, _ := setupAuthUsecase(t)

	email := "test@example.com"
	password := "password"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.NoError(t, err)

	expectedUserID := model.GenerateNewUserID()
	expectedUser := model.User{ID: expectedUserID, Email: email, Password: string(hashedPassword)}

	userRepo.EXPECT().FindByEmail(email).Return(expectedUser, nil)

	userID, err := authUsecase.SignIn(email, password)
	assert.NoError(t, err)
	assert.Equal(t, expectedUserID, userID)
}

func TestAuthUsecase_SignIn_UserNotFound(t *testing.T) {
	authUsecase, userRepo, _ := setupAuthUsecase(t)

	email := "test@example.com"
	password := "password"

	userRepo.EXPECT().FindByEmail(email).Return(model.User{}, apperr.NewApplicationError(apperr.ErrNotFound, fmt.Sprintf("User with email %s not found", email), errors.New("error")))

	_, err := authUsecase.SignIn(email, password)

	require.Error(t, err)

	appErr, ok := err.(*apperr.ApplicationError)

	require.True(t, ok)
	assert.Equal(t, apperr.ErrNotFound, appErr.Code)
	assert.Equal(t, fmt.Sprintf("User with email %s not found", email), appErr.Message)
}

func TestAuthUsecase_SignIn_PasswordMisMatch(t *testing.T) {
	authUsecase, userRepo, _ := setupAuthUsecase(t)

	email := "test@example.com"
	password := "password"

	userRepo.EXPECT().FindByEmail(email).Return(model.User{}, apperr.NewApplicationError(apperr.ErrUnauthorized, "Authentication failed. Please check your email and password.", errors.New("error")))

	_, err := authUsecase.SignIn(email, password)

	require.Error(t, err)

	appErr, ok := err.(*apperr.ApplicationError)
	require.True(t, ok)
	assert.Equal(t, apperr.ErrUnauthorized, appErr.Code)
	assert.Equal(t, "Authentication failed. Please check your email and password.", appErr.Message)
}

func TestAuthUsecase_SignUpForAdmin_Success(t *testing.T) {
	authUsecase, userRepo, profileRepo := setupAuthUsecase(t)

	id := model.GenerateNewUserID()
	name := "John"
	email := "test@example.com"
	password := "password"

	profile := model.Profile{
		UserID: id,
		Name:   name,
		Role:   model.Admin,
	}

	userRepo.EXPECT().Insert(gomock.Any(), email, gomock.Any()).Return(model.GenerateNewUserID(), nil)
	profileRepo.EXPECT().Insert(gomock.Any()).Return(profile, nil)

	userID, err := authUsecase.SignUpForAdmin(name, email, password)

	require.NoError(t, err)
	assert.NotEmpty(t, userID)
}

func TestAuthUsecase_SignUpForAdmin_EmailExists(t *testing.T) {
	authUsecase, userRepo, _ := setupAuthUsecase(t)

	name := "John"
	email := "test@example.com"
	password := "password"

	userRepo.EXPECT().Insert(gomock.Any(), email, gomock.Any()).Return(model.UserID(""), apperr.NewApplicationError(apperr.ErrBadReqeust, "This email address cannot be used", errors.New("error")))

	userID, err := authUsecase.SignUpForAdmin(name, email, password)

	require.Error(t, err)
	assert.Empty(t, userID)

	appErr, ok := err.(*apperr.ApplicationError)
	assert.True(t, ok)
	assert.Equal(t, apperr.ErrBadReqeust, appErr.Code)
	assert.Equal(t, "This email address cannot be used", appErr.Message)
}

func TestAuthUsecase_SignUpForGeneral_Success(t *testing.T) {
	authUsecase, userRepo, profileRepo := setupAuthUsecase(t)

	id := model.GenerateNewUserID()
	name := "John"
	email := "test@example.com"
	password := "password"

	profile := model.Profile{
		UserID: id,
		Name:   name,
		Role:   model.General,
	}

	userRepo.EXPECT().Insert(gomock.Any(), email, gomock.Any()).Return(model.GenerateNewUserID(), nil)
	profileRepo.EXPECT().Insert(gomock.Any()).Return(profile, nil)

	userID, err := authUsecase.SignUpForAdmin(name, email, password)

	require.NoError(t, err)
	assert.NotEmpty(t, userID)
}

func TestAuthUsecase_SignUpForGeneral_EmailExists(t *testing.T) {
	authUsecase, userRepo, _ := setupAuthUsecase(t)

	name := "John"
	email := "test@example.com"
	password := "password"

	userRepo.EXPECT().Insert(gomock.Any(), email, gomock.Any()).Return(model.UserID(""), apperr.NewApplicationError(apperr.ErrBadReqeust, "This email address cannot be used", errors.New("error")))

	userID, err := authUsecase.SignUpForAdmin(name, email, password)

	require.Error(t, err)
	assert.Empty(t, userID)

	appErr, ok := err.(*apperr.ApplicationError)
	assert.True(t, ok)
	assert.Equal(t, apperr.ErrBadReqeust, appErr.Code)
	assert.Equal(t, "This email address cannot be used", appErr.Message)
}
