package main
import (
	"testing"
)

func TestCleanInputs(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "pokedex bulbasaur Charmander PIKACHU",
			expected: []string{"pokedex", "bulbasaur", "charmander", "pikachu"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("lens of slices doesnt match")
		}
		for i, _ := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%v and %v do not match", actual, c.expected)
			}
		}
	}
}
