package service

import (
	"CurrencyTask/services/currency/entity"
	"CurrencyTask/services/currency/errorsx"
	"context"
	"log"
)

func (s service) GetCurrencyByDate(ctx context.Context, date string) (float64, error) {
	err := validateDate(date)
	if err != nil {
		return 0, err
	}

	rate, err := s.repository.GetCurrencyByDate(ctx, date)
	if err != nil {
		log.Println("getCurrencyByDate service err:", err)
		return 0, err
	}
	return rate, nil
}

func (s service) GetRateHistory(ctx context.Context, firstDate, lastDate string) ([]entity.Currency, error) {
	isValid, err := validateDates(firstDate, lastDate)
	if err != nil || isValid == false {
		return nil, err
	}

	rateHistory, err := s.repository.GetRateHistory(ctx, firstDate, lastDate)
	if err != nil {
		log.Println("getCurrencyByDate service err:", err)
		return nil, err
	}
	if len(rateHistory) == 0 {
		log.Println("getCurrencyByDate service err: rateHistory is nil")
		return nil, errorsx.RateDoesNotExistError
	}
	return rateHistory, nil
}

func (s service) SaveTodaysCurrency(ctx context.Context, currency entity.Currency) error {
	err := s.repository.SaveTodaysCurrency(ctx, currency)
	if err != nil {
		log.Println("getCurrencyByDate service err:", err)
		return err
	}
	return nil
}
