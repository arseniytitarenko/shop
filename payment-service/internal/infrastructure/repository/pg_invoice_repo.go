package repository

import (
	"gorm.io/gorm"
)

type PgInvoiceRepo struct {
	db *gorm.DB
}

func NewPgInvoiceRepo(db *gorm.DB) *PgInvoiceRepo {
	return &PgInvoiceRepo{db: db}
}
