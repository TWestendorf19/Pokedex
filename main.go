package main

import (
	"time"

	"github.com/TWestendorf19/Pokedex/internal/pokeapi"
)

func main() {
	pokeclient := pokeapi.NewClient(5 * time.Second)
	cf := &config{
		pokeAPIClient: pokeclient,
	}
	startRepl(cf)
}
