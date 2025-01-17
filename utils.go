package main

import (
	"math/rand"
)

func attemptToCatchPokemon(experience int) bool {
	baseChance := 0.6              // 60% base chance
	minExperience := 50            // Min exp
	reductionPerExperience := 0.001 // Lowering chances by 0.1% for every 1 exp point

	extraExperience := experience - minExperience
	finalChance := baseChance - float64(extraExperience)*reductionPerExperience

	// Chance cant go below 1%
	if finalChance < 0.01 {
		finalChance = 0.01
	}

	// Random throw
	return rand.Float64() < finalChance
}