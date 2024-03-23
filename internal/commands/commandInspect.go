package commands

import (
	"errors"
	"fmt"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/pokeapi"
	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func CommandInspect(config models.Config, pokeClient *pokeapi.PokeClient, pokemon *string) (models.Config, error) {
	if pokemon == nil {
		return config, errors.New("this command requires a pokemon")
	}

	val, ok := pokeClient.PokeDex.Get(*pokemon)

	if !ok {
		fmt.Println("you have not caught that pokemon")
		return config, nil
	}

	// There has to be a better way to do this
	fmt.Printf("Name: %v\n", val.Name)
	fmt.Printf("Height: %v\n", val.Height)
	fmt.Printf("Weight: %v\n", val.Weight)

	// Especially this
	fmt.Printf("Stats:\n")
	for _, statBlock := range val.Stats {
		fmt.Printf("  -%v: %v\n", statBlock.Stat.Name, statBlock.BaseStat)
	}

	// And this
	fmt.Printf("Types:\n")
	for _, typeBlock := range val.Types {
		fmt.Printf("  - %v\n", typeBlock.Type.Name)
	}

	return config, nil
}
