package service

import (
	"payment/internal/application/port/out"
)

type AccountService struct {
	accountRepo out.AccountRepo
}

func NewAccountService(accountRepo out.AccountRepo) *AccountService {
	return &AccountService{accountRepo: accountRepo}
}
