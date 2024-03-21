package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/commands"
	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/pokeapi"
	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	offset := 0
	config := models.Config{Next: fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", offset), Previous: ""}
	pokeClient := pokeapi.NewClient()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		usrInput := commands.UsrCommandCleaner(scanner.Text())
		usrCommand, err := commands.CliCommandDistributer(usrInput)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		config, err = usrCommand.Callback(config, pokeClient)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
