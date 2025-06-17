package in

import (
	"context"
	"github.com/google/uuid"
	"order/internal/domain"
)

type OrderUseCase interface {
	NewOrder(ctx context.Context, userID uuid.UUID, amount uint, description string) (*domain.Order, error)
	GetOrderList(ctx context.Context, userID uuid.UUID) ([]domain.Order, error)
	GetOrder(ctx context.Context, orderID uuid.UUID) (*domain.Order, error)
}
