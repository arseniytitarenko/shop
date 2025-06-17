package main

import (
	"context"
	"log"
	"payment/internal/application/constants"
	"payment/internal/application/service"
	"payment/internal/domain"
	"payment/internal/infrastructure/conn"
	"payment/internal/infrastructure/repository"
	"payment/internal/infrastructure/worker"
	"payment/internal/presentation/handler"
	"payment/internal/presentation/router"
)

func main() {
	db := conn.NewPostgres()
	err := db.AutoMigrate(&domain.Account{}, &domain.Outbox{}, &domain.Inbox{})
	if err != nil {
		log.Fatal(err)
	}

	accountRepo := repository.NewPgAccountRepo(db)

	connection := conn.NewRabbitMQConn()
	pubBroker, err := repository.NewRabbitPub(connection, constants.ExchangeName)
	if err != nil {
		log.Fatal(err)
	}
	subBroker, err := repository.NewRabbitSub(connection, constants.ExchangeName, constants.QueueName,
		constants.TopicNameIn)

	txManager := repository.NewGormTxManager(db)

	accountService := service.NewAccountService(txManager, accountRepo)

	ctx := context.Background()
	go worker.RunInboxWorker(ctx, txManager, subBroker)
	go worker.RunProcessingOrderWorker(ctx, txManager, constants.ProcessingInterval,
		constants.TopicNameIn, constants.TopicNameOut)
	go worker.RunOutboxWorker(ctx, txManager, pubBroker, constants.OutboxInterval)

	accountHandler := handler.NewAccountHandler(accountService)

	r := router.SetupRouter(accountHandler)
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
