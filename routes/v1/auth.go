package v1

import (
	"reqship-api/services"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	api := router.Group("auth")
	{
		api.POST("login", services.Login)
		api.POST("signup", services.SignUp)
		api.GET(":username", services.CheckAuth)
	}
}
