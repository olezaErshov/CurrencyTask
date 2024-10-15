package service

import (
	"CurrencyTask/services/currency/config"
	"CurrencyTask/services/currency/entity"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Worker struct {
	currencyService Servicer
	externalURL     string
	RunImmediately  bool
	RunTime         time.Time
}

func NewWorker(currencyService Servicer, workerConfig config.WorkerConfig) *Worker {
	return &Worker{
		currencyService: currencyService,
		externalURL:     workerConfig.ExternalUrl,
		RunImmediately:  workerConfig.FetchingOnStart,
	}
}

func (w *Worker) Start() {
	if w.RunImmediately {
		go w.pullData()
	}

	go func() {
		for {
			nextRun := w.getNextRunTime(time.Now())

			time.Sleep(time.Until(nextRun))

			go w.pullData()
		}
	}()
}

func (w *Worker) getNextRunTime(currentTime time.Time) time.Time {
	nextRun := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), w.RunTime.Hour(), w.RunTime.Minute(), 0, 0, currentTime.Location())

	if currentTime.After(nextRun) {
		nextRun = nextRun.Add(24 * time.Hour)
	}

	return nextRun
}

func (w *Worker) pullData() {
	url := fmt.Sprintf(w.externalURL)

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}

	var currencyResp entity.CurrencyResponse
	if err := json.NewDecoder(resp.Body).Decode(&currencyResp); err != nil {
		return
	}

	var currencyData entity.Currency
	currencyData.Rate = currencyResp.Rates["usd"]
	currencyData.Date = currencyResp.Date

	err = w.currencyService.SaveTodaysCurrency(context.TODO(), currencyData)
	if err != nil {
		return
	}
}
