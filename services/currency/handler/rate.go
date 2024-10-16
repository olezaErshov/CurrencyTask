package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (h Handler) RateByDay(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()

	date := c.Query("date")
	if date == "" {
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	exchangeRate, err := h.service.GetCurrencyByDate(ctx, date)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			errorText(c.Writer, "time limit exceeded", http.StatusInternalServerError)
			return
		}
		log.Println(err)
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	j, err := json.Marshal(exchangeRate)
	if err != nil {
		log.Println("GetRateByDay handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	_, err = c.Writer.Write(j)
	if err != nil {
		log.Println("GetRateByDay handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func (h Handler) RateByDaysInterval(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()

	firstDate := c.Query("first_date")
	if firstDate == "" {
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	lastDate := c.Query("last_date")
	if lastDate == "" {
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	exchangeRateHistory, err := h.service.GetRateHistory(ctx, firstDate, lastDate)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			errorText(c.Writer, "time limit exceeded", http.StatusInternalServerError)
			return
		}
		log.Println(err)
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	j, err := json.Marshal(exchangeRateHistory)
	if err != nil {
		log.Println("GetRateByDay handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	_, err = c.Writer.Write(j)
	if err != nil {
		log.Println("GetRateByDay handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
