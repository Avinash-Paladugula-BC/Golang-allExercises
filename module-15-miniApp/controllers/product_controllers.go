package controllers

import (
	"bc_training_app/config"
	"bc_training_app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	var products []models.Product
	// Assuming Ratings are also imp
	config.DB.Preload("Ratings").Find(&products)
	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	result := config.DB.Preload("Ratings").First(&product, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	// using ShouldBindJSON instead of BindJSON bcz this will return err but doesn't stop execution
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&product)
	c.JSON(http.StatusCreated, product)
}
