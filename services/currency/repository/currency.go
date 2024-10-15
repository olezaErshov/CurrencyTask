package repository

import (
	"CurrencyTask/services/currency/entity"
	"CurrencyTask/services/gateway/errorsx"
	"context"
	"database/sql"
	"errors"
)

func (r repository) GetCurrencyByDate(ctx context.Context, date string) (float32, error) {
	var rate float32
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
			return 0, errorsx.UserDoesNotExistError
		}
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return rate, nil
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

	query := "INSERT INTO currency (rate, name, date) VALUES ($1, $2, $3) RETURNING id"
	row := tx.QueryRowContext(ctx, query, currency.Rate, currency.Usd, currency.Date)

	if err = row.Err(); err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
