package api

import (
	handler "authenticate-me/src/api/handlers"
	middleware "authenticate-me/src/api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", handler.Initial)

	router.POST("/api/register", middleware.ValidateUserBody(), handler.RegisterUser)
	router.POST("/api/login", middleware.ValidateUserBody(), handler.LoginUser)
}
