package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	cleanedResult := strings.Fields(output)
	return cleanedResult
}

func startRepl() {

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
		command.callback()
	}

}
