package main

import (
	"context"
	"log"
	"order/internal/application/constants"
	"order/internal/application/service"
	"order/internal/domain"
	"order/internal/infrastructure/conn"
	"order/internal/infrastructure/repository"
	"order/internal/infrastructure/worker"
	"order/internal/presentation/handler"
	"order/internal/presentation/router"
)

func main() {
	db := conn.NewPostgres()

	err := db.AutoMigrate(&domain.Order{}, &domain.Outbox{}, &domain.Inbox{})
	if err != nil {
		log.Fatal(err)
	}

	rabbitConn := conn.NewRabbitMQConn()

	pubRepo, err := repository.NewRabbitPub(rabbitConn, constants.ExchangeName)
	if err != nil {
		log.Fatal(err)
	}

	subRepo, err := repository.NewRabbitSub(rabbitConn, constants.ExchangeName, constants.QueueName,
		constants.TopicTypeIn)
	if err != nil {
		log.Fatal(err)
	}

	orderRepo := repository.NewPgOrderRepo(db)
	outboxRepo := repository.NewPgOutboxRepo(db)
	txManager := repository.NewGormTxManager(db)

	accountService := service.NewOrderService(txManager, orderRepo, outboxRepo)

	ctx := context.Background()
	go worker.RunInboxWorker(ctx, txManager, subRepo)
	go worker.RunOutboxWorker(ctx, txManager, pubRepo, constants.OutboxInterval)
	go worker.RunProcessingOrderWorker(ctx, txManager, constants.ProcessingInterval, constants.TopicTypeIn)

	accountHandler := handler.NewOrderHandler(accountService)

	r := router.SetupRouter(accountHandler)
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
