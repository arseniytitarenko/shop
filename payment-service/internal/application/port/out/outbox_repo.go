package out

import (
	"context"
	"github.com/google/uuid"
	"payment/internal/domain"
)

type OutboxRepo interface {
	NewOutbox(ctx context.Context, outbox *domain.Outbox) error
	GetUnprocessed(ctx context.Context, limit int) ([]domain.Outbox, error)
	MarkProcessed(ctx context.Context, id uuid.UUID) error
}
