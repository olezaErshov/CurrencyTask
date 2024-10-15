package service

import "context"

type CurrencyRepository interface {
	GetCurrencyByDate(ctx context.Context, date string) (float32, error)
}

type Repositorier interface {
	CurrencyRepository
}
