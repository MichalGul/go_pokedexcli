package pokeapi

import (
	"net/http"
	"time"
)

//Http client
type Client struct {
	httpClient http.Client
}

//create new client
func NewClient(timeout time.Duration) Client {
	return Client {
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}