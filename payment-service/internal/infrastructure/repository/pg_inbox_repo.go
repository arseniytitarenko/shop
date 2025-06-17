package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"payment/internal/domain"
	"time"
)

type PgInboxRepo struct {
	db *gorm.DB
}

func NewPgInboxRepo(db *gorm.DB) *PgInboxRepo {
	return &PgInboxRepo{db: db}
}

func (r *PgInboxRepo) NewInbox(ctx context.Context, inbox *domain.Inbox) error {
	return r.db.WithContext(ctx).Create(inbox).Error
}

func (r *PgInboxRepo) GetUnprocessed(ctx context.Context, limit int, msgType string) ([]domain.Inbox, error) {
	var messages []domain.Inbox
	err := r.db.WithContext(ctx).
		Where("processed_at IS NULL").
		Where("type = ?", msgType).
		Order("created_at ASC").
		Limit(limit).
		Find(&messages).Error
	return messages, err
}

func (r *PgInboxRepo) MarkProcessed(ctx context.Context, id uuid.UUID) error {
	now := time.Now()
	return r.db.WithContext(ctx).
		Model(&domain.Inbox{}).
		Where("id = ?", id).
		Update("processed_at", &now).Error
}
