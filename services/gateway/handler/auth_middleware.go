package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func (h Handler) authMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.GetHeader(authHeader)
		if token == "" {
			log.Println("auth middleware: access token not found")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		url := fmt.Sprintf("http://localhost:8082/validate")

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("auth middleware: something went wrong")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		req.Header.Set(authHeader, token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("auth middleware: error iin getting response")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("auth middleware: error in reaing body")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		switch resp.StatusCode {
		case http.StatusOK:
			log.Println("auth middleware: everything is ok")
			c.JSON(http.StatusOK, gin.H{"token": string(body)})
			return
		case http.StatusBadRequest:
			log.Println("auth middleware error:", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
			return
		case http.StatusUnauthorized:
			log.Println("auth middleware error:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		default:
			log.Println("auth middleware error: unexpected status code", resp.StatusCode)
			c.AbortWithStatusJSON(resp.StatusCode, gin.H{"error": "Unexpected response", "body": string(body)})
			return
		}
	}

}
