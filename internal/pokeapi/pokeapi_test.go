package pokeapi

import (
	"testing"
	"time"

	"github.com/derjabineli/pokedex/internal/pokecache"
)

func TestExploreLocation(t *testing.T) {
	cache := pokecache.NewCache(30 * time.Second)
	cases := []struct {
		input    string
		expected string
	}{
	  {
		input: "explore h",
		expected: "",
	  },
	  {
		input: "mt-coronet-2f",
	  	expected: "mt-coronet-2f",
	  },
	  {
		input: "canalave-city-area",
	  	expected: "canalave-city-area",
	  },
	}

	for _, c := range cases {
		actual := ExploreLocation(c.input, cache)
		if c.expected != actual.Name {
			t.Errorf("Expected %s; Got %s", c.expected, actual.Name)
        	t.Fail()
		}
	}
}