package routers

import (
	"test-fase-3/controllers"

	"github.com/gin-gonic/gin"
)

func registerRouter(r *gin.RouterGroup){
	r.POST("", controllers.Register)
}

func loginRouter(r *gin.RouterGroup){
	r.POST("", controllers.Login)
}