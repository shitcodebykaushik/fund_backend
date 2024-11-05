package routes

import (
	"fund/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/signup", controllers.SignupHandler)
	router.POST("/login", controllers.LoginHandler)
}
