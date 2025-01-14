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

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type ExploreResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

func (c *Client) ExploreLocation(url string, location string) (ExploreResponse, error) {
	exploreUrl := url + "/" + location

	elem, exists := c.cache.Get(exploreUrl)
	if exists {
		exploreResponse := ExploreResponse{}
		unmarshalErr := json.Unmarshal(elem, &exploreResponse)
		if unmarshalErr != nil {
			return ExploreResponse{}, fmt.Errorf("error unmarshal response data for %s error: %v", url, unmarshalErr)
		}
		return exploreResponse, nil
	}

	//todo usse cache here get
	fmt.Printf("URL EXPLORE: %s \n", exploreUrl)
	req, err := http.NewRequest("GET", exploreUrl, nil)
	if err != nil {
		return ExploreResponse{}, fmt.Errorf("error creating Get request for %s error: %v", exploreUrl, err)
	}

	exploreResponse, err := c.httpClient.Do(req)
	if err != nil {
		return ExploreResponse{}, fmt.Errorf("error getting response for %s error: %v", url, err)
	}
	defer exploreResponse.Body.Close()

	if exploreResponse.StatusCode > 299{
		return ExploreResponse{}, fmt.Errorf("location status code request not OK :%s", exploreResponse.Status)
	}

	rawBytes, err := io.ReadAll(exploreResponse.Body)
	if err != nil {
		return ExploreResponse{}, fmt.Errorf("failed to read raw bytes of response %v", err)
	}

	c.cache.Add(url, rawBytes)

	exploreLocations := ExploreResponse{}
	unmarshalErr := json.Unmarshal(rawBytes, &exploreLocations)
	if unmarshalErr != nil {
		return ExploreResponse{}, fmt.Errorf("error unmarshal response data for %s error: %v", url, unmarshalErr)
	}

	return exploreLocations, nil

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