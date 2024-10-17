package errorsx

import "errors"

var (
	UserDoesNotExistError                   = errors.New("user does not exist")
	WrongDateFormatError                    = errors.New("wrong date format. Date must be in format YYYY-MM-DD")
	FirstDateEqualOrHigherThenLastDateError = errors.New("first date equal or higher than last date")
)
