package service

import "context"

type Currency interface {
	GetCurrencyByDate(ctx context.Context, date string) (float32, error)
}

type Servicer interface {
	CurrencyRepository
}
