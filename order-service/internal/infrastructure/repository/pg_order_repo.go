package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"order/internal/domain"
)

type PgOrderRepo struct {
	db *gorm.DB
}

func NewPgOrderRepo(db *gorm.DB) *PgOrderRepo {
	return &PgOrderRepo{db: db}
}

func (o *PgOrderRepo) NewOrder(ctx context.Context, order *domain.Order) error {
	return o.db.WithContext(ctx).Create(order).Error
}

func (o *PgOrderRepo) GetOrderList(ctx context.Context, userID uuid.UUID) ([]domain.Order, error) {
	var orders []domain.Order
	err := o.db.WithContext(ctx).Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (o *PgOrderRepo) GetOrder(ctx context.Context, orderID uuid.UUID) (*domain.Order, error) {
	var order domain.Order
	if err := o.db.WithContext(ctx).Where("order_id = ?", orderID).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *PgOrderRepo) SaveOrder(ctx context.Context, order *domain.Order) error {
	return o.db.WithContext(ctx).Save(order).Error
}
