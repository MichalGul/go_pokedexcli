package pokeapi

import (
	"net/http"
	"time"
	"github.com/MichalGul/go_pokedexcli/internal/pokecache"
)

//Http client
type Client struct {
	httpClient http.Client
	cache *pokecache.Cache
}

//create new client
func NewClient(timeout time.Duration) Client {
	cache := pokecache.NewCache(timeout)

	return Client {
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: &cache,
	}
}