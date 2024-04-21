package initializers

import "authenticate-me/src/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
