package commands

import (
	"errors"
	"os"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/pokeapi"
	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func CommandExit(config models.Config, pokeClient *pokeapi.PokeClient, option *string) (models.Config, error) {
	if option != nil {
		return config, errors.New("this command does not support an additional option")
	}
	os.Exit(0)
	return config, nil
}
