package controllers

import (
	"bc_training_app/config"
	"bc_training_app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetReviews(c *gin.Context) {
	productID := c.Param("id")
	var reviews []models.Rating
	config.DB.Where("product_id = ?", productID).Find(&reviews)
	c.JSON(http.StatusOK, reviews)
}

func CreateReview(c *gin.Context) {
	productID := c.Param("id")
	var review models.Rating

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	review.ProductID = uint(id)

	if review.Rating < 1 || review.Rating > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rating must be between 1 and 5"})
		return
	}

	config.DB.Create(&review)
	c.JSON(http.StatusCreated, review)
}
