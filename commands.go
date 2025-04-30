package main

import (
	"fmt"
	"os"
	"math/rand"
)


type cliCommand struct {
	name		string
	description string
	callback 	func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return  map[string]cliCommand{
		"exit": {
			name:		 "exit",
			description: "Exit the Pokedex",
			callback: 	 commandExit,
		},
		"help": {
			name:		 "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name:        "map",
			description: "Gets next page of locations",
		   	callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Gets previous page of locations",
			callback: commandMapb,
		},
		"explore": {
			name: 		 "explore",
			description: "Get exploration info for the location",
			callback: commandExplore,
		},
		"catch": {
			name: 		 "catch",
			description: "Attempt to catch pokemon",
			callback: commandCatch,
		},
		"pokedex": {
			name:   	 "pokedex",
			description: "List all pokemon in pokedex",
			callback: commandPokedex,
		},
		"inspect": {
			name:  		 "inspect",
			description: "List Pokemon Details from pokedex",
			callback: commandInspect,
		},
	}
}

func commandExit(config *config, parameters []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp(config *config, parameters []string) error {
	commands := getCommands()
	fmt.Println()
	fmt.Println("List of all commands: ")
	fmt.Println()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()

	return nil
}

func commandMap(config *config, parameters []string) error {
	locationsResp, err := config.pokeapiClient.ListLocations(config.nextLocationURL, &config.pokecache)
	if err != nil {
		return err
	}

	config.nextLocationURL = locationsResp.Next
	config.previousLocationURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(config *config, parameters []string) error {
	locationsResp, err := config.pokeapiClient.ListLocations(config.previousLocationURL, &config.pokecache)
	if err != nil {
		return err
	}

	config.nextLocationURL = locationsResp.Next
	config.previousLocationURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore(config *config, parameters []string) error {
	fmt.Printf("Exploring %s...\n ", parameters[0])
	exploreResp, err := config.pokeapiClient.ExploreLocation(parameters[0], &config.pokecache)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter:= range exploreResp.PokemonEncounters {
		fmt.Println("- ", encounter.Pokemon.Name)
	}
	return nil	
}

func commandCatch(config *config, parameters []string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", parameters[0])

	pokemonResp, err := config.pokeapiClient.GetPokemon(parameters[0], &config.pokecache)
	if err != nil {
		return err
	}

	baseExperience := pokemonResp.BaseExperience

	chance := baseExperience / 50
	
	if (chance == rand.Intn(chance+1)) {
		fmt.Printf("%s was caught!\n", parameters[0])
		config.pokedex[parameters[0]] = pokemonResp
		return nil
	}

	fmt.Printf("%s escaped!\n", parameters[0])
	return nil
}

func commandPokedex(config *config, parameters []string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemonResp := range config.pokedex {
		fmt.Printf("- %s\n", pokemonResp.Name)
	}

	return nil
}

func commandInspect(config *config, parameters []string) error {
	pokemon, ok := config.pokedex[parameters[0]]

	if !ok {
		fmt.Printf("%s is not in the pokedex.\n", parameters[0])
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, pokemon_type := range pokemon.Types {
		fmt.Printf(" - %s\n", pokemon_type.Type.Name)
	}

	return nil
}
