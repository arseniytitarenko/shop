package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order/internal/application/dto"
	"order/internal/application/errs"
	"order/internal/application/port/in"
	"order/internal/presentation/response"
)

type OrderHandler struct {
	orderUseCase in.OrderUseCase
}

func NewOrderHandler(orderUseCase in.OrderUseCase) *OrderHandler {
	return &OrderHandler{orderUseCase: orderUseCase}
}

func (h *OrderHandler) NewOrder(c *gin.Context) {
	var orderRequest dto.NewOrderRequest
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		response.HandleError(c, errs.InvalidRequest)
		return
	}

	order, err := h.orderUseCase.NewOrder(orderRequest.UserID, orderRequest.Amount, orderRequest.Description)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, order.OrderID)
}

func (h *OrderHandler) GetOrderList(c *gin.Context) {
	var orderListRequest dto.OrderListRequest
	if err := c.ShouldBindJSON(&orderListRequest); err != nil {
		response.HandleError(c, errs.InvalidRequest)
		return
	}

	orders, err := h.orderUseCase.GetOrderList(orderListRequest.UserID)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	orderResponse := dto.OrderListResponse{
		OrderLength: len(orders),
		UserID:      orderListRequest.UserID,
		OrderList:   []dto.OrderResponse{},
	}
	for _, order := range orders {
		orderResponse.OrderList = append(orderResponse.OrderList, dto.OrderResponse{
			OrderID:     order.OrderID,
			Amount:      order.Amount,
			Description: order.Description,
			Status:      order.Status,
		})
	}
	c.JSON(http.StatusOK, orderResponse)
}

func (h *OrderHandler) GetOrderStatus(c *gin.Context) {
	var orderRequest dto.OrderRequest
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		response.HandleError(c, errs.InvalidRequest)
		return
	}

	order, err := h.orderUseCase.GetOrder(orderRequest.OrderID)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	orderStatusResponse := dto.OrderStatusResponse{
		Status: order.Status,
	}
	c.JSON(http.StatusOK, orderStatusResponse)
}
