package commands

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/pokeapi"
	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func CommandCatch(config models.Config, pokeClient *pokeapi.PokeClient, pokemon *string) (models.Config, error) {
	if pokemon == nil {
		return config, errors.New("this command requires a pokemon")
	}

	fmt.Println(rand.Intn(100))

	return config, nil
}
