package main

import (
	"fmt"
	"os"
)

func commandExit(cf *config) error {
	fmt.Println("Exiting Pokedex...")
	os.Exit(0)
	return nil
}

func commandExitSpec(cf *config, spec string) error {
	fmt.Printf("Exiting Pokedex... %s", spec)
	os.Exit(0)
	return nil
}
