package pokeapi

import (
    "context"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

// GetPokemon fetches details about a specific pokemon
func (c *Client) GetPokemon(ctx context.Context, pokemonName string) (PokemonResponse, error) {
    url := fmt.Sprintf("%s/pokemon/%s", baseURL, pokemonName)

    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return PokemonResponse{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return PokemonResponse{}, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return PokemonResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return PokemonResponse{}, err
    }

    var pokemon PokemonResponse
    err = json.Unmarshal(body, &pokemon)
    if err != nil {
        return PokemonResponse{}, err
    }

    return pokemon, nil
}