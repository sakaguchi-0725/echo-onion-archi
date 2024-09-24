package model

type Loan struct {
	ID         LoanID
	BookID     BookID
	UserID     UserID
	LoanDate   Date
	ReturnDate *Date
	Status     LoanStatus
}

func NewLoan(id LoanID, bookID BookID, userID UserID, loanDate Date, returnDate *Date) Loan {
	return Loan{
		ID:         id,
		BookID:     bookID,
		UserID:     userID,
		LoanDate:   loanDate,
		ReturnDate: returnDate,
	}
}
