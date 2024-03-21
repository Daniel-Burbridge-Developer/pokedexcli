package commands

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/pokeapi"
	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func CommandMap(config models.Config, pokeClient *pokeapi.PokeClient) (models.Config, error) {
	locationsData := models.LocationsData{}

	if config.Next == nil {
		return config, errors.New("no next page exists")
	}

	body, err := pokeClient.RequestLocationData(config.Next)
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
