package commands

import (
	"errors"
	"strings"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func CliCommandBuilder() map[string]models.CliCommand {
	return map[string]models.CliCommand{
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

func CliCommandDistributer(command string) (models.CliCommand, error) {
	commands := CliCommandBuilder()
	value, exists := commands[command]
	if exists {
		return value, nil
	} else {
		return models.CliCommand{Name: "", Description: "", Callback: nil}, errors.New("command not found")
	}
}

func UsrCommandCleaner(input string) string {
	return strings.Trim(strings.ToLower(input), "")
}
