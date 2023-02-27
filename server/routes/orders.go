package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var orderCollection *mongo.Collection = OpenCollection(Client, "orders")

func AddOrders(c *gin.Context) {

}

func GetProducts(c *gin.Context) {

}
func GetProductsByCategory(c *gin.Context) {

}
func ProductByID(c *gin.Context) {

}
func GetOrders(c *gin.Context) {

}
func OrderByID(c *gin.Context) {

}
func UpdateStatusPayment(c *gin.Context) {

}
func DeleteProductByID(c *gin.Context) {

}
