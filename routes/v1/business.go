package v1

import (
	"reqship-api/middleware"
	"reqship-api/services"

	"github.com/gin-gonic/gin"
)

func BusinessRoutes(router *gin.RouterGroup) {
	api := router.Group("business")
	{
		api.GET("", services.GetAllBusinesses)
		api.POST("", middleware.GetUserDetails, services.CreateBusiness)
		api.PUT(":business_id", middleware.GetUserDetails, middleware.GetBusinessDetails, services.UpdateBusiness)
		api.DELETE(":business_id", middleware.GetUserDetails, middleware.GetBusinessDetails, services.DeleteBusiness)
	}
}
