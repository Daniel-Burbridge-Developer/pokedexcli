package commands

import (
	"errors"
	"os"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/pokeapi"
	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func CommandExit(config models.Config, pokeClient *pokeapi.PokeClient, location *string) (models.Config, error) {
	if location != nil {
		return config, errors.New("this command does not support a location")
	}
	os.Exit(0)
	return config, nil
}
