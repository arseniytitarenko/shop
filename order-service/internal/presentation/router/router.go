package router

import (
	"github.com/gin-gonic/gin"
	"order/internal/presentation/handler"
)

func SetupRouter(orderHandler *handler.OrderHandler) *gin.Engine {
	r := gin.Default()
	r.POST("/orders", orderHandler.NewOrder)
	r.GET("/orders", orderHandler.GetOrderList)
	r.POST("/orders/status", orderHandler.GetOrderStatus)
	return r
}
