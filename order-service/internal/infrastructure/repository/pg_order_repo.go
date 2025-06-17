package repository

import (
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

func (o *PgOrderRepo) NewOrder(order *domain.Order) error {

}
func (o *PgOrderRepo) GetOrderList(userID uuid.UUID) ([]domain.Order, error) {

}

func (o *PgOrderRepo) GetOrder(orderID uuid.UUID) (*domain.Order, error) {

}
