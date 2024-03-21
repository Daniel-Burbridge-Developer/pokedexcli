package commands

import (
	"errors"
	"fmt"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/pokeapi"
	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func CommandExplore(config models.Config, pokeClient *pokeapi.PokeClient, location *string) (models.Config, error) {
	if location == nil {
		return config, errors.New("this command requires a location")
	}

	fmt.Println("Nothing broken yet!")
	return config, nil
}
