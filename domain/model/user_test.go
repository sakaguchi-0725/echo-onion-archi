package model_test

import (
	"testing"

	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestNewUser_Success(t *testing.T) {
	id := model.GenerateNewUserID()
	name := "Suzuki"
	email := "test@example.com"
	role := model.Admin

	user, err := model.NewUser(id, name, email, role)

	assert.NoError(t, err)
	assert.Equal(t, id, user.ID)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, role, user.Role)
}

func TestNewUser_Failure(t *testing.T) {
	id := model.GenerateNewUserID()
	name := ""
	email := "test@example.com"
	role := model.Admin

	user, err := model.NewUser(id, name, email, role)

	assert.EqualError(t, err, "name is required")
	assert.Equal(t, model.User{}, user)
}

func TestRecreateUser(t *testing.T) {
	id := model.GenerateNewUserID()
	name := "Suzuki"
	email := "test@example.com"
	role := model.Admin

	user := model.RecreateUser(id, name, email, role)

	assert.Equal(t, id, user.ID)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, role, user.Role)
}
