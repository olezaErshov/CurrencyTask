package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
)

func (h Handler) GetCurrency(c *gin.Context) {
	date := c.Query("date")
	if date == "" {
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	resp, err := GetRateInCurrenService(date)
	if err != nil {
		log.Println(err)
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	c.Data(http.StatusOK, "application/json", resp)
}

func GetRateInCurrenService(date string) ([]byte, error) {
	serviceURL, err := url.Parse("http://localhost:8001/api/v1/rate/date")
	if err != nil {
		return nil, err
	}

	queryParams := url.Values{}
	queryParams.Set("date", date)
	serviceURL.RawQuery = queryParams.Encode()

	resp, err := http.Get(serviceURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("external service returned non-OK status: %d", resp.StatusCode)
	}

	body, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (h Handler) GetRateHistory(c *gin.Context) {

}
