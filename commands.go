package main

import (
	"fmt"
	"os"

	"github.com/derjabineli/pokedex/internal/pokeapi"
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
  next, previous := pokeapi.GetLocation(cfg.Next)
  cfg.Previous = previous
  cfg.Next = next
  return nil
}

func commandMapb(cfg *Config) error {
  if cfg.Previous == "" {
    fmt.Println("you're on the first page")
    return nil
  } else {
    pokeapi.GetLocation(cfg.Previous)
    return nil
  }
}

