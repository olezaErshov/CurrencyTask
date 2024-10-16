package handler

import (
	"CurrencyTask/services/gateway/entity"
	"CurrencyTask/services/gateway/errorsx"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type tokenResponse struct {
	Token string `json:"token"`
}

func (h Handler) signIn(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, requestExpiredInSeconds*time.Second)
	defer cancel()

	var input entity.User
	if err := c.ShouldBind(&input); err != nil {
		log.Println(err)
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByCreds(input.Login, input.Password)
	if err != nil {
		switch {
		case errors.Is(err, errorsx.UserDoesNotExistError):
			errorText(c.Writer, "User not found", http.StatusInternalServerError)
			return
		default:
			log.Println("SignIn handler error:", err)
			errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}
	url := fmt.Sprintf("http://auth-generator:8080/generate?login=%s", user.Login)

	respBody, err := requestInAuthService(ctx, url)
	if err != nil {
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}

	response := tokenResponse{
		Token: string(respBody),
	}

	j, err := json.Marshal(response)
	if err != nil {
		log.Println("SignUp handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set(authHeader, response.Token)

	c.Writer.WriteHeader(http.StatusOK)
	_, err = c.Writer.Write(j)
	if err != nil {
		log.Println("SignIn handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func requestInAuthService(ctx context.Context, url string) ([]byte, error) {

	resp, err := executeRequest(ctx, "GET", url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Println(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return body, nil
}
