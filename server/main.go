package main

import (
	"os"
	"react-go-store/server/store"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	r.GET("/api/products", store.AllOrders)
	r.GET("/api/products/{id}", store.ProductByID)
	r.POST("/api/orders", store.CreateOrders)
	r.GET("/api/orders", store.AllOrders)
	r.GET("/api/orders/{id}", store.OrderByID)
	r.PUT("/api/orders/{id}/status", store.UpdateStatusPayment)
	r.Run(":" + port)	
}

