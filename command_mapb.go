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
	locationsResp, err := cf.pokeAPIClient.ListAreas(cf.previousURL)
	if err != nil {
		return err
	}

	// reassign next and previous location URLs to iterate forward
	cf.nextURL = locationsResp.Next
	cf.previousURL = locationsResp.Previous

	// print out all locations in area
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapbSpec(cf *config, location string) error {
	fmt.Printf("Printing out all areas in location: '%s'...", location)

	// get location response
	locationsResp, err := cf.pokeAPIClient.ListLocations(location)
	if err != nil {
		return err
	}

	// print out all locations in area
	for _, loc := range locationsResp.Areas {
		fmt.Println(loc.Name)
	}

	return nil
}
