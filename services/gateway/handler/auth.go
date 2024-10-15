package handler

import (
	"CurrencyTask/services/gateway/entity"
	"CurrencyTask/services/gateway/errorsx"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func (h Handler) signIn(c *gin.Context) {
	var input entity.User
	if err := c.ShouldBind(&input); err != nil {
		log.Println(err)
		errorText(c.Writer, "Something went wrong", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByCreds(input.Login, input.Password)
	if err != nil {
		switch err {
		case errorsx.UserDoesNotExistError:
			errorText(c.Writer, "User not found", http.StatusInternalServerError)
			return
		default:
			log.Println("SignIn handler error:", err)
			errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}

	url := fmt.Sprintf("http://0.0.0.0:8080/generate?login=%s", user.Login)

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}

	token := string(body)

	j, err := json.Marshal(token)
	if err != nil {
		log.Println("SignUp handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")

	c.Writer.Header().Set(authHeader, token)

	c.Writer.WriteHeader(http.StatusOK)
	_, err = c.Writer.Write(j)
	if err != nil {
		log.Println("SignIn handler error:", err)
		errorText(c.Writer, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
