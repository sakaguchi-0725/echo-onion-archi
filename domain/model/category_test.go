package model_test

import (
	"testing"

	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestNewCategory_Success(t *testing.T) {
	categoryStr := "Programming"
	category, err := model.NewCategory(categoryStr)

	assert.NoError(t, err)
	assert.Equal(t, categoryStr, category.Name)
}

func TestNewCategory_Failure(t *testing.T) {
	category, err := model.NewCategory("")

	assert.EqualError(t, err, "name is required")
	assert.Equal(t, model.Category{}, category)
}
