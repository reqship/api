package routes

import (
	v1 "reqship-api/routes/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.Use(cors.Default())

	router_v1 := router.Group("api/v1")
	{
		v1.AuthRoutes(router_v1)
		v1.BusinessRoutes(router_v1)
		v1.ItemsRoutes(router_v1)
	}

	router.Run(":8080")
}
