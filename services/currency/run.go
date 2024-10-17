package currency

import (
	"CurrencyTask/services/currency/config"
	"CurrencyTask/services/currency/database"
	"CurrencyTask/services/currency/handler"
	"CurrencyTask/services/currency/repository"
	"CurrencyTask/services/currency/service"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewDB(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("successfully connected to db ")

	r := repository.NewRepository(db)
	s := service.NewService(r)
	w := service.NewWorker(s, cfg.Worker)
	w.Run()

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
