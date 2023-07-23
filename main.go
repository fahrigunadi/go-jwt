package main

import (
	"os"

	"github.com/fahrigunadi/go-jwt/controllers"
	"github.com/fahrigunadi/go-jwt/initializers"
	"github.com/fahrigunadi/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run(":" + os.Getenv("PORT"))
}
