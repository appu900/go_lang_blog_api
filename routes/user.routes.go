package routes

import (
	"blog-api/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userController *controller.UserController) {
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
}

