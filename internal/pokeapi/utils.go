package pokeapi

import (
	"math/rand"
)

func attemptToCatchPokemon(experience int) bool {
	baseChance := 0.6              // 60% bazowa szansa
	minExperience := 50            // Minimalne doświadczenie
	reductionPerExperience := 0.001 // Zmniejszenie szansy o 0.1% za każdy 1 punkt doświadczenia

	// Oblicz szansę na złapanie
	extraExperience := experience - minExperience
	finalChance := baseChance - float64(extraExperience)*reductionPerExperience

	// Upewnij się, że szansa nie spada poniżej 1%
	if finalChance < 0.01 {
		finalChance = 0.01
	}

	// Rzut losowy
	return rand.Float64() < finalChance
}