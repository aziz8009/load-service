package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/aziz8009/load-service/internal/domain"
	"github.com/aziz8009/load-service/internal/repository"
)

type LoanUsecase struct {
	repo *repository.LoanRepository
}

func NewLoanUsecase(repo *repository.LoanRepository) *LoanUsecase {
	return &LoanUsecase{repo: repo}
}

// CreateLoan - state awal: "proposed"
func (uc *LoanUsecase) CreateLoan(loan *domain.Loan) error {
	now := time.Now()

	loan.State = domain.StateProposed
	loan.CreatedAt = now
	loan.UpdatedAt = loan.CreatedAt
	return uc.repo.Create(loan)
}

// ApproveLoan - hanya boleh approve jika masih "proposed"
func (uc *LoanUsecase) ApproveLoan(id uint, approvedBy, approvalPhoto string) error {
	loan, err := uc.repo.GetByID(id)
	if err != nil {
		return err
	}
	if loan.State != domain.StateProposed {
		return errors.New("loan is not in proposed state")
	}

	now := time.Now()

	loan.State = domain.StateApproved
	loan.ApprovedBy = approvedBy
	loan.ApprovalPhoto = approvalPhoto
	loan.ApprovalDate = &now
	loan.UpdatedAt = now

	return uc.repo.Update(loan)
}

// DisburseLoan - hanya jika loan sudah "approved"
func (uc *LoanUsecase) DisburseLoan(id uint, disbursedBy, file string, agreementSigned string) error {
	loan, err := uc.repo.GetByID(id)
	if err != nil {
		return err
	}
	if loan.State != domain.StateApproved {
		return errors.New("loan is not in approved state")
	}

	now := time.Now()

	loan.State = domain.StateDisbursed
	loan.DisbursedBy = disbursedBy
	loan.AgreementSigned = agreementSigned
	loan.DisbursementDate = &now
	loan.AgreementLetter = file
	loan.UpdatedAt = *loan.DisbursementDate

	return uc.repo.Update(loan)
}

func (uc *LoanUsecase) InvestLoan(id uint, investorID string) error {
	loan, err := uc.repo.GetByID(id)
	if err != nil {
		return err
	}

	if loan.State != "approved" {
		return fmt.Errorf("loan is not in approved status")
	}

	now := time.Now()
	loan.State = "invested"
	loan.BorrowerID = investorID
	loan.CreatedAt = now

	return uc.repo.Update(loan)
}

func (uc *LoanUsecase) GetByID(id uint) (*domain.Loan, error) {
	return uc.repo.GetByID(id)
}

func (uc *LoanUsecase) GetAll() ([]*domain.Loan, error) {
	return uc.repo.GetAll()
}
