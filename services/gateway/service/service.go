package service

import "CurrencyTask/services/gateway/entity"

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
		return entity.User{}, err
	}
	return user, nil
}
