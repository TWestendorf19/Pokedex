package main

import (
	"fmt"
)

func commandMap(cf *config) error {
	fmt.Println()
	fmt.Println("Printing out next 20 areas...")
	// get location response
	areasResp, err := cf.pokeAPIClient.ListAreas(cf.nextURL)
	if err != nil {
		return err
	}

	// reassign next and previous location URLs to iterate forward
	cf.nextURL = areasResp.Next
	cf.previousURL = areasResp.Previous

	// print out all areas
	for _, loc := range areasResp.Results {
		fmt.Print("	")
		fmt.Println(loc.Name)
	}

	fmt.Println()

	return nil
}

func commandMapSpec(cf *config, location string) error {
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
