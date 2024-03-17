package commands

import (
	"os"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func CommandExit(config models.Config) (models.Config, error) {
	os.Exit(0)
	return config, nil
}
