package handler

import (
	"context"
	"log"
	"net/http"
)

func executeRequest(ctx context.Context, httpMethod, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, httpMethod, url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}
