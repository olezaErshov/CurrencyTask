package handler

import (
	"CurrencyTask/services/currency/errorsx"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type rateResponse struct {
	Rate float64 `json:"rate"`
}

func (h Handler) RateByDay(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()

	date := c.Query("date")
	if date == "" {
		log.Println("date is empty")
		errorText(c.Writer, "something went wrong", http.StatusBadRequest)
		return
	}

	exchangeRate, err := h.service.GetCurrencyByDate(ctx, date)
	if err != nil {
		switch {
		case errors.Is(err, errorsx.RateDoesNotExistError):
			errorText(c.Writer, "rate from this date doesn't exist", http.StatusNotFound)
			return
		case errors.Is(err, context.DeadlineExceeded):
			errorText(c.Writer, "time limit exceeded", http.StatusInternalServerError)
			return
		default:
			log.Println("getRateByDay handler err:", err)
			errorText(c.Writer, "something went wrong", http.StatusInternalServerError)
			return
		}
	}

	j, err := json.Marshal(rateResponse{Rate: exchangeRate})
	if err != nil {
		log.Println("getRateByDay handler error:", err)
		errorText(c.Writer, "something went wrong", http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	_, err = c.Writer.Write(j)
	if err != nil {
		log.Println("getRateByDay handler error:", err)
		errorText(c.Writer, "something went wrong", http.StatusInternalServerError)
		return
	}
}

func (h Handler) RateHistory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()

	firstDate := c.Query("first_date")
	if firstDate == "" {
		log.Println("rateHistory handler err: first_date is empty")
		errorText(c.Writer, "something went wrong", http.StatusBadRequest)
		return
	}

	lastDate := c.Query("last_date")
	if lastDate == "" {
		log.Println("rateHistory handler err: last_date is empty")
		errorText(c.Writer, "something went wrong", http.StatusBadRequest)
		return
	}

	exchangeRateHistory, err := h.service.GetRateHistory(ctx, firstDate, lastDate)
	if err != nil {
		switch {
		case errors.Is(err, errorsx.RateDoesNotExistError):
			errorText(c.Writer, "rate from this date doesn't exist", http.StatusNotFound)
			return
		case errors.Is(err, context.DeadlineExceeded):
			errorText(c.Writer, "time limit exceeded", http.StatusInternalServerError)
			return
		default:
			log.Println("rateHistory handler err:", err)
			errorText(c.Writer, "something went wrong", http.StatusBadRequest)
			return
		}
	}

	j, err := json.Marshal(exchangeRateHistory)
	if err != nil {
		log.Println("rateHistory handler error:", err)
		errorText(c.Writer, "something went wrong", http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	_, err = c.Writer.Write(j)
	if err != nil {
		log.Println("rateHistory handler error:", err)
		errorText(c.Writer, "something went wrong", http.StatusInternalServerError)
		return
	}
}
