package main

import (
	"fmt"
	"github.com/MichalGul/go_pokedexcli/internal/pokeapi"
)

func commandMapb(config *config, _ string) error {
	prevUrl := config.Previous
	if config.Previous == "" {
		prevUrl = pokeapi.LocationEndpoint
	}
	currentLocations, err := config.pokeapiClient.GetLocations(prevUrl)
	// fmt.Println(prevUrl)
	if err != nil {
		fmt.Printf("error with mapb command %v", err)
		return nil
	}

	//display names
	locationsSlice := currentLocations.Results
	for _, locationName := range locationsSlice {
		fmt.Println(locationName.Name)
	}

	config.Previous = currentLocations.Previous
	config.Next = currentLocations.Next

	return nil
}
