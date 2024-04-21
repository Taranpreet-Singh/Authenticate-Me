package main

import (
	"authenticate-me/src/api"
	"authenticate-me/src/initializers"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()

	api.SetupRoutes(router)
	router.Run(":" + os.Getenv("SERVER_PORT"))
}
