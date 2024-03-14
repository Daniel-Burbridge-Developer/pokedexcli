package main

import "fmt"

func commandHelp() error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, value := range cliCommandBuilder() {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	fmt.Println("")
	return nil
}
