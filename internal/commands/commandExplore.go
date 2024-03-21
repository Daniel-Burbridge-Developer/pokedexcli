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

	url := fmt.Sprint("https://pokeapi.co/api/v2/location-area/", *location)

	body, err := pokeClient.ExploreLocation(url)
	if err != nil {
		return config, err
	}

	fmt.Println(body)

	return config, nil
}
