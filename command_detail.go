package main

import (
	"fmt"
)

func commandDetailSpec(cf *config, commandName string) error {
	fmt.Println()
	fmt.Printf("Getting details for command '%s'...", commandName)
	fmt.Println()

	cmd, exists := getCommands()[commandName]

	if exists {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	} else {
		fmt.Printf("Command '%s' not found. For more help use 'help' command.", commandName)
	}
	fmt.Println()
	return nil
}

func commandDetail(cf *config) error {
	fmt.Println()
	fmt.Println("Cannot use command 'detail' without second command to lookup.")
	return nil
}
