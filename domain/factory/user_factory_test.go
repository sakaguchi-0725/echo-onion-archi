package factory_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/factory"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	mocks "github.com/sakaguchi-0725/echo-onion-arch/mocks/domain/service"
	"github.com/stretchr/testify/assert"
)

func setupUserServiceMock(t *testing.T) (*gomock.Controller, *mocks.MockUserService) {
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockUserService(ctrl)
	return ctrl, mockRepo
}

func TestUserFactory_CreateNewUser_Success(t *testing.T) {
	ctrl, mockUserService := setupUserServiceMock(t)
	defer ctrl.Finish()

	mockUserService.EXPECT().IsEmailUnique("john@example.com").Return(true, nil)

	userFactory := factory.NewUserFactory(mockUserService)

	user, err := userFactory.CreateNewUser("John Doe", "john@example.com", "admin")

	assert.NoError(t, err)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@example.com", user.Email)
	assert.Equal(t, model.Admin, user.Role)
}

func TestUserFactory_CreateNewUser_EmailAlreadyExists(t *testing.T) {
	ctrl, mockUserService := setupUserServiceMock(t)
	defer ctrl.Finish()

	mockUserService.EXPECT().IsEmailUnique("john@example.com").Return(false, nil)

	userFactory := factory.NewUserFactory(mockUserService)

	user, err := userFactory.CreateNewUser("John Doe", "john@example.com", "admin")

	assert.Error(t, err)
	assert.EqualError(t, err, "email already exists")
	assert.Equal(t, model.User{}, user)
}

func TestUserFactory_CreateNewUser_InvalidUserRole(t *testing.T) {
	ctrl, mockUserService := setupUserServiceMock(t)
	defer ctrl.Finish()

	mockUserService.EXPECT().IsEmailUnique("john@example.com").Return(true, nil)

	userFactory := factory.NewUserFactory(mockUserService)

	user, err := userFactory.CreateNewUser("John Doe", "john@example.com", "invalid_role")

	assert.Error(t, err)
	assert.EqualError(t, err, "invalid user role")
	assert.Equal(t, model.User{}, user)
}

func TestUserFactory_CreateNewUser_UserServiceError(t *testing.T) {
	ctrl, mockUserService := setupUserServiceMock(t)
	defer ctrl.Finish()

	mockUserService.EXPECT().IsEmailUnique("john@example.com").Return(false, errors.New("user service error"))

	userFactory := factory.NewUserFactory(mockUserService)

	user, err := userFactory.CreateNewUser("John Doe", "john@example.com", "admin")

	assert.Error(t, err)
	assert.EqualError(t, err, "user service error")
	assert.Equal(t, model.User{}, user)
}
