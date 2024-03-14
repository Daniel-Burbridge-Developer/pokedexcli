package commands

import "fmt"

func CommandHelp() error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, value := range CliCommandBuilder() {
		fmt.Printf("%v: %v\n", value.Name, value.Description)
	}
	fmt.Println("")
	return nil
}
