package in

import (
	"github.com/google/uuid"
	"payment/internal/domain"
)

type AccountUseCase interface {
	NewAccount(userID uuid.UUID) error
	GetAccount(userID uuid.UUID) (*domain.Account, error)
	ReplenishAccount(userID uuid.UUID, amount uint) error
}
