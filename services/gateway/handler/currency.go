package handler

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetRate(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()

	date := c.Query("date")
	currencyUrl := fmt.Sprintf("%s/rate/date", h.cfg.CurrencyService)

	resp, statusCode, err := GetRateInCurrenService(ctx, currencyUrl, date)
	if err != nil {
		log.Println("getCurrency handler error:", err)
		errorText(c.Writer, "something went wrong", http.StatusBadRequest)
		return
	}
	c.Data(statusCode, "application/json", resp)
}

func (h Handler) GetRateHistory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()

	firstDate := c.Query("first_date")
	lastDate := c.Query("last_date")

	currencyUrl := fmt.Sprintf("%s/rate/history", h.cfg.CurrencyService)

	resp, statusCode, err := GetRateHistoryInCurrenService(ctx, currencyUrl, firstDate, lastDate)
	if err != nil {
		log.Println("getRateHistory handler error:", err)
		errorText(c.Writer, "something went wrong", http.StatusBadRequest)
		return
	}
	c.Data(statusCode, "application/json", resp)
}

func GetRateInCurrenService(ctx context.Context, serviceUrl, date string) ([]byte, int, error) {
	serviceURL, err := url.Parse(serviceUrl)
	if err != nil {
		return nil, 0, err
	}

	queryParams := url.Values{}
	queryParams.Set("date", date)
	serviceURL.RawQuery = queryParams.Encode()

	resp, err := executeRequest(ctx, "GET", serviceURL.String())
	if err != nil {
		return nil, 0, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}

func GetRateHistoryInCurrenService(ctx context.Context, serviceUrl, firstDate, lastDate string) ([]byte, int, error) {
	serviceURL, err := url.Parse(serviceUrl)
	if err != nil {
		return nil, 0, err
	}

	queryParams := url.Values{}
	queryParams.Set("first_date", firstDate)
	queryParams.Set("last_date", lastDate)
	serviceURL.RawQuery = queryParams.Encode()

	resp, err := executeRequest(ctx, "GET", serviceURL.String())
	if err != nil {
		return nil, 0, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}
