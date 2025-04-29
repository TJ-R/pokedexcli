package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokecache"
	"strings"
)

type config struct {
	pokeapiClient pokeapi.Client
	pokecache pokecache.Cache
	nextLocationURL *string
	previousLocationURL *string
}

func startRepl(config *config) {

	fmt.Println("Welcome to the Pokedex!")
	scanner := bufio.NewScanner(os.Stdin)
	for  {
		fmt.Print("Pokedex > ")
		
		scanner.Scan()

		message := scanner.Text()

		if len(message) != 0 {
			cleanMessage := cleanInput(scanner.Text())

			// Use the first word of the message as the command
			command, ok := getCommands()[cleanMessage[0]]
			if ok {
				err := command.callback(config)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}

	}
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text);
	lower := strings.ToLower(trimmed);
	words := strings.Fields(lower);

	return words
}


