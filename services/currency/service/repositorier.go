package service

import (
	"CurrencyTask/services/currency/entity"
	"context"
)

type CurrencyRepository interface {
	GetCurrencyByDate(ctx context.Context, date string) (float32, error)
	SaveTodaysCurrency(ctx context.Context, currency entity.Currency) error
}

type Repositorier interface {
	CurrencyRepository
}
