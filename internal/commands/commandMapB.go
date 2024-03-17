package commands

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/pokeapi"
	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func CommandMapB(config models.Config) (models.Config, error) {
	locationsData := models.LocationsData{}

	if config.Previous == nil {
		return config, errors.New("no previous page exists")
	}

	body, err := pokeapi.RequestLocationData(config.Previous)
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
