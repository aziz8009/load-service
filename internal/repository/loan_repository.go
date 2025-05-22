package repository

import (
	"github.com/aziz8009/load-service/internal/domain"
	"gorm.io/gorm"
)

type LoanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) *LoanRepository {
	return &LoanRepository{db: db}
}

func (r *LoanRepository) Create(loan *domain.Loan) error {
	return r.db.Create(loan).Error
}

func (r *LoanRepository) GetByID(id uint) (*domain.Loan, error) {
	var loan domain.Loan
	err := r.db.First(&loan, id).Error
	return &loan, err
}

func (r *LoanRepository) GetAll() ([]*domain.Loan, error) {
	var loans []*domain.Loan
	if err := r.db.Preload("Investors").Find(&loans).Error; err != nil {
		return nil, err
	}
	return loans, nil
}

func (r *LoanRepository) Update(loan *domain.Loan) error {
	return r.db.Save(loan).Error
}
