package handler

import (
	"CurrencyTask/services/gateway/entity"
	"CurrencyTask/services/gateway/errorsx"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type tokenResponse struct {
	Token string `json:"token"`
}

func (h Handler) signIn(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()

	var input entity.User
	if err := c.ShouldBind(&input); err != nil {
		log.Println("signIn handler error:", err)
		errorText(c.Writer, "something went wrong", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByCreds(input.Login, input.Password)
	if err != nil {
		switch {
		case errors.Is(err, errorsx.UserDoesNotExistError):
			errorText(c.Writer, "user not found", http.StatusNotFound)
			return
		default:
			log.Println("signIn handler error:", err)
			errorText(c.Writer, "something went wrong", http.StatusInternalServerError)
			return
		}
	}
	url := fmt.Sprintf("%s/generate?login=%s", h.cfg.AuthGenerator, user.Login)

	respBody, err := requestInAuthService(ctx, url)
	if err != nil {
		errorText(c.Writer, "something went wrong", http.StatusInternalServerError)
		return
	}

	response := tokenResponse{
		Token: string(respBody),
	}

	j, err := json.Marshal(response)
	if err != nil {
		log.Println("signIn handler error:", err)
		errorText(c.Writer, "something went wrong", http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set(authHeader, response.Token)

	c.Writer.WriteHeader(http.StatusOK)
	_, err = c.Writer.Write(j)
	if err != nil {
		log.Println("signIn handler error:", err)
		errorText(c.Writer, "something went wrong", http.StatusInternalServerError)
		return
	}
}

func requestInAuthService(ctx context.Context, url string) ([]byte, error) {

	resp, err := executeRequest(ctx, "GET", url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("requestInAuthService error:", err)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("requestInAuthService error:", err)
		return nil, err
	}
	return body, nil
}
