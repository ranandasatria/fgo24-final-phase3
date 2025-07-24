package routers

import (
	"test-fase-3/controllers"
	"test-fase-3/middlewares"

	"github.com/gin-gonic/gin"
)

func transactionsRouter(r *gin.RouterGroup) {
	r.Use(middlewares.VerifyToken())
	r.GET("", controllers.GetTransactions)
	r.POST("", controllers.CreateTransaction)
}
