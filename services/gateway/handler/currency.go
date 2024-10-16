package handler

import (
	"CurrencyTask/services/gateway/errorsx"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type rateResponse struct {
	Rate float64 `json:"rate"`
}

func (h Handler) GetCurrency(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()

	date := c.Query("date")
	if date == "" {
		log.Println("GetCurrency handler error: date is empty")
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}
	var serviceResponse rateResponse

	resp, err := GetRateInCurrenService(ctx, date)
	if err != nil {
		log.Println("GetCurrency handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}
	serviceResponse.Rate, err = strconv.ParseFloat(string(resp), 64) //TODO надо красиво вернуть в виде rate:12301 и написать метод, предотвращающий повторную запись в бд по несколько раз на дню
	if err != nil {
		log.Println("GetCurrency handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	j, err := json.Marshal(serviceResponse)
	if err != nil {
		log.Println("GetCurrency handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	_, err = c.Writer.Write(j)
	if err != nil {
		log.Println("GetCurrency handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func GetRateInCurrenService(ctx context.Context, date string) ([]byte, error) {
	serviceURL, err := url.Parse("http://currency:8001/api/v1/rate/date")
	if err != nil {
		return nil, err
	}

	queryParams := url.Values{}
	queryParams.Set("date", date)
	serviceURL.RawQuery = queryParams.Encode()

	resp, err := executeRequest(ctx, "GET", serviceURL.String())
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errorsx.CurrencyServiceError
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (h Handler) GetRateHistory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()

	firstDate := c.Query("first_date")
	if firstDate == "" {
		log.Println("GetRateHistory handler error: first_date is empty")
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	lastDate := c.Query("last_date")
	if lastDate == "" {
		log.Println("GetRateHistory handler error: last_date is empty")
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	resp, err := GetRateHistoryInCurrenService(ctx, firstDate, lastDate)
	if err != nil {
		log.Println("GetRateHistory handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}
	c.Data(http.StatusOK, "application/json", resp)
}

func GetRateHistoryInCurrenService(ctx context.Context, firstDate, lastDate string) ([]byte, error) {
	serviceURL, err := url.Parse("http://currency:8001/api/v1/rate/history") //TODO занести эту строку или целиков в конфиг или по частям
	if err != nil {
		return nil, err
	}

	queryParams := url.Values{}
	queryParams.Set("first_date", firstDate)
	queryParams.Set("last_date", lastDate)
	serviceURL.RawQuery = queryParams.Encode()

	resp, err := executeRequest(ctx, "GET", serviceURL.String())
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errorsx.CurrencyServiceError
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
