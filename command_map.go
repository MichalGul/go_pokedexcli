package main

import (
	"fmt"
)

func commandMap(config *config, _ string) error {
	currentLocations, err := config.pokeapiClient.GetLocations(config.Next)
	if err != nil {
		fmt.Printf("error with map command %v", err)
		return nil
	}

	//display names
	locationsSlice := currentLocations.Results
	for _, locationName := range locationsSlice {
		fmt.Println(locationName.Name)
	}

	config.Next = currentLocations.Next
	config.Previous = currentLocations.Previous

	return nil
}
