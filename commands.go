package main

import (
	"fmt"
	"os"
)


type cliCommand struct {
	name		string
	description string
	callback 	func(*config) error
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
	}
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp(config *config) error {
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

func commandMap(config *config) error {
	locationsResp, err := config.pokeapiClient.ListLocations(config.nextLocationURL)
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

func commandMapb(config *config) error {
	locationsResp, err := config.pokeapiClient.ListLocations(config.previousLocationURL)
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
