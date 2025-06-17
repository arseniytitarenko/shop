package in

import (
	"context"
	"github.com/google/uuid"
	"payment/internal/domain"
)

type AccountUseCase interface {
	NewAccount(ctx context.Context, userID uuid.UUID) error
	GetAccount(ctx context.Context, userID uuid.UUID) (*domain.Account, error)
	ReplenishAccount(ctx context.Context, userID uuid.UUID, amount uint) error
}
