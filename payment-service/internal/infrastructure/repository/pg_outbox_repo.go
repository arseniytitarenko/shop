package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"payment/internal/domain"
	"time"
)

type PgOutboxRepo struct {
	db *gorm.DB
}

func NewPgOutboxRepo(db *gorm.DB) *PgOutboxRepo {
	return &PgOutboxRepo{db: db}
}

func (r *PgOutboxRepo) NewOutbox(ctx context.Context, outbox *domain.Outbox) error {
	return r.db.WithContext(ctx).Create(outbox).Error
}

func (r *PgOutboxRepo) GetUnprocessed(ctx context.Context, limit int) ([]domain.Outbox, error) {
	var messages []domain.Outbox
	err := r.db.WithContext(ctx).
		Where("processed_at IS NULL").
		Order("created_at ASC").
		Limit(limit).
		Find(&messages).Error
	return messages, err
}

func (r *PgOutboxRepo) MarkProcessed(ctx context.Context, id uuid.UUID) error {
	now := time.Now()
	return r.db.WithContext(ctx).
		Model(&domain.Outbox{}).
		Where("id = ?", id).
		Update("processed_at", &now).Error
}
