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
  callback func(*Config) error
}

type Config struct {
  Next string
  Previous string
  Cache *pokecache.Cache
}

func commandExit(*Config) error {
  fmt.Print("Closing the Pokedex... Goodbye!\n")
  os.Exit(0)
  return nil
}

func commandHelp(*Config) error {
  fmt.Println()
  fmt.Print("Welcome to the Pokedex! \n")
  fmt.Print("Usage:\n\n")
  for _, cmd := range getCommands() {
    fmt.Printf("%v: %v \n", cmd.name, cmd.description)
  }
  return nil
}

func commandMap(cfg *Config) error {
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

  printResults(result)

  fmt.Printf("Previous: %v \n Next: %v \n\n", result.Previous, result.Next)
  
  if cfg.Previous == "" {
    cfg.Previous = cfg.Next
  } else {
    cfg.Previous = result.Previous
  }
  cfg.Next = result.Next

  return nil
}

func commandMapb(cfg *Config) error {
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
    printResults(result)
      return nil
  }
}

func printResults(result pokeapi.Location) {
  for _, location := range result.Results {
    fmt.Println(location.Name)
  }
}

