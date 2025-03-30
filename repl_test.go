package main 

import (
	"testing"
)

func TestCleanInput(t* testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: "    multiple    spaces   between     words   ",
			expected: []string{"multiple", "spaces", "between", "words"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Error: cleanInput(%q) returned %d words, expected %d", c.input, len(actual), len(c.expected))
			t.Fail()
		}

		for i := range actual {
			if i >= len(c.expected) {
				break
			}
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Error: cleanInput(%q)[%d] = %q, expected %q", c.input, i, word, expectedWord)
				t.Fail()
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message 
			// and fail the test
		}
	}
}
