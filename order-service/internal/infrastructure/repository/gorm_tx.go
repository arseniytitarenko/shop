package repository

import (
	"context"
	"gorm.io/gorm"
	"order/internal/application/port/out"
)

type GormTxRepo struct {
	orderRepo  out.OrderRepo
	outboxRepo out.OutboxRepo
	inboxRepo  out.InboxRepo
}

func (r *GormTxRepo) OrderRepo() out.OrderRepo {
	return r.orderRepo
}

func (r *GormTxRepo) OutboxRepo() out.OutboxRepo {
	return r.outboxRepo
}

func (r *GormTxRepo) InboxRepo() out.InboxRepo {
	return r.inboxRepo
}

type GormTxManager struct {
	db *gorm.DB
}

func NewGormTxManager(db *gorm.DB) *GormTxManager {
	return &GormTxManager{db: db}
}

func (t *GormTxManager) Exec(ctx context.Context, fn func(ctx context.Context, tx out.TxRepo) error) error {
	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txRepo := &GormTxRepo{
			orderRepo:  NewPgOrderRepo(tx),
			outboxRepo: NewPgOutboxRepo(tx),
			inboxRepo:  NewPgInboxRepo(tx),
		}
		return fn(ctx, txRepo)
	})
}
