package main

import (
	"fmt"
	"github.com/MichalGul/go_pokedexcli/internal/pokeapi"
)

func commandExplore(config *config, parameter string) error {

	if parameter == "" {
		fmt.Println("Error: No location name provided. Example usage: explore <area_name>")
		return nil
	}

	exploredLocations, err := config.pokeapiClient.ExploreLocation(pokeapi.LocationEndpoint, parameter)
	if err != nil {
		fmt.Printf("error with explore command %v \n", err)
		return err
	}

	//display pokemons
	pokemonEncounterSlice := exploredLocations.PokemonEncounters
	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range pokemonEncounterSlice {
		fmt.Printf(" - %s \n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
