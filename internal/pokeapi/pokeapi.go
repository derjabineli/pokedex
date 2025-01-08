package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/derjabineli/pokedex/internal/pokecache"
)

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetLocation(url string, cache *pokecache.Cache) Location {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Could not find Locations. Error: %v \n", err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	cache.Add(url, body)

	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and body: %s\n", res.StatusCode, body)
	}
	if err != nil {
		fmt.Printf("Could not find Locations. Error: %v \n", err)
	}
	
	result := Location{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func ExploreLocation(location string, cache *pokecache.Cache) LocationData {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + location)
	if err != nil {
		fmt.Printf("Location Explore Failed. Error: %v \n", err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	cache.Add(location, body)

	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and body: %s\n", res.StatusCode, body)
	}
	if err != nil {
		fmt.Printf("Location Explore Failed. Error: %v \n", err)
	}

	result := LocationData{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("Location Explore Failed. Error: %v \n", err)
	}
	return result
}