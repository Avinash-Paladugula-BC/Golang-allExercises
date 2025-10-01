package routes


import (
	"bc_training_app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/api/products", controllers.GetAllProducts)
	r.GET("/api/products/:id", controllers.GetProduct)
	r.POST("/api/products/create", controllers.CreateProduct)
	r.GET("/api/products/:id/reviews", controllers.GetReviews)
	r.POST("/api/products/:id/reviews/create", controllers.CreateReview)
}
