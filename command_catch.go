package main

import (
	"fmt"

	"github.com/MichalGul/go_pokedexcli/internal/pokeapi"
)

func commandCatch(config *config, pokemonName string) error {

	if pokemonName == "" {
		return fmt.Errorf("error: No pokemon name provided. Example usage: catch <pokemon-name>")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := config.pokeapiClient.CatchPokemon(pokeapi.PokemonEndpoint, pokemonName)
	if err != nil {
		fmt.Printf("error with catch command %v \n", err)
		return err
	}

	pokeWasCaught := attemptToCatchPokemon(pokemon.BaseExperience)

	if pokeWasCaught {
		fmt.Printf("%s was caught!\n", pokemonName)
		config.pokedex.OwnedPokemon[pokemonName] = pokemon
		fmt.Printf("You may now inspect your %s with inspect command!\n", pokemonName)

	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
