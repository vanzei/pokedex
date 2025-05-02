package main

import (
    "encoding/json"
    "fmt"
    "net/http"

)

var offset = 0
const page = 20
var limit = 0

type LocationAreaResponse struct {
    Results []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
    Next string `json:"next"`
    Prev string `json:"previous"`
}

func commandMap(mapb bool) error {
	if limit > 0 && mapb == false {
		offset = limit
	}
	limit = offset + page
    apiURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?offset=%d&limit=%d", offset, limit)
    resp, err := http.Get(apiURL)
    if err != nil {
        fmt.Println("Error fetching data:", err)
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Printf("Error: received status code %d\n", resp.StatusCode)
        return fmt.Errorf("failed to fetch data")
    }

    var data LocationAreaResponse
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        fmt.Println("Error decoding response:", err)
        return err
    }

    // Display the results
    fmt.Println("Location Areas:")
    for _, result := range data.Results {
        fmt.Printf("- %s\n", result.Name)
    }

    return nil
}

func commandMapb() error {
    if offset - page >= 0 {
		offset -= page
	} else {
		fmt.Println("No previous pages available.")
		return commandMap(true)
	}
	return commandMap(true)
	}
        
    
