package pokeapi

import (
	"net/http"
	"io"
	"encoding/json"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreasResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationsResp := LocationAreasResponse{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return locationsResp, nil
}

