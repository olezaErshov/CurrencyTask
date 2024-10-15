package service

import "CurrencyTask/services/gateway/entity"

type User interface {
	GetUserByCreds(login, password string) (entity.User, error)
}

type Servicer interface {
	User
}
