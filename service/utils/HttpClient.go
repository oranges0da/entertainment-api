package utils

import (
	"net/http"
	"time"
)

func HttpClient() *http.Client {
	// create a new client
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	return client
}
