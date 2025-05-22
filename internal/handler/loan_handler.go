package handler

import (
	"net/http"
	"strconv"

	"github.com/aziz8009/load-service/internal/domain"
	"github.com/aziz8009/load-service/internal/usecase"
	"github.com/gin-gonic/gin"
)

type LoanHandler struct {
	usecase *usecase.LoanUsecase
}

type InvestLoanRequest struct {
	InvestorID string `json:"investor_id"`
}

func NewLoanHandler(uc *usecase.LoanUsecase) *LoanHandler {
	return &LoanHandler{usecase: uc}
}

// POST /loans
func (h *LoanHandler) CreateLoan(c *gin.Context) {
	var req domain.Loan
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.CreateLoan(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Loan created", "loan": req})
}

// PUT /loans/:id/approve
func (h *LoanHandler) ApproveLoan(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan ID"})
		return
	}

	var req struct {
		ApprovedBy    string `json:"approved_by"`
		ApprovalPhoto string `json:"approval_photo"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.usecase.ApproveLoan(uint(id), req.ApprovedBy, req.ApprovalPhoto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan approved"})
}

// PUT /loans/:id/disburse
func (h *LoanHandler) DisburseLoan(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan ID"})
		return
	}

	var req struct {
		DisbursedBy     string `json:"disbursed_by"`
		AgreementSigned string `json:"agreement_signed"`
		AgreementFile   string `json:"agreement_file"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.usecase.DisburseLoan(uint(id), req.DisbursedBy, req.AgreementFile, req.AgreementSigned)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan disbursed"})
}

func (h *LoanHandler) InvestLoan(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req InvestLoanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.usecase.InvestLoan(uint(id), req.InvestorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan invested"})
}

func (h *LoanHandler) GetAllLoans(c *gin.Context) {
	loans, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, loans)
}

func (h *LoanHandler) GetLoanDetail(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan id"})
		return
	}

	loan, err := h.usecase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "loan not found"})
		return
	}
	c.JSON(http.StatusOK, loan)
}
