package main

import (
	"fmt"
	"go-jwt/controllers"
	"go-jwt/initializers"
	"go-jwt/middleware"

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
	// r.GET("/validate", controllers.Validate)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run(":3000")
	// r.GET("/signup", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run()

	fmt.Println("hello")
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>")
}
