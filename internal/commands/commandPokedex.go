package commands

import (
	"errors"
	"fmt"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/pokeapi"
	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func CommandPokedex(config models.Config, pokeClient *pokeapi.PokeClient, option *string) (models.Config, error) {
	if option != nil {
		return config, errors.New("this command does not support an additional option")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokeClient.PokeDex.Entries {
		fmt.Printf(" - %v\n", pokemon.Name)
	}

	return config, nil
}
