package service

import (
	"CurrencyTask/services/currency/entity"
	"context"
	"log"
)

func (s service) GetCurrencyByDate(ctx context.Context, date string) (float32, error) {
	rate, err := s.repository.GetCurrencyByDate(ctx, date)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return rate, nil
}

func (s service) SaveTodaysCurrency(ctx context.Context, currency entity.Currency) error {
	log.Println(currency)
	err := s.repository.SaveTodaysCurrency(ctx, currency)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
