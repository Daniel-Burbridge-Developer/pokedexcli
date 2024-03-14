package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cliCommandBuilder() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

}

func cliCommandDistributer(command string) (cliCommand, error) {
	commands := cliCommandBuilder()
	value, exists := commands[command]
	if exists {
		return value, nil
	} else {
		return cliCommand{name: "", description: "", callback: nil}, errors.New("command not found")
	}
}

func usrCommandCleaner(input string) string {
	return strings.Trim(strings.ToLower(input), "")
}

func commandHelp() error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, value := range cliCommandBuilder() {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	fmt.Println("")
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
