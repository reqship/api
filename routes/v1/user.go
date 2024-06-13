package v1

import (
	"reqship-api/middleware"
	"reqship-api/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	api := router.Group("user")
	{
		api.GET(":user_id/orders", middleware.CheckUserAccess, services.GetOrdersByUserId)
	}
}
