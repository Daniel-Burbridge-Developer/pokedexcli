package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/commands"
	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	offset := 0
	config := models.Config{Next: fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", offset), Previous: ""}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		usrInput := commands.UsrCommandCleaner(scanner.Text())
		usrCommand, err := commands.CliCommandDistributer(usrInput)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		usrCommand.Callback(config)
		offset, config = handleConfig(offset, usrCommand, config)
	}
}

func handleConfig(offset int, command models.CliCommand, config models.Config) (int, models.Config) {
	// Offset Hard-limits coded in for now, find a better way to do this later.
	// Probably should just have the return statement at the end and have values changed in the if statements but this should work for now.
	if command.Name == "map" {
		if offset < 200 {
			offset += 20
			return offset, models.Config{Next: fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", offset), Previous: fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", (offset - 20))}
		} else {
			return offset, models.Config{Next: "", Previous: "https://pokeapi.co/api/v2/location-area/?offset=180&limit=20"}
		}
	}

	if command.Name == "mapB" {
		if offset > 0 {
			offset -= 20
			return offset, models.Config{Next: fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", offset+20), Previous: fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", offset)}
		} else {
			return offset, models.Config{Next: "https://pokeapi.co/api/v2/location-area/?offset=180&limit=40", Previous: ""}
		}
	}

	return offset, config
}
