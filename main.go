package main

import (
	"assisment2/controllers"
	"assisment2/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDb()
	router := gin.Default()

	router.POST("/orders", controllers.AddOrder)
	router.GET("/orders", controllers.GetOrders)
	router.DELETE("/orders/:orderId", controllers.DeleteOrder)
	router.PUT("/orders/:orderId", controllers.EditOrder)
	router.Run(":8080")
}
