package domain

import "time"

type LoanState string

const (
	StateProposed  LoanState = "proposed"
	StateApproved  LoanState = "approved"
	StateInvested  LoanState = "invested"
	StateDisbursed LoanState = "disbursed"
)

type Loan struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	BorrowerID      string    `gorm:"size:100;not null"`
	PrincipalAmount float64   `gorm:"not null"`
	Rate            float64   `gorm:"not null"`
	ROI             float64   `gorm:"not null"`
	State           LoanState `gorm:"type:varchar(20);default:'proposed'"`

	AgreementLetter string `gorm:"size:255"`

	ApprovedBy    string `gorm:"size:100"`
	ApprovalPhoto string `gorm:"size:255"`
	ApprovalDate  *time.Time

	DisbursedBy      string `gorm:"size:100"`
	DisbursementDate *time.Time
	AgreementSigned  string `gorm:"size:255"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
