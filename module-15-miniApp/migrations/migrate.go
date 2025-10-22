package migrations


import (
	"bc_training_app/models"

	"gorm.io/gorm"
)

func AutoMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.Product{}, &models.Rating{}, &models.User{})
	// db.AutoMigrate(&models.User{})
}
