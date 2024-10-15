package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResponse struct {
	Error string
}

func errorText(w http.ResponseWriter, error string, code int) {
	h := w.Header()

	h.Set("Content-Type", "application/json")
	h.Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	ErrResp := errorResponse{Error: error}
	j, err := json.Marshal(ErrResp)
	if err != nil {
		log.Println(err)
	}
	_, err = w.Write(j)
	if err != nil {
		log.Println(err)
	}
	log.Println(code, error)
}
