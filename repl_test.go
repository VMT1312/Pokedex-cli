package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  leading and trailing spaces  ",
			expected: []string{"leading", "and", "trailing", "spaces"},
		},
		{
			input:    "HEllo WOrld",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		if len(actual) != len(c.expected) {
			t.Errorf("Expected length %d, got %d for input '%s'", len(c.expected), len(actual), c.input)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word != expectedWord {
				t.Errorf("Expected '%s', got '%s' for input '%s'", expectedWord, word, c.input)
			}
		}
	}
}
