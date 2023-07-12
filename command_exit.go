package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Exiting Pokedex...")
	os.Exit(0)
	return nil
}

func commandExitSpec(spec string) error {
	fmt.Printf("Exiting Pokedex... %s", spec)
	os.Exit(0)
	return nil
}
