package repository

import (
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
	return 0, nil
}
