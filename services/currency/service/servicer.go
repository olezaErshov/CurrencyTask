package service

import (
	"CurrencyTask/services/currency/entity"
	"context"
)

type Currency interface {
	GetCurrencyByDate(ctx context.Context, date string) (float32, error)
	GetRateHistory(ctx context.Context, firstDate, lastDate string) ([]entity.Currency, error)
	SaveTodaysCurrency(ctx context.Context, currency entity.Currency) error
}

type Servicer interface {
	CurrencyRepository
}
