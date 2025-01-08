package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/derjabineli/pokedex/internal/pokeapi"
	"github.com/derjabineli/pokedex/internal/pokecache"
)

var Commands map[string]cliCommand

type cliCommand struct {
  name string
  description string
  callback func(*Config, string) error
}

type Config struct {
  Next string
  Previous string
  Cache *pokecache.Cache
  Pokemon map[string]pokeapi.Pokemon
}

func commandExit(cfg *Config, param string) error {
  if param != "" {
    return fmt.Errorf("too many values passed")
  }
  fmt.Print("Closing the Pokedex... Goodbye!\n")
  os.Exit(0)
  return nil
}

func commandHelp(cfg *Config, param string) error {
  if param != "" {
    return fmt.Errorf("too many values passed")
  }
  fmt.Println()
  fmt.Print("Welcome to the Pokedex! \n")
  fmt.Print("Usage:\n\n")
  for _, cmd := range getCommands() {
    fmt.Printf("%v: %v \n", cmd.name, cmd.description)
  }
  return nil
}

func commandMap(cfg *Config, param string) error {
  if param != "" {
    return fmt.Errorf("too many values passed")
  }

  result := pokeapi.Location{}

  entry, exists := cfg.Cache.Get(cfg.Next)
  if exists {
    err := json.Unmarshal(entry, &result)
	  if err != nil {
		  fmt.Println(err)
	  }
  } else {
    result = pokeapi.GetLocation(cfg.Next, cfg.Cache)
  }

  printLocations(result)
  
  if cfg.Previous == "" {
    cfg.Previous = cfg.Next
  } else {
    cfg.Previous = result.Previous
  }
  cfg.Next = result.Next

  return nil
}

func commandMapb(cfg *Config, param string) error {
  if param != "" {
    return fmt.Errorf("too many values passed")
  }

  if cfg.Previous == "" {
    fmt.Println("you're on the first page")
    return nil
  } else {
    result := pokeapi.Location{}
    entry, exists := cfg.Cache.Get(cfg.Previous)
    if exists {
      err := json.Unmarshal(entry, &result)
	    if err != nil {
		    fmt.Println(err)
	    }
    } else {
      result = pokeapi.GetLocation(cfg.Previous, cfg.Cache)
    }
    cfg.Previous = result.Previous
    cfg.Next = result.Next
    printLocations(result)
      return nil
  }
}

func printLocations(result pokeapi.Location) {
  for _, location := range result.Results {
    fmt.Println(location.Name)
  }
}

func commandExplore(cfg *Config, param string) error {
  if param == "" {
    fmt.Print("Please provide a location that you'd like to explore \n")
    return nil
  }

  result := pokeapi.LocationData{}

  location, exists := cfg.Cache.Get(param)
  if exists {
    err := json.Unmarshal(location, &result)
	    if err != nil {
		    fmt.Println(err)
	    }
    } else {
      result = pokeapi.ExploreLocation(param, cfg.Cache)
  }

  if result.Name != "" {
    printPokemon(param, result)
    return nil
  } else {
    fmt.Printf("Hmm.. %v doesn't seem to exist\n", param)
    return nil
  }
}

func printPokemon(city string, result pokeapi.LocationData) {
  fmt.Printf("Exploring %v...\n", city)
  fmt.Println("Found Pokemon:")
  for _, pokemon := range result.PokemonEncounters {
    fmt.Printf("- %v\n", pokemon.Pokemon.Name)
  }
}

func commandCatch(cfg *Config, pokemon string) error {
  fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)

  data, err := pokeapi.CatchPokemon(pokemon)
  if err != nil {
    return fmt.Errorf("%v escaped!\n", pokemon)
  } else {
    fmt.Printf("%v was caught!\n", pokemon)
    cfg.Pokemon[pokemon] = data
    return nil
  }
}

func commandInspect(cfg *Config, pokemon string) error {
  data, exists := cfg.Pokemon[pokemon]
  if !exists {
    return fmt.Errorf("you have not caught that pokemon")
  }
  fmt.Printf(`Name: %v 
Height: %v
Weight: %v
`, data.Name, data.Height, data.Weight)

  fmt.Print("Stats:\n")
  for i := range data.Stats {
    fmt.Printf("  -%v: %v\n", data.Stats[i].Stat.Name, data.Stats[i].BaseStat)
  }
  fmt.Print("Types:\n")
  for i := range data.Types {
    fmt.Printf("  - %v\n", data.Types[i].Type.Name)
  }
  return nil
}

func commandPokedex(cfg *Config, param string) error {
  if len(cfg.Pokemon) == 0 {
    return fmt.Errorf("no pokemon in pokedex")
  }
  fmt.Print("Your Pokedex:\n")
  for name := range cfg.Pokemon {
    fmt.Printf("  - %v\n", name)
  }
  return nil
}