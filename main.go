package main

import (
	"time"
	"pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)

	config := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(config)
}
