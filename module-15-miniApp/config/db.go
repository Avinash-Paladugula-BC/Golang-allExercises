package config


import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectDB() {
	host := "localhost"
	port := "5432" // 5432 is the default port used by postgres
	user := "postgres"
	password := "Avi@2004"
	dbname := "bc_training_app"

	db_info := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	var err error
	DB, err = gorm.Open(postgres.Open(db_info), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Creating all the required tables in the database
	// Commenting the below lines since the tables are created already
	// err = DB.AutoMigrate(&models.Product{}, &models.Rating{}, &models.User{})
	// Created the User table for authentication, didn't do it first, so created later
	// err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Table creation using AutoMigration failed:", err)
	}

	log.Println("Database connected and migrated successfully!")
}