package errorsx

import "errors"

var (
	RateDoesNotExistError = errors.New("rate does not exist from this date")
)
