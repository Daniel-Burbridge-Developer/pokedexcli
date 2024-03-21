package commands

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/pokeapi"
	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func CommandMapB(config models.Config, pokeClient *pokeapi.PokeClient, location *string) (models.Config, error) {
	if location != nil {
		return config, errors.New("this command does not support a location")
	}

	locationsData := models.LocationsData{}

	if config.Previous == nil {
		return config, errors.New("no previous page exists")
	}

	body, err := pokeClient.RequestLocationData(config.Previous)
	if err != nil {
		return config, err
	}

	json.Unmarshal(body, &config)
	json.Unmarshal(body, &locationsData)

	for _, result := range locationsData.Results {
		fmt.Printf("%s\n", result.Name)
	}

	return config, nil
}
