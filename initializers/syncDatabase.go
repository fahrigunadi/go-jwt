package initializers

import "github.com/fahrigunadi/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
