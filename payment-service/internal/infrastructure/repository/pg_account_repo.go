package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"payment/internal/domain"
)

type PgAccountRepo struct {
	db *gorm.DB
}

func NewPgAccountRepo(db *gorm.DB) *PgAccountRepo {
	return &PgAccountRepo{db: db}
}

func (p *PgAccountRepo) NewAccount(account *domain.Account) error {
	return p.db.Create(account).Error
}

func (p *PgAccountRepo) GetAccount(accountID uuid.UUID) (*domain.Account, error) {
	var account domain.Account
	if err := p.db.Where("user_id = ?", accountID).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (p *PgAccountRepo) ReplenishAccount(accountID uuid.UUID, amount uint) error {
	return p.db.Transaction(func(tx *gorm.DB) error {
		var account domain.Account
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("user_id = ?", accountID).
			First(&account).Error; err != nil {
			return err
		}
		account.Balance += amount
		return tx.Save(&account).Error
	})
}
