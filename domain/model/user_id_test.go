package model_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestNewUserID_Success(t *testing.T) {
	validID := uuid.New().String()
	userID, err := model.NewUserID(validID)

	assert.NoError(t, err)
	assert.Equal(t, model.UserID(validID), userID)
}

func TestNewUserID_Failure(t *testing.T) {
	invalidID := "invalid_uuid"
	userID, err := model.NewUserID(invalidID)

	assert.EqualError(t, err, "invalid user ID format")
	assert.Equal(t, model.UserID(""), userID)
}

func TestGenerateNewUserID(t *testing.T) {
	userID := model.GenerateNewUserID()

	parsedUUID, err := uuid.Parse(userID.String())
	assert.NoError(t, err)
	assert.Equal(t, userID, model.UserID(parsedUUID.String()))
}

func TestString_UserID(t *testing.T) {
	uuidStr := uuid.New().String()
	userID, err := model.NewUserID(uuidStr)

	assert.NoError(t, err)
	assert.Equal(t, uuidStr, userID.String())
}
