package model_test

import (
	"testing"

	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestNewBookStatus_Success(t *testing.T) {
	validStatus := "available"
	bookStatus, err := model.NewBookStatus(validStatus)

	assert.NoError(t, err)
	assert.Equal(t, model.Available, bookStatus)
}

func TestNewBookStatus_Failure(t *testing.T) {
	invalidStatus := "example"
	bookStatus, err := model.NewBookStatus(invalidStatus)

	assert.Error(t, err)
	assert.EqualError(t, err, "invalid book status")
	assert.Equal(t, model.BookStatus(""), bookStatus)
}

func TestString_BookStatus(t *testing.T) {
	statusStr := "loaned"
	bookStatus, err := model.NewBookStatus(statusStr)

	assert.NoError(t, err)
	assert.Equal(t, statusStr, bookStatus.String())
}
