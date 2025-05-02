package pokeapi

import (
    "context" // Fix for undefined context
    "encoding/json"
    "io"
    "net/http"

    "go.opentelemetry.io/otel"          // Fix for undefined otel
    "go.opentelemetry.io/otel/metric"   // Fix for undefined metric
)

var (
    meterProvider = otel.GetMeterProvider()
    meter         = meterProvider.Meter("pokeapi")
    apiCallCounter, _ = meter.Int64Counter(
        "api_calls",
        metric.WithDescription("Counts the number of API calls made to the PokeAPI"),
    )
)

func (c *Client) ListLocations(ctx context.Context, pageURL *string) (RespShallowLocations, error) {
    apiCallCounter.Add(ctx, 1)
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
