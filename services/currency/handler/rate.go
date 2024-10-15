package handler

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type dateRequest struct {
	Date string `json:"date"`
}

func (h Handler) RateByDay(c *gin.Context) {
	var input dateRequest
	if err := c.ShouldBind(&input); err != nil {
		log.Println(err)
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	exchangeRate, err := h.service.GetCurrencyByDate(context.TODO(), input.Date)
	if err != nil {
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

type IntervalRequest struct {
	FirstDate string `json:"first_date"`
	LastDate  string `json:"last_date"`
}

func (h Handler) RateByDaysInterval(c *gin.Context) {
	var input IntervalRequest
	if err := c.ShouldBind(&input); err != nil {
		log.Println(err)
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	exchangeRateHistory, err := h.service.GetRateHistory(context.TODO(), input.FirstDate, input.LastDate)
	if err != nil {
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
