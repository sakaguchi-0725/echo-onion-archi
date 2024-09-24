//go:generate mockgen -source=loan_repository.go -destination=../../mocks/domain/repository/mock_loan_repository.go -package=mocks

package repository

import "github.com/sakaguchi-0725/echo-onion-arch/domain/model"

type LoanFilter struct {
	BookID *model.BookID
	UserID *model.UserID
	Status *model.LoanStatus
}

type LoanRepository interface {
	Insert(loan model.Loan) (model.Loan, error)
	Update(loan model.Loan) (model.Loan, error)
	FindAll() ([]model.Loan, error)
	FindByFilter(filter LoanFilter) ([]model.Loan, error)
	FindByID(id model.LoanID) (model.Loan, error)
	DeleteByID(id model.LoanID) error
}
