package commands

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/pokeapi"
	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func CommandExplore(config models.Config, pokeClient *pokeapi.PokeClient, location *string) (models.Config, error) {
	if location == nil {
		return config, errors.New("this command requires a location")
	}

	fmt.Printf("Exploring %s...\n", *location)

	exploredLocationsData := models.ExploredLocationData{}

	url := fmt.Sprint("https://pokeapi.co/api/v2/location-area/", *location)

	body, err := pokeClient.ExploreLocation(url)
	if err != nil {
		return config, err
	}

	// I think this is actually an error, since the data I'm getting back could be the direct link? But it's working
	json.Unmarshal(body, &exploredLocationsData)

	fmt.Println("Found Pokemon:")
	for _, result := range exploredLocationsData.PokemonEncounters {
		fmt.Printf(" - %s\n", result.Pokemon.Name)
	}

	return config, nil
}
