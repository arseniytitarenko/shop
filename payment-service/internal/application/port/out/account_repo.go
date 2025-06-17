package out

import (
	"github.com/google/uuid"
	"payment/internal/domain"
)

type AccountRepo interface {
	NewAccount(account *domain.Account) error
	GetAccount(accountID uuid.UUID) (*domain.Account, error)
	ReplenishAccount(accountID uuid.UUID, amount uint) error
}
