package main

import "fmt"


func commandInspect(config *config, pokemonName string) error {
	if pokemonName == "" {
		return fmt.Errorf("error: No pokemon name provided. Example usage: inspect <pokemon-name>")
		
	}

	pokemon, exists := config.pokedex.OwnedPokemon[pokemonName]
	if !exists {
		fmt.Println("You do not own that pokemon")
		return nil
	}

	pokemon.Display(pokemonName)

	return nil
}