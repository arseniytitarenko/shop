package out

import (
	"context"
	"github.com/google/uuid"
	"order/internal/domain"
)

type OrderRepo interface {
	NewOrder(ctx context.Context, order *domain.Order) error
	GetOrderList(ctx context.Context, userID uuid.UUID) ([]domain.Order, error)
	GetOrder(ctx context.Context, orderID uuid.UUID) (*domain.Order, error)
	SaveOrder(ctx context.Context, order *domain.Order) error
}
