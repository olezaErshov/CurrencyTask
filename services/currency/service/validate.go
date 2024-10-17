package service

import (
	"CurrencyTask/services/currency/errorsx"
	"time"
)

func validateDate(date string) error {
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return errorsx.WrongDateFormatError
	}
	return nil
}

func validateDates(firstDateStr, lastDateStr string) error {
	layout := "2006-01-02"

	firstDate, err := time.Parse(layout, firstDateStr)
	if err != nil {
		return errorsx.WrongDateFormatError
	}

	lastDate, err := time.Parse(layout, lastDateStr)
	if err != nil {
		return errorsx.WrongDateFormatError
	}

	if !firstDate.Before(lastDate) {
		return errorsx.FirstDateEqualOrHigherThenLastDateError
	}

	return nil
}
