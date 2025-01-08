package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/derjabineli/pokedex/internal/pokecache"
)

func startRepl() {  
	config := &Config{Next: "https://pokeapi.co/api/v2/location-area", Previous: "", Cache: pokecache.NewCache(30 * time.Second)}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(` ________  ________  ___  __    _______   ________  _______      ___    ___ 
|\   __  \|\   __  \|\  \|\  \ |\  ___ \ |\   ___ \|\  ___ \    |\  \  /  /|
\ \  \|\  \ \  \|\  \ \  \/  /|\ \   __/|\ \  \_|\ \ \   __/|   \ \  \/  / /
 \ \   ____\ \  \\\  \ \   ___  \ \  \_|/_\ \  \ \\ \ \  \_|/__  \ \    / / 
  \ \  \___|\ \  \\\  \ \  \\ \  \ \  \_|\ \ \  \_\\ \ \  \_|\ \  /     \/  
   \ \__\    \ \_______\ \__\\ \__\ \_______\ \_______\ \_______\/  /\   \  
    \|__|     \|_______|\|__| \|__|\|_______|\|_______|\|_______/__/ /\ __\ 
                                                                |__|/ \|__|
____________________________________________________________________________			
																`)

  for {
    fmt.Print("Pokedex > ")
    scanner.Scan()
    
    input := cleanInput(scanner.Text())
    if len(input) == 0 {
      continue
    }

    command, exists  := getCommands()[input[0]]
	param := ""
	if len(input) == 2 {
		param = input[1]
	}
    
	if exists {
		err := command.callback(config, param)
		if err != nil{
			fmt.Println(err)
		}
		continue
	} else {
		fmt.Println("Unknown Command")
      continue
    }
  }
}

func cleanInput(input string) []string {
	input = strings.ToLower(input)
	cleaned := strings.Fields(input)
	return cleaned
  }

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
	  "map": {
		name: "map",
		description: "Displays the names of 20 location areas in the Pokemon world",
		callback: commandMap,
	  },
	  "mapb": {
		name: "mapb",
		description: "Display the names of the previously displayed location areas",
		callback: commandMapb,
	  },
	  "explore": {
		name: "explore",
		description: "Displays pokemon at a provided location",
		callback: commandExplore,
	  },
	  "help": {
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	  },
	  "exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	  },
	}
}