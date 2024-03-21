package commands

import (
	"os"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/pokeapi"
	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func CommandExit(config models.Config, pokeClient *pokeapi.PokeClient) (models.Config, error) {
	os.Exit(0)
	return config, nil
}
