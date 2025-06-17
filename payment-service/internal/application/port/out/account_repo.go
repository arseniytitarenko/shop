package out

import (
	"context"
	"github.com/google/uuid"
	"payment/internal/domain"
)

type AccountRepo interface {
	NewAccount(ctx context.Context, account *domain.Account) error
	GetAccount(ctx context.Context, accountID uuid.UUID) (*domain.Account, error)
	SaveAccount(ctx context.Context, account *domain.Account) error
}
