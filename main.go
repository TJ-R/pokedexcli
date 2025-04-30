package main

import (
	"time"
	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(1 * time.Minute)
	pokeDex := make(map[string]pokeapi.PokemonResponse)

	config := &config{
		pokeapiClient: pokeClient,
		pokecache: *pokeCache,
		pokedex: pokeDex,
	}

	startRepl(config)
}
