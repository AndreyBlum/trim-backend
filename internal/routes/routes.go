package routes

import "github.com/gin-gonic/gin"

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		AuthRoutes(api)
		TransactionRoutes(api)
	}
	return r
}
