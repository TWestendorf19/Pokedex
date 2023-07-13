package main

import (
	"fmt"
)

func commandMap(cf *config) error {
	// get location response
	locationsResp, err := cf.pokeAPIClient.ListLocations(cf.nextURL)
	if err != nil {
		return err
	}

	// reassign next and previous location URLs
	cf.nextURL = locationsResp.Next
	cf.previousURL = locationsResp.Previous

	// print out all locations in area
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapSpec(cf *config, location string) error {
	fmt.Printf("Printing out all areas in location: '%s'...", location)
	*cf.currentURL = "https://pokeapi.co/api/v2/location-area/" + location
	// get location response
	locationsResp, err := cf.pokeAPIClient.ListLocations(cf.currentURL)
	if err != nil {
		return err
	}

	// reassign next and previous location URLs
	cf.nextURL = locationsResp.Next
	cf.previousURL = locationsResp.Previous

	// print out all locations in area
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
