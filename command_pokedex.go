package main

import "fmt"

func commandPokedex(config *config, _ string) error {

	fmt.Println("Your Pokedex:")
	for key, _ := range config.pokedex.OwnedPokemon {
		fmt.Printf("- %s \n", key)
	}
	return nil
}