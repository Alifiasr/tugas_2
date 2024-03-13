package router

import (
	"tugas__2/controller"

	"github.com/gin-gonic/gin"
)

func StarOrder(ord *controller.OrderDB) *gin.Engine {
	router := gin.Default()

	router.POST("/orders", ord.CreateOrders)
	router.GET("/orders", ord.GetOrders)
	router.PUT("/orders/:orderId", ord.UpdateOrders)
	router.DELETE("/orders/:orderId", ord.DeleteOrders)

	return router
}
