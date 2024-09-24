package model_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestNewBookID_Success(t *testing.T) {
	validUUID := uuid.New().String()
	bookID, err := model.NewBookID(validUUID)

	assert.NoError(t, err)
	assert.Equal(t, model.BookID(validUUID), bookID)
}

func TestNewBookID_Failure(t *testing.T) {
	invalidUUID := "invalid_uuid"
	bookID, err := model.NewBookID(invalidUUID)

	assert.Error(t, err)
	assert.EqualError(t, err, "invalid book ID format")
	assert.Equal(t, model.BookID(""), bookID)
}

func TestGenerateNewBookID(t *testing.T) {
	bookID := model.GenerateNewBookID()

	parsedUUID, err := uuid.Parse(bookID.String())
	assert.NoError(t, err)
	assert.Equal(t, bookID, model.BookID(parsedUUID.String()))
}

func TestString_BookID(t *testing.T) {
	uuidStr := uuid.New().String()
	bookID, err := model.NewBookID(uuidStr)

	assert.NoError(t, err)
	assert.Equal(t, uuidStr, bookID.String())
}
