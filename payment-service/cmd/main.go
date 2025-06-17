package cmd

import (
	"log"
	"payment/internal/application/service"
	"payment/internal/domain"
	"payment/internal/infrastructure/database"
	"payment/internal/infrastructure/repository"
	"payment/internal/presentation/handler"
	"payment/internal/presentation/router"
)

func main() {
	db := database.NewPostgres()
	err := db.AutoMigrate(&domain.Account{})
	if err != nil {
		log.Fatal(err)
	}

	accountRepo := repository.NewPgAccountRepo(db)

	accountService := service.NewAccountService(accountRepo)

	accountHandler := handler.NewAccountHandler(accountService)

	r := router.SetupRouter(accountHandler)
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
