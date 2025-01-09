package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello  world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "pikachu charmander",
			expected: []string{"pikachu", "charmander"},
		},
		{
			input:    "Bulbasaur SQUIRTLE",
			expected: []string{"bulbasaur", "squirtle"},
		},
	}

	for _, c := range testCases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Test failed actual length %v, expected length of input %v", len(actual), len(c.expected))
			t.Fail()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Actual workd %s don't match expected %s", word, expectedWord)
				t.Fail()
			}

		}
	}

}
