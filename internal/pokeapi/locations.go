package pokeapi

import (
	"net/http"
	"io"
	"encoding/json"
	"pokedexcli/internal/pokecache"
)

func (c *Client) ListLocations(pageURL *string, cache *pokecache.Cache) (LocationAreasResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cachedResp, ok := cache.Get(url)
	if ok {
		locationsResp := LocationAreasResponse{}
		err := json.Unmarshal(cachedResp, &locationsResp)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return locationsResp, nil
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

	cache.Add(url, dat)

	locationsResp := LocationAreasResponse{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return locationsResp, nil
}

func (c *Client) ExploreLocation(areaName string, cache *pokecache.Cache) (ExploreAreaResponse, error) {
	url := baseURL + "/location-area" + "/" + areaName

	cachedResp, ok := cache.Get(url)
	if ok {
		exploreResp := ExploreAreaResponse{}
		err := json.Unmarshal(cachedResp, &exploreResp)
		if err != nil {
			return ExploreAreaResponse{}, err
		}
		return exploreResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ExploreAreaResponse{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ExploreAreaResponse{}, nil
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return ExploreAreaResponse{}, nil
	}

	cache.Add(url, dat)
	
	exploreResp := ExploreAreaResponse{}
	err = json.Unmarshal(dat, &exploreResp)
	if err != nil {
		return ExploreAreaResponse{}, err
	}

	return exploreResp, nil

}

