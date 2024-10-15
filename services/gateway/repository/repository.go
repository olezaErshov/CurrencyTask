package repository

import (
	"CurrencyTask/services/gateway/entity"
	"CurrencyTask/services/gateway/errorsx"
	"CurrencyTask/services/gateway/service"
)

type repository struct{}

func NewRepository() service.Repositorier {
	return repository{}
}

func (r repository) GetUserByCreds(login, password string) (entity.User, error) {
	users := []entity.User{
		{Login: "oleg", Password: "pass1"},
		{Login: "egor", Password: "pass2"},
	}

	for _, user := range users {
		if user.Login == login && user.Password == password {
			return user, nil
		}
	}
	return entity.User{}, errorsx.UserDoesNotExistError
}
