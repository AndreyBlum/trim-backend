package routes

import (
	authhandler "trim/internal/handler/auth"
	"trim/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/register", authhandler.Register)
	rg.POST("/login", authhandler.Login)
	rg.GET("/me", middleware.AuthMiddleware(), authhandler.Profile)
}
