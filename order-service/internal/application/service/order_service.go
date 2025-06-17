package service

import (
	"github.com/google/uuid"
	"order/internal/application/port/out"
	"order/internal/domain"
)

type OrderService struct {
	orderRepo out.OrderRepo
}

func NewOrderService(orderRepo out.OrderRepo) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

func (s *OrderService) NewOrder(userID uuid.UUID, amount uint, description string) (*domain.Order, error) {
	order := &domain.Order{
		UserID:      userID,
		Amount:      amount,
		Description: description,
		OrderID:     uuid.New(),
		Status:      domain.StatusNew,
	}
	err := s.orderRepo.NewOrder(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *OrderService) GetOrderList(userID uuid.UUID) ([]domain.Order, error) {
	return s.orderRepo.GetOrderList(userID)
}

func (s *OrderService) GetOrder(orderID uuid.UUID) (*domain.Order, error) {
	return s.orderRepo.GetOrder(orderID)
}
