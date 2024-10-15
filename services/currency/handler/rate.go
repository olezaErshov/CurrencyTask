package handler

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h Handler) RateByDay(c *gin.Context) {
	date := c.Query("date")
	if date == "" {
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	exchangeRate, err := h.service.GetCurrencyByDate(context.TODO(), date)
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

func (h Handler) RateByDaysInterval(c *gin.Context) {
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

	exchangeRateHistory, err := h.service.GetRateHistory(context.TODO(), firstDate, lastDate)
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
