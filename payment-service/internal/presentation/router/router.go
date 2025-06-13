package router

import (
	"github.com/gin-gonic/gin"
	"payment/internal/presentation/handler"
)

func SetupRouter(accountHandler *handler.AccountHandler) *gin.Engine {
	r := gin.Default()
	r.POST("/account", accountHandler.NewAccount)
	r.GET("/account", accountHandler.GetAccount)
	return r
}
