package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"payment/internal/application/errs"
	"payment/internal/application/port/out"
	"payment/internal/domain"
)

type AccountService struct {
	accountRepo out.AccountRepo
	txManager   out.Tx
}

func NewAccountService(txManager out.Tx, accountRepo out.AccountRepo) *AccountService {
	return &AccountService{txManager: txManager, accountRepo: accountRepo}
}

func (s *AccountService) NewAccount(ctx context.Context, userID uuid.UUID) error {
	account := &domain.Account{
		UserID: userID,
	}
	err := s.accountRepo.NewAccount(ctx, account)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return errs.ErrAccountAlreadyExists
	}
	return err
}

func (s *AccountService) ReplenishAccount(ctx context.Context, userID uuid.UUID, amount uint) error {
	return s.txManager.Exec(ctx, func(ctx context.Context, tx out.TxRepo) error {
		account, err := tx.AccountRepo().GetAccount(ctx, userID)
		if err != nil {
			return err
		}
		account.Balance += amount
		return tx.AccountRepo().SaveAccount(ctx, account)
	})
}

func (s *AccountService) GetAccount(ctx context.Context, userID uuid.UUID) (*domain.Account, error) {
	return s.accountRepo.GetAccount(ctx, userID)
}
