package commands

import (
	"errors"
	"fmt"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/pokeapi"
	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func CommandHelp(config models.Config, pokeClient *pokeapi.PokeClient, location *string) (models.Config, error) {
	if location != nil {
		return config, errors.New("this command does not support a location")
	}

	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, value := range CliCommandBuilder() {
		fmt.Printf("%v: %v\n", value.Name, value.Description)
	}
	fmt.Println("")
	return config, nil
}
