package v1

import (
	"reqship-api/middleware"
	"reqship-api/services"

	"github.com/gin-gonic/gin"
)

func ItemsRoutes(router *gin.RouterGroup) {
	api := router.Group("items")
	{
		api.POST("", middleware.GetBusinessDetails, services.CreateItem)
		api.GET(":business_id", services.GetItemsByBusiness)
	}
}
