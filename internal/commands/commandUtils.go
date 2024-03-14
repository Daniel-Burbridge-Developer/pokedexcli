package commands

import (
	"errors"
	"strings"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func CliCommandBuilder() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"map": {
			Name:        "map",
			Description: "Scrolls through pokemon locations 20 at a time",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapB",
			Description: "Scrolls backwards through pokemon locations 20 at a time or displays an error if no previous locations",
			Callback:    CommandMapB,
		},
	}

}

func CliCommandDistributer(command string) (CliCommand, error) {
	commands := CliCommandBuilder()
	value, exists := commands[command]
	if exists {
		return value, nil
	} else {
		return CliCommand{Name: "", Description: "", Callback: nil}, errors.New("command not found")
	}
}

func UsrCommandCleaner(input string) string {
	return strings.Trim(strings.ToLower(input), "")
}
