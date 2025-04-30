package pokeapi

import (
	"net/http"
	"io"
	"encoding/json"
	"pokedexcli/internal/pokecache"
)

func (c *Client) GetPokemon(name string, cache *pokecache.Cache) (PokemonResponse, error) {
	url := baseURL + "/pokemon/" + name

	cachedResp, ok := cache.Get(url)
	if ok {
		pokemonResp := PokemonResponse{}
		err := json.Unmarshal(cachedResp, &pokemonResp)
		if err != nil {
			return PokemonResponse{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	cache.Add(url, dat)
	
	pokemonResp := PokemonResponse{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return PokemonResponse{} , err
	}

	return pokemonResp, nil
}
