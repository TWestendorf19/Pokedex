package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	// assign base URL + location to url
	url := baseURL + "/location-area"

	// if pageURL is not empty assign to url
	if pageURL != nil {
		url = *pageURL
	}

	// new GET request and if an error is found, return
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// send request and store response, if error is found return
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	// read all data from response body, if error found, return
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	// Unmarshal
	err = json.Unmarshal(dat, &locationsResp)

	if err != nil {
		return RespShallowLocations{}, err
	}

	// If we made it this far, everything worked OK, return locations response and nil error
	return locationsResp, nil
}
