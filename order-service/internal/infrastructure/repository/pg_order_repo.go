package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"order/internal/domain"
)

type PgOrderRepo struct {
	db *gorm.DB
}

func NewPgOrderRepo(db *gorm.DB) *PgOrderRepo {
	return &PgOrderRepo{db: db}
}

func (o *PgOrderRepo) NewOrder(order *domain.Order) error {
	log.Println("mock")
	return o.db.Create(order).Error
}

func (o *PgOrderRepo) GetOrderList(userID uuid.UUID) ([]domain.Order, error) {
	var orders []domain.Order
	err := o.db.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (o *PgOrderRepo) GetOrder(orderID uuid.UUID) (*domain.Order, error) {
	var order domain.Order
	if err := o.db.Where("order_id = ?", orderID).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
