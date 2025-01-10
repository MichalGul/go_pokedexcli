package pokeapi

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


// Struct client method to get locations
func (c *Client) GetLocations(url string) (Locations, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, fmt.Errorf("error creating Get request for %s error: %v",url, err)
	}

	locationResponse, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, fmt.Errorf("error getting response for %s error: %v", url, err)
	}
	defer locationResponse.Body.Close()

	
	if locationResponse.StatusCode > 299{
		return Locations{}, fmt.Errorf("location status code request not OK :%s", locationResponse.Status)
	}

	var locations Locations
	jsonDecoder := json.NewDecoder(locationResponse.Body)
	decodeError := jsonDecoder.Decode(&locations)
	if decodeError != nil {
		return Locations{}, fmt.Errorf("error decoding response data for %s error: %v", url, decodeError)
	}

	return locations, nil


}