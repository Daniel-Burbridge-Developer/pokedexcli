package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/commands"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		usrInput := commands.UsrCommandCleaner(scanner.Text())
		usrCommand, err := commands.CliCommandDistributer(usrInput)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		usrCommand.Callback()
	}
}
