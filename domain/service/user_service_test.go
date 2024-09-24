package service_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/service"
	mocks "github.com/sakaguchi-0725/echo-onion-arch/mocks/domain/repository"
	"github.com/stretchr/testify/assert"
)

func setupUserRepositoryMock(t *testing.T) (*gomock.Controller, *mocks.MockUserRepository) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockUserRepository(ctrl)
	return ctrl, mockRepo
}

func TestIsEmailUnique_Success(t *testing.T) {
	ctrl, mockRepo := setupUserRepositoryMock(t)
	defer ctrl.Finish()

	mockRepo.EXPECT().FindByEmail("test@example.com").Return(model.User{}, nil)

	userService := service.NewUserService(mockRepo)

	isUnique, err := userService.IsEmailUnique("test@example.com")

	assert.NoError(t, err)
	assert.True(t, isUnique)
}

func TestIsEmailUnique_EmailExists(t *testing.T) {
	ctrl, mockRepo := setupUserRepositoryMock(t)
	defer ctrl.Finish()

	mockRepo.EXPECT().FindByEmail("test@example.com").Return(model.User{Email: "test@example.com"}, nil)

	userService := service.NewUserService(mockRepo)

	isUnique, err := userService.IsEmailUnique("test@example.com")

	assert.NoError(t, err)
	assert.False(t, isUnique)
}

func TestIsEmailUnique_Error(t *testing.T) {
	ctrl, mockRepo := setupUserRepositoryMock(t)
	defer ctrl.Finish()

	mockRepo.EXPECT().FindByEmail("test@example.com").Return(model.User{}, errors.New("database error"))

	userService := service.NewUserService(mockRepo)

	isUnique, err := userService.IsEmailUnique("test@example.com")

	assert.Error(t, err)
	assert.False(t, isUnique)
}
