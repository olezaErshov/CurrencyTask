package handler

import (
	"CurrencyTask/services/gateway/errorsx"
	"time"
)

func validateDate(date string) error {
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return err
	}
	return nil
}

func validateDates(firstDateStr, lastDateStr string) (bool, error) {
	layout := "2006-01-02"

	firstDate, err := time.Parse(layout, firstDateStr)
	if err != nil {
		return false, errorsx.WrongDateFormatError
	}

	lastDate, err := time.Parse(layout, lastDateStr)
	if err != nil {
		return false, errorsx.WrongDateFormatError
	}

	if !firstDate.Before(lastDate) {
		return false, errorsx.FirstDateEqualOrHigherThenLastDateError
	}

	return true, nil
}
