package router

import (
	"github.com/gin-gonic/gin"
	"order/internal/presentation/handler"
	"order/internal/presentation/middleware"
	"time"
)

func SetupRouter(orderHandler *handler.OrderHandler) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.TimeoutMiddleware(3 * time.Second))

	orders := r.Group("/orders")
	{
		orders.POST("", orderHandler.NewOrder)
		orders.GET("", orderHandler.GetOrderList)
		orders.GET("/status", orderHandler.GetOrderStatus)
	}

	return r
}
