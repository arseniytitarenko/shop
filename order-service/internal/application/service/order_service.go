package service

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"order/internal/application/constants"
	"order/internal/application/port/out"
	"order/internal/domain"
	"time"
)

type OrderService struct {
	txManager  out.Tx
	orderRepo  out.OrderRepo
	outboxRepo out.OutboxRepo
}

func NewOrderService(tx out.Tx, orderRepo out.OrderRepo, outboxRepo out.OutboxRepo) *OrderService {
	return &OrderService{txManager: tx, orderRepo: orderRepo, outboxRepo: outboxRepo}
}

func (s *OrderService) NewOrder(ctx context.Context, userID uuid.UUID, amount uint, description string) (*domain.Order, error) {
	order := &domain.Order{
		UserID:      userID,
		Amount:      amount,
		Description: description,
		OrderID:     uuid.New(),
		Status:      domain.StatusNew,
	}

	id := uuid.New()
	payloadBytes, _ := json.Marshal(map[string]interface{}{
		"id":       id,
		"order_id": order.OrderID,
		"user_id":  userID,
		"amount":   amount,
	})

	outbox := &domain.Outbox{
		ID:        id,
		Type:      constants.TopicTypeOut,
		Payload:   string(payloadBytes),
		CreatedAt: time.Now(),
	}

	err := s.txManager.Exec(ctx, func(ctx context.Context, tx out.TxRepo) error {
		if err := tx.OrderRepo().NewOrder(ctx, order); err != nil {
			return err
		}
		return tx.OutboxRepo().NewOutbox(ctx, outbox)
	})
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *OrderService) GetOrderList(ctx context.Context, userID uuid.UUID) ([]domain.Order, error) {
	return s.orderRepo.GetOrderList(ctx, userID)
}

func (s *OrderService) GetOrder(ctx context.Context, orderID uuid.UUID) (*domain.Order, error) {
	return s.orderRepo.GetOrder(ctx, orderID)
}
