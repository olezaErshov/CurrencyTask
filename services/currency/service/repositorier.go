package service

import (
	"CurrencyTask/services/currency/entity"
	"context"
)

type CurrencyRepository interface {
	GetCurrencyByDate(ctx context.Context, date string) (float64, error)
	GetRateHistory(ctx context.Context, firstDate, lastDate string) ([]entity.Currency, error)
	SaveTodaysCurrency(ctx context.Context, currency entity.Currency) error
}

type Repositorier interface {
	CurrencyRepository
}
