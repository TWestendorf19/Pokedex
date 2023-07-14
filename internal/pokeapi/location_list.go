package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) ListAreas(pageURL *string) (AreasResponse, error) {
	// assign base URL + location to url
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

	locationsResp := AreasResponse{}
	// Unmarshal
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return AreasResponse{}, err
	}

	// If we made it this far, everything worked OK, return locations response and nil error
	return locationsResp, nil
}

func (client *Client) ListLocations(location string) (LocationsResponse, error) {
	url := baseURL + "/location/" + location

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("issue creating new request: ")
		return LocationsResponse{}, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		fmt.Println("issue parsing request")
		return LocationsResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("issue reading body")
		return LocationsResponse{}, err
	}

	areaResp := LocationsResponse{}

	err = json.Unmarshal(dat, &areaResp)
	if err != nil {
		fmt.Println("issue unmarshaling response")
		return LocationsResponse{}, err
	}

	return areaResp, nil
}
