package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"react-go-store/server/routes"
)

var r = gin.Default()

func main() {
	port := os.Getenv("PORT")

	// C
	r.POST("/orders/create", routes.AddOrders)
	// R
	r.GET("/products", routes.GetProducts)
	r.GET("/products/category", routes.GetProductsByCategory)
	r.GET("/products/category/{id}", routes.ProductByID)
	r.GET("/orders", routes.GetOrders)
	r.GET("/orders/{id}", routes.OrderByID)
	// U
	r.PUT("/orders/{id}/status", routes.UpdateStatusPayment)
	// D
	r.DELETE("/products/category/{id}", routes.DeleteProductByID)
	r.Run(":" + port)
}
