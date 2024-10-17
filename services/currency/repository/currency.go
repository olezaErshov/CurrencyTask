package repository

import (
	"CurrencyTask/services/currency/entity"
	"CurrencyTask/services/currency/errorsx"
	"context"
	"database/sql"
	"errors"
)

func (r repository) GetCurrencyByDate(ctx context.Context, date string) (float64, error) {
	var rate float64
	tx, err := r.postgres.Begin()
	if err != nil {
		return 0, err
	}

	go func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := "SELECT rate FROM currency WHERE date = $1"

	row := tx.QueryRowContext(ctx, query, date)

	if err = row.Scan(&rate); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errorsx.RateDoesNotExistError
		}
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return rate, nil
}

func (r repository) GetRateHistory(ctx context.Context, firstDate, lastDate string) ([]entity.Currency, error) {
	var (
		rate float32
		date string
	)
	tx, err := r.postgres.Begin()
	if err != nil {
		return nil, err
	}

	go func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := "SELECT rate,date FROM currency WHERE date BETWEEN $1 AND $2"

	rows, err := r.postgres.Query(query, firstDate, lastDate)
	if err != nil {
		return nil, err
	}

	exchangeRateHistory := make([]entity.Currency, 0)
	for rows.Next() {
		err = rows.Scan(&rate, &date)
		if err != nil {
			return nil, err
		}
		exchangeRateHistory = append(exchangeRateHistory, entity.Currency{Rate: rate, Date: date})
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return exchangeRateHistory, nil
}

func (r repository) SaveTodaysCurrency(ctx context.Context, currency entity.Currency) error {
	tx, err := r.postgres.Begin()
	if err != nil {
		return err
	}

	go func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := "INSERT INTO currency (rate, name, date) VALUES ($1,'usd',$2) RETURNING id"
	row := tx.QueryRowContext(ctx, query, currency.Rate, currency.Date)

	if err = row.Err(); err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
