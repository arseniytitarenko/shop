package repository

import (
	"context"
	"gorm.io/gorm"
	"payment/internal/application/port/out"
)

type GormTxRepo struct {
	accountRepo out.AccountRepo
	inboxRepo   out.InboxRepo
	outboxRepo  out.OutboxRepo
}

func (r *GormTxRepo) AccountRepo() out.AccountRepo {
	return r.accountRepo
}

func (r *GormTxRepo) InboxRepo() out.InboxRepo {
	return r.inboxRepo
}

func (r *GormTxRepo) OutboxRepo() out.OutboxRepo {
	return r.outboxRepo
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
			accountRepo: NewPgAccountRepo(tx),
			inboxRepo:   NewPgInboxRepo(tx),
			outboxRepo:  NewPgOutboxRepo(tx),
		}
		return fn(ctx, txRepo)
	})
}
