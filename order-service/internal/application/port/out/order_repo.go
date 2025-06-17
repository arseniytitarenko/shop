package out

import (
	"github.com/google/uuid"
	"order/internal/domain"
)

type OrderRepo interface {
	NewOrder(order *domain.Order) error
	GetOrderList(userID uuid.UUID) ([]domain.Order, error)
	GetOrder(orderID uuid.UUID) (*domain.Order, error)
}
