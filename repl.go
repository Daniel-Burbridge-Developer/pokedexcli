package main
import (
	"github.com/Daniel-Burbridge-Developer/pokedexcli/commands
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		usrInput := usrCommandCleaner(scanner.Text())
		usrCommand, err := cliCommandDistributer(usrInput)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		usrCommand.callback()
	}
}
