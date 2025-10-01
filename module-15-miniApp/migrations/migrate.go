package migrations


import (
	"bc_training_app/models"

	"gorm.io/gorm"
)

func AutoMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.Product{}, &models.Rating{})
	db.AutoMigrate(&models.User{})
}
