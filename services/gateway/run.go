package gateway

import (
	"CurrencyTask/services/gateway/config"
	"CurrencyTask/services/gateway/handler"
	"CurrencyTask/services/gateway/repository"
	"CurrencyTask/services/gateway/service"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	r := repository.NewRepository()
	s := service.NewService(r)
	h := handler.NewHandler(s)

	router := handler.InitRoutes(&h)

	server := newServer(router, cfg.Server)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func newServer(router http.Handler, cfg config.ServerConfig) *http.Server {
	servAddr := fmt.Sprintf("%s:%v", cfg.Host, cfg.Port)
	return &http.Server{
		Addr:    servAddr,
		Handler: router,
	}
}
