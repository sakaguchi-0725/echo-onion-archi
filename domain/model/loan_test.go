package model_test

import (
	"testing"
	"time"

	"github.com/sakaguchi-0725/echo-onion-arch/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestNewLoan(t *testing.T) {
	loanID := model.GenerateNewLoanID()
	bookID := model.GenerateNewBookID()
	userID := model.GenerateNewUserID()

	now := time.Now()
	loanDate := model.NewDate(now)

	loan := model.NewLoan(loanID, bookID, userID, loanDate, nil)

	assert.Equal(t, loanID, loan.ID)
	assert.Equal(t, bookID, loan.BookID)
	assert.Equal(t, userID, loan.UserID)
	assert.Equal(t, loanDate, loan.LoanDate)
}
