package model_test

import (
	"testing"

	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestNewUserRole_Success(t *testing.T) {
	roleStr := "admin"
	role, err := model.NewUserRole(roleStr)

	assert.NoError(t, err)
	assert.Equal(t, model.Admin, role)
}

func TestNewUserRole_Failure(t *testing.T) {
	roleStr := "invalid"
	role, err := model.NewUserRole(roleStr)

	assert.EqualError(t, err, "invalid user role")
	assert.Equal(t, model.UserRole(""), role)
}

func TestString_UserRole(t *testing.T) {
	roleStr := "admin"
	role, err := model.NewUserRole(roleStr)

	assert.NoError(t, err)
	assert.Equal(t, roleStr, role.String())
}
