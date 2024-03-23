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

		"explore": {
			Name:        "explore",
			Description: "Returns a list of all Pokemon that can be found in a location",
			Callback:    CommandExplore,
		},
	}

}

func CliCommandDistributer(command string) (models.CliCommand, *string, error) {
	commands := CliCommandBuilder()
	text := strings.Split(command, " ")
	command = text[0]
	value, exists := commands[command]
	if exists {
		if len(text) > 1 {
			option := text[1]
			return value, &option, nil
		}
		return value, nil, nil
	} else {
		return models.CliCommand{Name: "", Description: "", Callback: nil}, nil, errors.New("command not found")
	}
}

func UsrCommandCleaner(input string) string {
	return strings.Trim(strings.ToLower(input), "")
}
