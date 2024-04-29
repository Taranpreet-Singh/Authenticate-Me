package api

import (
	"authenticate-me/src/models"
	"authenticate-me/src/service"
	"authenticate-me/src/utils/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initial(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"messaage": "Application is running!!",
	})
}

func RegisterUser(c *gin.Context) {
	//* Get request body passed from middleware
	body, exists := c.Get("requestBody")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{constants.Error: constants.ErrUserNotFound.Error()})
		return
	}
	requestBody := body.(models.UserRequestBody)

	//* Register user
	user, customErr := service.RegisterUser(&requestBody)
	if customErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{customErr.ErrorType: customErr.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{constants.Success: "User registered successfully", constants.User: user})
}

func LoginUser(c *gin.Context) {
	//* Get request body passed from middleware
	body, exists := c.Get("requestBody")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{constants.Error: constants.ErrUserNotFound.Error()})
		return
	}
	requestBody := body.(models.UserRequestBody)

	//* Login user
	token, customErr := service.LoginUser(&requestBody)
	if customErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{customErr.ErrorType: customErr.Error.Error()})
		return
	}

	c.JSON(http.StatusAccepted, token)
}
