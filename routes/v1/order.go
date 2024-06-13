package v1

import (
	"reqship-api/middleware"
	"reqship-api/services"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.RouterGroup) {
	api := router.Group("order")
	{
		api.POST("", middleware.GetUserDetails, services.NewOrder)
		api.GET(":order_id", middleware.GetUserDetails, services.GetOrderContent)
		api.POST(":order_id", middleware.GetBusinessDetails, services.CompleteOrder)
	}
}
