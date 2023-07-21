package main

import (
	"errors"
	"fmt"
)

func commandLocations(cf *config) error {

	fmt.Println()
	fmt.Println("Printing next 20 locations...")
	// get location response
	locationsResp, err := cf.pokeAPIClient.ListLocations(cf.nextURL)
	if err != nil {
		return err
	}

	// reassign next and previous location URLs to iterate forward
	cf.nextURL = locationsResp.Next
	cf.previousURL = locationsResp.Previous

	// print out all locations
	for _, loc := range locationsResp.Results {
		fmt.Print(" | ")
		fmt.Println(loc.Name)
	}

	fmt.Println()

	return nil
}

func commandLocationsBack(cf *config, back string) error {

	if back != "back" && back != "Back" {
		fmt.Println()
		return errors.New("invalid special command, use 'locations back' to go back through list")
	}

	if cf.previousURL == nil {
		return errors.New("cannot go back--at start of list")
	}

	fmt.Println()
	fmt.Println("Printing previous 20 locations...")

	// get location response
	locationsResp, err := cf.pokeAPIClient.ListLocations(cf.previousURL)
	if err != nil {
		return err
	}

	// reassign next and previous location URLs to iterate forward
	cf.nextURL = locationsResp.Next
	cf.previousURL = locationsResp.Previous

	// print out all locations
	for _, loc := range locationsResp.Results {
		fmt.Print(" | ")
		fmt.Println(loc.Name)
	}

	fmt.Println()

	return nil
}
