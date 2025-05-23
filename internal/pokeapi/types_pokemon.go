package pokeapi

// PokemonResponse represents the response from the pokemon endpoint
type PokemonResponse struct {
    ID             int    `json:"id"`
    Name           string `json:"name"`
    BaseExperience int    `json:"base_experience"`
    Height         int    `json:"height"`
    Weight         int    `json:"weight"`
    Stats          []struct {
        BaseStat int `json:"base_stat"`
        Effort   int `json:"effort"`
        Stat     struct {
            Name string `json:"name"`
            URL  string `json:"url"`
        } `json:"stat"`
    } `json:"stats"`
    Types []struct {
        Slot int `json:"slot"`
        Type struct {
            Name string `json:"name"`
            URL  string `json:"url"`
        } `json:"type"`
    } `json:"types"`
    Sprites struct {
        FrontDefault string `json:"front_default"`
    } `json:"sprites"`
}