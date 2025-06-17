package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"payment/internal/application/errs"
	"payment/internal/domain"
)

type PgAccountRepo struct {
	db *gorm.DB
}

func NewPgAccountRepo(db *gorm.DB) *PgAccountRepo {
	return &PgAccountRepo{db: db}
}

func (p *PgAccountRepo) NewAccount(ctx context.Context, account *domain.Account) error {
	return p.db.WithContext(ctx).Create(account).Error
}

func (p *PgAccountRepo) GetAccount(ctx context.Context, accountID uuid.UUID) (*domain.Account, error) {
	var account domain.Account
	err := p.db.WithContext(ctx).Where("user_id = ?", accountID).First(&account).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errs.ErrAccountNotFound
	}
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (p *PgAccountRepo) SaveAccount(ctx context.Context, account *domain.Account) error {
	return p.db.WithContext(ctx).Save(account).Error
}
