package api

import (
	"authenticate-me/src/models"
	"authenticate-me/src/utils/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticateUser(c *gin.Context) {

}

func ValidateUserBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get email and pass from request body
		var body models.UserRequestBody
		if c.Bind(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				constants.Error: "Failed to load body",
			})
			return
		}
		c.Set("requestBody", body)

		c.Next()
	}
}
