package out

import (
	"context"
	"github.com/google/uuid"
	"payment/internal/domain"
)

type InboxRepo interface {
	NewInbox(ctx context.Context, inbox *domain.Inbox) error
	GetUnprocessed(ctx context.Context, limit int, msgType string) ([]domain.Inbox, error)
	MarkProcessed(ctx context.Context, id uuid.UUID) error
}
