package main

import (
	"errors"
	"fmt"
)

func commandMapb(cf *config) error {

	if cf.previousURL == nil {
		return errors.New("cannot go back--already on first page")
	}
	// get location response
	areasResp, err := cf.pokeAPIClient.ListAreas(cf.previousURL)
	if err != nil {
		return err
	}

	// reassign next and previous location URLs to iterate forward
	cf.nextURL = areasResp.Next
	cf.previousURL = areasResp.Previous

	// print out all areas
	for _, loc := range areasResp.Results {
		fmt.Println(loc.Name)
	}

	fmt.Println()

	return nil
}

func commandMapbSpec(cf *config, location string) error {
	fmt.Println()
	fmt.Printf("Printing out all areas in location: '%s'...\n", location)

	// get location response
	locationsResp, err := cf.pokeAPIClient.ListLocationsAreas(location)
	if err != nil {
		return err
	}

	// print out all areas in location
	for _, loc := range locationsResp.Areas {
		fmt.Print("	")
		fmt.Println(loc.Name)
	}

	fmt.Println()

	return nil
}
