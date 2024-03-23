package models

import "github.com/Daniel-Burbridge-Developer/pokedexcli/internal/pokeapi"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(config Config, pokeClient *pokeapi.PokeClient, option *string) (Config, error)
}
