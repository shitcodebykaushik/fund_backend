package routes

import (
	"fund/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/create", controllers.CreateEscrow)
		api.POST("/approve/:id", controllers.ApproveEscrow)
		api.POST("/cancel/:id", controllers.CancelEscrow)
	}
}
