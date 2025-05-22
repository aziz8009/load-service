package model

import "time"

type Loan struct {
	ID            uint `gorm:"primaryKey"`
	BorrowerID    string
	Principal     float64
	Rate          float64
	ROI           float64
	State         string
	AgreementLink string
	ApprovedAt    *time.Time
	ApprovedBy    string
	VisitedProof  string
	DisbursedAt   *time.Time
	DisbursedBy   string
	DisburseProof string
	Investments   []Investment
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Investment struct {
	ID       uint `gorm:"primaryKey"`
	LoanID   uint
	Investor string
	Amount   float64
}
