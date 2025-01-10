package main

import (
	"time"
	"github.com/MichalGul/go_pokedexcli/internal/pokeapi"
)

func main() {
	client:=pokeapi.NewClient(5* time.Second)
	config := &config{
		pokeapiClient: client,
		Next: pokeapi.LocationEndpoint,
		Previous: "",
	}
	startRepl(config)
}
