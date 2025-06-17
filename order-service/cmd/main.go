package main

import (
	"log"
	"order/internal/application/service"
	"order/internal/domain"
	"order/internal/infrastructure/database"
	"order/internal/infrastructure/repository"
	"order/internal/presentation/handler"
	"order/internal/presentation/router"
)

func main() {
	db := database.NewPostgres()
	err := db.AutoMigrate(&domain.Order{})
	if err != nil {
		log.Fatal(err)
	}

	orderRepo := repository.NewPgOrderRepo(db)

	accountService := service.NewOrderService(orderRepo)

	accountHandler := handler.NewOrderHandler(accountService)

	r := router.SetupRouter(accountHandler)
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
