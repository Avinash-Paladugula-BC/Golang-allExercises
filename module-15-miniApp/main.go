package main

import (
	"log"
	"os"

	"bc_training_app/config"
	"bc_training_app/migrations"  
	"bc_training_app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	// Migration already executed, so need to comment
	migrations.AutoMigrations(config.DB)
	r := gin.Default()
	routes.RegisterRoutes(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server is running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
