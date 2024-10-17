package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(authHeader)
		if token == "" {
			log.Println("auth middleware: access token not found")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		body, statusCode, err := h.executeRequestToAuthService(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		}

		switch statusCode {
		case http.StatusOK:
			log.Println("auth middleware: everything is ok")
			return
		case http.StatusBadRequest:
			log.Println("auth middleware error:", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		case http.StatusUnauthorized:
			log.Println("auth middleware error:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		default:
			log.Println("auth middleware error: unexpected status code", statusCode)
			c.AbortWithStatusJSON(statusCode, gin.H{"error": "unexpected response", "body": string(body)})
			return
		}
	}
}

func (h Handler) executeRequestToAuthService(token string) ([]byte, int, error) {
	url := fmt.Sprintf("%s/validate", h.cfg.AuthGenerator)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("executeRequestToAuthService: something went wrong")
		return nil, 0, err
	}

	req.Header.Set(authHeader, token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("executeRequestToAuthService: error in getting response")
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("executeRequestToAuthService: error in reading body")
		return nil, 0, err
	}
	return body, resp.StatusCode, err
}
