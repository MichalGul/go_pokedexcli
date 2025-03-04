package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/MichalGul/go_pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *config, parameter string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	pokedex pokeapi.Pokedex
	Next string
	Previous string
}


func getCommandsRegister() map[string]cliCommand {
	return map[string]cliCommand{"exit": {
		name:        "exit",
		description: "Exit the pokedex",
		callback:    commandExit,
	},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays next avaliable locations in Pokemon world",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays previous avaliable locations in Pokemon world",
			callback: commandMapb,
		},
		"explore": {
			name: "explore <area-name>",
			description: "Displays pokemons in selected location. Example: explore eterna-forest-area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch <pokemon-name>",
			description: "Attempts to catch pokemon and adds it to pokedex",
			callback: commandCatch,
		},
		"inspect":{
			name: "inspect <pokemon-name>",
			description: "Inspects caught pokemon displaying its stats",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Displays caught pokemons",
			callback: commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	cleanedResult := strings.Fields(output)
	return cleanedResult
}

func startRepl(config *config) {

	inputScanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		inputScanner.Scan()

		err := inputScanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		userInput := cleanInput(inputScanner.Text())

		if len(userInput) == 0 {
			continue
		}
		commandName := userInput[0]
		command, ok := getCommandsRegister()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		// Extract second argument (if it exists)
		secondArg := ""
		if len(userInput) > 1 {
			secondArg = userInput[1]
		}

		comand_err := command.callback(config, secondArg)
		if comand_err != nil {
			fmt.Printf("Command %s failed %v \n", commandName, comand_err)
		}
	}

}
