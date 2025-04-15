package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  tacooo  belll  ",
			expected: []string{"tacooo", "belll"},
		},
		{
			input:    "  hello  world  I am herr ",
			expected: []string{"hello", "world", "i", "am", "herr"},
		},
		{
			input:    "superfunnymonkey 2",
			expected: []string{"superfunnymonkey", "2"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Test Failed, expected len did not match")
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Test Failed, actual %s did not match expected %s", actual[i], c.expected[i])
			}
		}
	}
}
