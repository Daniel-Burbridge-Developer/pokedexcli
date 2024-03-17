package commands

import (
	"fmt"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func CommandHelp(config models.Config) (models.Config, error) {
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
