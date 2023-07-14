package main

import (
	"fmt"
)

func commandHelp(cf *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("---------------Usage---------------")
	for _, cmd := range getCommands() {
		fmt.Printf("%s:\n %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandHelpSpec(cf *config, spec string) error {
	fmt.Println()

	// if nameonly is added, show only the names
	if spec == "nameonly" {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println()
		fmt.Println("---------------Names---------------")
		for _, cmd := range getCommands() {
			fmt.Print("	")
			fmt.Printf("%s\n", cmd.name)
		}
	} else { // else print error message
		fmt.Println("Invalid special command, use detail or help to see command usage info.")
	}

	fmt.Println()

	return nil
}
