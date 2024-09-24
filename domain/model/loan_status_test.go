package model_test

import (
	"testing"

	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestNewLoanStatus_Success(t *testing.T) {
	validStatus := "borrowed"
	status, err := model.NewLoanStatus(validStatus)

	assert.NoError(t, err)
	assert.Equal(t, model.Borrowed, status)
}

func TestNewLoanStatus_Failure(t *testing.T) {
	invalidStatus := "invalid_status"
	status, err := model.NewLoanStatus(invalidStatus)

	assert.EqualError(t, err, "invalid loan status")
	assert.Equal(t, model.LoanStatus(""), status)
}

func TestString_LoanStatus(t *testing.T) {
	statusStr := "borrowed"
	status, err := model.NewLoanStatus(statusStr)

	assert.NoError(t, err)
	assert.Equal(t, statusStr, status.String())
}
