package main

import (
	"errors"
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
		"map": {
			name:        "map",
			description: "Scrolls through pokemon locations 20 at a time",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapB",
			description: "Scrolls backwards through pokemon locations 20 at a time or displays an error if no previous locations",
			callback:    commandMapB,
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
