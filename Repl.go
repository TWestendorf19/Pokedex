package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/TWestendorf19/Pokedex/internal/pokeapi"
)

type cliCommand struct {
	name         string
	description  string
	callback     func(*config) error
	callbackspec func(*config, string) error
}

type config struct {
	nextURL       *string
	previousURL   *string
	pokeAPIClient pokeapi.Client
}

func startRepl(cf *config) {
	reader := bufio.NewScanner(os.Stdin)

	// infinite loop
	for {
		// print pokedex prompt
		fmt.Print("Pokedex > ")

		// get input from user
		reader.Scan()

		// assign lowercase substrings to words array
		words := cleanInput(reader.Text())

		// if nothing has been entered, go back to top
		if len(words) == 0 {
			continue
		}

		// get name of command from first string in words
		commandName := words[0]

		// if more than the command was typed, store second word in commandSpec
		if len(words) > 1 {
			commandSpec := words[1]

			// search for command by commandName string
			command, exists := getCommands()[commandName]

			// if command is found, call that commands function
			if exists {
				err := command.callbackspec(cf, commandSpec)
				if err != nil {
					fmt.Println(err)
				}
				continue
			} else { // otherwise print unknown command message and prompt for help command
				fmt.Println("Unknown command, type help for command list.")
				continue
			}

		} else { // if only 1 word was given

			// search for command by commandName string
			command, exists := getCommands()[commandName]

			// if command is found, call that commands function
			if exists {
				err := command.callback(cf)
				if err != nil {
					fmt.Println(err)
				}
				continue
			} else { // otherwise print unknown command message and prompt for help command
				fmt.Println("Unknown command, type help for command list.")
				continue
			}
		}
	}
}

func cleanInput(text string) []string {
	// set all characters to lowercase
	output := strings.ToLower(text)

	// gets a slice of substrings separated by whitespace
	words := strings.Fields(output)

	return words
}

func getCommands() map[string]cliCommand {
	// return the map of all cli commands
	return map[string]cliCommand{
		"help": {
			name:         "help",
			description:  "Lists names and descriptions of all commands. Use 'help nameonly' to only show names.",
			callback:     commandHelp,
			callbackspec: commandHelpSpec,
		},
		"detail": {
			name:         "detail",
			description:  "Provides details on a command given after call. Usage: detail <command>",
			callback:     commandDetail,
			callbackspec: commandDetailSpec,
		},
		"map": {
			name:         "map",
			description:  "Displays the names of 20 locations areas in the Pokemon world. Each subsequent call will display the next 20 locations. Using the special command you can pick an area directly (hopefully).",
			callback:     commandMap,
			callbackspec: commandMapSpec,
		},
		"mapb": {
			name:         "mapb",
			description:  "Similar to map command, but iterates backwards through the list. Using the special command is the same as the map special command.",
			callback:     commandMapb,
			callbackspec: commandMapbSpec,
		},
		"exit": {
			name:         "exit",
			description:  "Exits the Pokedex.",
			callback:     commandExit,
			callbackspec: commandExitSpec,
		},
	}
}
