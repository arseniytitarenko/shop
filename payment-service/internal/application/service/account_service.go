package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"payment/internal/application/errs"
	"payment/internal/application/port/out"
	"payment/internal/domain"
)

type AccountService struct {
	accountRepo out.AccountRepo
}

func NewAccountService(accountRepo out.AccountRepo) *AccountService {
	return &AccountService{accountRepo: accountRepo}
}

func (s *AccountService) NewAccount(userID uuid.UUID) error {
	account := &domain.Account{
		UserID: userID,
	}
	err := s.accountRepo.NewAccount(account)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return errs.ErrAccountAlreadyExists
	}
	return err
}

func (s *AccountService) ReplenishAccount(userID uuid.UUID, amount uint) error {
	return s.accountRepo.ReplenishAccount(userID, amount)
}

func (s *AccountService) GetAccount(userID uuid.UUID) (*domain.Account, error) {
	return s.accountRepo.GetAccount(userID)
}
