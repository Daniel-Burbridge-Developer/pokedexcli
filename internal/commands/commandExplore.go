package commands

import (
	"fmt"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/pokeapi"
	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func CommandExplore(config models.Config, pokeClient *pokeapi.PokeClient, location *string) (models.Config, error) {
	fmt.Println("Nothing broken yet!")
	return config, nil
}
