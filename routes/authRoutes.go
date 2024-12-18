package routes

import (
	"fund/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Authentication routes
	router.POST("/signup", controllers.SignupHandler)
	router.POST("/login", controllers.LoginHandler)

	// Escrow API routes
	api := router.Group("/api")
	{
		api.POST("/create", controllers.CreateEscrow)
		api.POST("/approve/:id", controllers.ApproveEscrow)
		api.POST("/cancel/:id", controllers.CancelEscrow)
	}
}
