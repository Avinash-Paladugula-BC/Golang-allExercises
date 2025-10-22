package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Description string
	Category    string
	Quantity    int
	Price       float64
	Image       string
	Variants    string

	Ratings []Rating `json:"ratings" gorm:"foreignKey:ProductID"`
}

type Rating struct {
	gorm.Model
	ProductID uint
	Name      string
	Review    string
	Rating    int
}

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex" json:"email"`
	Password string
}
