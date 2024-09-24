package model_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestNewLoanID_Success(t *testing.T) {
	validUUID := uuid.New().String()
	loanID, err := model.NewLoanID(validUUID)

	assert.NoError(t, err)
	assert.Equal(t, model.LoanID(validUUID), loanID)
}

func TestNewLoanID_Failure(t *testing.T) {
	validUUID := "invalid_uuid"
	loanID, err := model.NewLoanID(validUUID)

	assert.Error(t, err)
	assert.EqualError(t, err, "invalid loan ID format")
	assert.Equal(t, model.LoanID(""), loanID)
}

func TestGenerateNewLoanID(t *testing.T) {
	loanID := model.GenerateNewLoanID()

	parsedUUID, err := uuid.Parse(loanID.String())
	assert.NoError(t, err)
	assert.Equal(t, loanID, model.LoanID(parsedUUID.String()))
}

func TestString_LoanID(t *testing.T) {
	uuidStr := uuid.New().String()
	loanID, err := model.NewLoanID(uuidStr)

	assert.NoError(t, err)
	assert.Equal(t, uuidStr, loanID.String())
}
