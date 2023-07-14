package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// List all areas in groups of 20
func (client *Client) ListAreas(pageURL *string) (AreasResponse, error) {
	url := baseURL + "/location-area"

	// if pageURL is not empty assign to url
	if pageURL != nil {
		url = *pageURL
	}

	// new GET request and if an error is found, return
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreasResponse{}, err
	}

	// send request and store response, if error is found return
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return AreasResponse{}, err
	}
	defer resp.Body.Close()

	// read all data from response body, if error found, return
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return AreasResponse{}, err
	}

	areasResp := AreasResponse{}
	// Unmarshal
	err = json.Unmarshal(dat, &areasResp)
	if err != nil {
		return AreasResponse{}, err
	}

	// If we made it this far, everything worked OK, return locations response and nil error
	return areasResp, nil
}

// List a specified location's areas
func (client *Client) ListLocationsAreas(location string) (LocDetailsResponse, error) {
	url := baseURL + "/location/" + location

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("issue creating new request: ")
		return LocDetailsResponse{}, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		fmt.Println("issue parsing request")
		return LocDetailsResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("issue reading body")
		return LocDetailsResponse{}, err
	}

	areaResp := LocDetailsResponse{}

	err = json.Unmarshal(dat, &areaResp)
	if err != nil {
		fmt.Println("issue unmarshaling response")
		return LocDetailsResponse{}, err
	}

	return areaResp, nil
}

// List all Locations in groups of 20
func (client *Client) ListLocations(pageURL *string) (LocsResponse, error) {
	// assign base URL + location to url
	url := baseURL + "/location"

	// if pageURL is not empty assign to url
	if pageURL != nil {
		url = *pageURL
	}

	// new GET request and if an error is found, return
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocsResponse{}, err
	}

	// send request and store response, if error is found return
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return LocsResponse{}, err
	}
	defer resp.Body.Close()

	// read all data from response body, if error found, return
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocsResponse{}, err
	}

	locationsResp := LocsResponse{}
	// Unmarshal
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocsResponse{}, err
	}

	// If we made it this far, everything worked OK, return locations response and nil error
	return locationsResp, nil
}
