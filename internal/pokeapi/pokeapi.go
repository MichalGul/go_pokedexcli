package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	elem, exists := c.cache.Get(url)
	if exists {
		locations := Locations{}
		unmarshalErr := json.Unmarshal(elem, &locations)
		if unmarshalErr != nil {
			return Locations{}, fmt.Errorf("error unmarshal response data for %s error: %v", url, unmarshalErr)
		}
		return locations, nil
	}

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

	rawBytes, err := io.ReadAll(locationResponse.Body)
	if err != nil {
		return Locations{}, fmt.Errorf("failed to read raw bytes of response %v", err)
	}
	c.cache.Add(url, rawBytes)

	locations := Locations{}
	unmarshalErr := json.Unmarshal(rawBytes, &locations)
	if unmarshalErr != nil {
		return Locations{}, fmt.Errorf("error unmarshal response data for %s error: %v", url, unmarshalErr)
	}

	return locations, nil
}


// // Struct client method to get locations
// func (c *Client) GetLocations(url string) (Locations, error) {

// 	elem, exists := c.cache.Get(url)
// 	if exists {
// 		locations := Locations{}
// 		unmarshalErr := json.Unmarshal(elem, &locations)
// 		if unmarshalErr != nil {
// 			return Locations{}, fmt.Errorf("error unmarshal response data for %s error: %v", url, unmarshalErr)
// 		}
// 		return locations, nil
// 	}

// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return Locations{}, fmt.Errorf("error creating Get request for %s error: %v",url, err)
// 	}

// 	locationResponse, err := c.httpClient.Do(req)
// 	if err != nil {
// 		return Locations{}, fmt.Errorf("error getting response for %s error: %v", url, err)
// 	}
// 	defer locationResponse.Body.Close()

	
// 	if locationResponse.StatusCode > 299{
// 		return Locations{}, fmt.Errorf("location status code request not OK :%s", locationResponse.Status)
// 	}

// 	rawBytes, err := io.ReadAll(locationResponse.Body)
// 	if err != nil {
// 		return Locations{}, fmt.Errorf("failed to read raw bytes of response %v", err)
// 	}
// 	c.cache.Add(url, rawBytes)

// 	//Todo probably needs to be change to parsing io.ReadAll(resp.Body) and json.Unmarshal(&locations)
// 	locations := Locations{}
// 	unmarshalErr := json.Unmarshal(rawBytes, &locations)
// 	if unmarshalErr != nil {
// 		return Locations{}, fmt.Errorf("error unmarshal response data for %s error: %v", url, unmarshalErr)
// 	}
// 	// jsonDecoder := json.NewDecoder(locationResponse.Body)
// 	// decodeError := jsonDecoder.Decode(&locations)
// 	// if decodeError != nil {
// 	// 	return Locations{}, fmt.Errorf("error decoding response data for %s error: %v", url, decodeError)
// 	// }

// 	return locations, nil
// }