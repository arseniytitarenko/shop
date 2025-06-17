package in

import (
	"github.com/google/uuid"
	"order/internal/domain"
)

type OrderUseCase interface {
	NewOrder(userID uuid.UUID, amount uint, description string) (*domain.Order, error)
	GetOrderList(userID uuid.UUID) ([]domain.Order, error)
	GetOrder(orderID uuid.UUID) (*domain.Order, error)
}
