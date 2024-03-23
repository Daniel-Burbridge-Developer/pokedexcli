package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/pokeapi"
	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/pokedex"
	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func CommandCatch(config models.Config, pokeClient *pokeapi.PokeClient, pokemon *string) (models.Config, error) {
	if pokemon == nil {
		return config, errors.New("this command requires a pokemon")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", *pokemon)
	pokemonData := pokedex.Pokemon{}

	url := fmt.Sprint("https://pokeapi.co/api/v2/pokemon/", *pokemon)

	body, err := pokeClient.RequestPokemonData(url)

	if err != nil {
		return config, err
	}

	// I think this is actually an error, since the data I'm getting back could be the direct link? But it's working
	json.Unmarshal(body, &pokemonData)

	baseXP := pokemonData.BaseExperience
	chanceToCatch := 100 - (baseXP * 2 / 10)
	if chanceToCatch < 20 {
		chanceToCatch = 20
	}

	if chanceToCatch > 95 {
		chanceToCatch = 95
	}

	catchThresholdToBeat := 100 - chanceToCatch

	throwStrength := rand.Intn(100) + 1

	if throwStrength > catchThresholdToBeat {
		fmt.Printf("%s was caught!\n", pokemonData.Name)
		_, ok := pokeClient.PokeDex.Get(pokemonData.Name)
		if !ok {
			pokeClient.PokeDex.Add(pokemonData.Name, pokemonData)
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemonData.Name)
	}

	return config, nil
}
