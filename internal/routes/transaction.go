package routes

import (
	"trim/internal/handler"
	"trim/internal/middleware"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(rg *gin.RouterGroup) {
	secure := rg.Group("", middleware.AuthMiddleware())
	secure.GET("/transactions", handler.GetTransactions)
	secure.POST("/transactions", handler.CreateTransaction)
}
