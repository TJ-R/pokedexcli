package main

import (
	"time"
	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(1 * time.Minute)

	config := &config{
		pokeapiClient: pokeClient,
		pokecache: *pokeCache,
	}

	startRepl(config)
}
