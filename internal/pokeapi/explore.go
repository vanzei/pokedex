package pokeapi

import (
    "context"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

// GetLocationArea fetches details about a specific location area and its Pokémon
func (c *Client) GetLocationArea(ctx context.Context, locationAreaName string) (LocationAreaResponse, error) {
    url := fmt.Sprintf("%s/location-area/%s", baseURL, locationAreaName)

    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return LocationAreaResponse{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return LocationAreaResponse{}, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return LocationAreaResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return LocationAreaResponse{}, err
    }

    var locationArea LocationAreaResponse
    err = json.Unmarshal(body, &locationArea)
    if err != nil {
        return LocationAreaResponse{}, err
    }

    return locationArea, nil
}