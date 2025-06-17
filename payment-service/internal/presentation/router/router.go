package router

import (
	"github.com/gin-gonic/gin"
	"payment/internal/presentation/handler"
)

func SetupRouter(accountHandler *handler.AccountHandler) *gin.Engine {
	r := gin.Default()
	r.POST("/accounts", accountHandler.NewAccount)
	r.GET("/accounts", accountHandler.GetAccount)
	r.POST("/accounts/deposit", accountHandler.ReplenishAccount)
	return r
}
