package dto

import (
	"github.com/google/uuid"
	"order/internal/domain"
)

type NewOrderRequest struct {
	UserID      uuid.UUID `json:"user_id"`
	Amount      uint      `json:"amount"`
	Description string    `json:"description"`
}

type NewOrderResponse struct {
	OrderID uuid.UUID `json:"order_id"`
}

type OrderResponse struct {
	OrderID     uuid.UUID     `json:"order_id"`
	Description string        `json:"description"`
	Amount      uint          `json:"amount"`
	Status      domain.Status `json:"status"`
}

type OrderListResponse struct {
	UserID      uuid.UUID       `json:"user_id"`
	OrderLength int             `json:"orders_length"`
	OrderList   []OrderResponse `json:"orders"`
}

type OrderStatusResponse struct {
	Status domain.Status `json:"status"`
}

type OrderListRequest struct {
	UserID uuid.UUID `json:"user_id"`
}

type OrderRequest struct {
	OrderID uuid.UUID `json:"order_id"`
}
