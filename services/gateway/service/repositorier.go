package service

import "CurrencyTask/services/gateway/entity"

type UserRepository interface {
	GetUserByCreds(login, password string) (entity.User, error)
}

type Repositorier interface {
	UserRepository
}
