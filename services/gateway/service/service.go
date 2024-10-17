package service

import (
	"CurrencyTask/services/gateway/entity"
	"log"
)

type service struct {
	repository Repositorier
}

func NewService(r Repositorier) Servicer {
	return service{
		repository: r,
	}
}

func (s service) GetUserByCreds(login, password string) (entity.User, error) {
	user, err := s.repository.GetUserByCreds(login, password)
	if err != nil {
		log.Println("getUserByCreds service error: ", err)
		return entity.User{}, err
	}
	return user, nil
}
