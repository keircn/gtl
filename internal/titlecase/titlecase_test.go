package titlecase

import "testing"

func TestToTitleCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic sentence",
			input:    "the quick brown fox",
			expected: "The Quick Brown Fox",
		},
		{
			name:     "with small words",
			input:    "the lord of the rings",
			expected: "The Lord of the Rings",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single word",
			input:    "hello",
			expected: "Hello",
		},
		{
			name:     "small word at beginning",
			input:    "a tale of two cities",
			expected: "A Tale of Two Cities",
		},
		{
			name:     "small word at end",
			input:    "what is this for",
			expected: "What Is This For",
		},
		{
			name:     "already capitalized",
			input:    "The Great Gatsby",
			expected: "The Great Gatsby",
		},
		{
			name:     "mixed case",
			input:    "tHe QuIcK bRoWn FoX",
			expected: "The Quick Brown Fox",
		},
		{
			name:     "with conjunctions",
			input:    "war and peace",
			expected: "War and Peace",
		},
		{
			name:     "with prepositions",
			input:    "up in the air",
			expected: "Up in the Air",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToTitleCase(tt.input)
			if result != tt.expected {
				t.Errorf("ToTitleCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestTitleWord(t *testing.T) {
	tests := []struct {
		name          string
		word          string
		isFirstOrLast bool
		expected      string
	}{
		{
			name:          "first word - small word",
			word:          "the",
			isFirstOrLast: true,
			expected:      "The",
		},
		{
			name:          "middle word - small word",
			word:          "the",
			isFirstOrLast: false,
			expected:      "the",
		},
		{
			name:          "last word - small word",
			word:          "of",
			isFirstOrLast: true,
			expected:      "Of",
		},
		{
			name:          "regular word",
			word:          "quick",
			isFirstOrLast: false,
			expected:      "Quick",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := titleWord(tt.word, tt.isFirstOrLast)
			if result != tt.expected {
				t.Errorf("titleWord(%q, %t) = %q, want %q", tt.word, tt.isFirstOrLast, result, tt.expected)
			}
		})
	}
}
