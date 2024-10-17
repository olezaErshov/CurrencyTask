package errorsx

import "errors"

var (
	RateDoesNotExistError                   = errors.New("rate does not exist from this date")
	WrongDateFormatError                    = errors.New("wrong date format. Date must be in format YYYY-MM-DD")
	FirstDateEqualOrHigherThenLastDateError = errors.New("first date equal or higher than last date")
)
