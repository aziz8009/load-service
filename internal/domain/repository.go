package domain

type LoanRepositoryInterface interface {
	Create(loan *Loan) error
	GetByID(id uint) (*Loan, error)
	Update(loan *Loan) error
}
