package http

import (
	"github.com/aziz8009/load-service/config"
	"github.com/aziz8009/load-service/internal/handler"
	"github.com/aziz8009/load-service/internal/repository"
	"github.com/aziz8009/load-service/internal/usecase"
	"github.com/aziz8009/load-service/pkg/database"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Load config
	cfg := config.LoadConfig()

	// Init DB dengan config
	db := database.InitDB(cfg)

	// Dependency Injection
	loanRepo := repository.NewLoanRepository(db)
	loanUC := usecase.NewLoanUsecase(loanRepo)
	loanHandler := handler.NewLoanHandler(loanUC)

	// Routes
	loanRoutes := r.Group("/loans")
	{
		loanRoutes.GET("/", loanHandler.GetAllLoans)
		loanRoutes.GET("/:id", loanHandler.GetLoanDetail)
		loanRoutes.POST("", loanHandler.CreateLoan)
		loanRoutes.PUT("/:id/approve", loanHandler.ApproveLoan)
		loanRoutes.PUT("/:id/disburse", loanHandler.DisburseLoan)
		loanRoutes.PUT("/:id/invest", loanHandler.InvestLoan)
	}

	return r
}
