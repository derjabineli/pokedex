package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
	  input    string
	  expected []string
  }{
    {
      input: "Charmander Bulbasaur PIKACHU",
      expected: []string{"charmander", "bulbasaur", "pikachu"},
    },
    {
      input: "Ivysaur Wartortle Butterfree",
      expected: []string{"ivysaur", "wartortle", "butterfree"},
    },
    {
      input: " piKaChU ",
      expected: []string{"pikachu"},
    },
    {
      input: "Metapod BlAsToise  ",
      expected: []string{"metapod", "blastoise"},
    },
  }

  for _, c := range cases {
	  actual := cleanInput(c.input)
    if len(actual) != len(c.expected) {
      t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
      continue
    }
	  for i := range actual {
		  word := actual[i]
		  expectedWord := c.expected[i]
      if word != expectedWord {
        t.Errorf("Expected %s; Got %s", expectedWord, word)
        t.Fail()
      }
	  }
  }
}
