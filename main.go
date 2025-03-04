package main

import (
	"time"
	"github.com/MichalGul/go_pokedexcli/internal/pokeapi"
)

func main() {
	client:=pokeapi.NewClient(5* time.Second, time.Minute*5)
	pokedex := pokeapi.Pokedex{
		OwnedPokemon: make(map[string]pokeapi.CaughtPokemon),
	}
	config := &config{
		pokeapiClient: client,
		pokedex: pokedex,
		Next: pokeapi.LocationEndpoint,
		Previous: "",
	}
	
	startRepl(config)
}
