package service

type service struct {
	repository Repositorier
}

func NewService(r Repositorier) Servicer {
	return service{
		repository: r,
	}
}
