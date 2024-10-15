package repository

import (
	"CurrencyTask/services/currency/service"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	postgres *sqlx.DB
}

func NewRepository(db *sqlx.DB) service.Repositorier {
	return repository{
		postgres: db,
	}
}
