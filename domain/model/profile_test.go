package model_test

import (
	"testing"

	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestNewProfile_Success(t *testing.T) {
	userID := model.UserID("user-123")
	name := "John Doe"
	role := model.UserRole("Admin")

	profile, err := model.NewProfile(userID, name, role)

	assert.NoError(t, err)
	assert.Equal(t, userID, profile.UserID)
	assert.Equal(t, name, profile.Name)
	assert.Equal(t, role, profile.Role)
}

func TestNewProfile_NameRequired(t *testing.T) {
	userID := model.UserID("user-123")
	name := ""
	role := model.UserRole("Admin")

	profile, err := model.NewProfile(userID, name, role)

	assert.Error(t, err)
	assert.EqualError(t, err, "name is required")
	assert.Equal(t, model.Profile{}, profile)
}

func TestNewProfile_NameLength(t *testing.T) {
	userID := model.UserID("user-123")
	name := "Jo" // Name is too short
	role := model.UserRole("Admin")

	profile, err := model.NewProfile(userID, name, role)

	assert.Error(t, err)
	assert.EqualError(t, err, "name must be 3 or more characters")
	assert.Equal(t, model.Profile{}, profile)
}

func TestRecreateProfile(t *testing.T) {
	userID := model.UserID("user-123")
	name := "John Doe"
	role := model.UserRole("Admin")

	profile := model.RecreateProfile(userID, name, role)

	assert.Equal(t, userID, profile.UserID)
	assert.Equal(t, name, profile.Name)
	assert.Equal(t, role, profile.Role)
}
