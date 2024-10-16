package errorsx

import "errors"

var (
	UserDoesNotExistError = errors.New("user does not exist")
	CurrencyServiceError  = errors.New("currency service error")
)
