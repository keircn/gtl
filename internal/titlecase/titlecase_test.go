package titlecase

import (
	"strings"
	"testing"
)

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
			result, err := ToTitleCase(tt.input)
			if err != nil {
				t.Errorf("ToTitleCase(%q) returned unexpected error: %v", tt.input, err)
				return
			}
			if result != tt.expected {
				t.Errorf("ToTitleCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToTitleCaseErrors(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedErr error
	}{
		{
			name:        "empty string",
			input:       "",
			expectedErr: ErrEmptyInput,
		},
		{
			name:        "whitespace only",
			input:       "   ",
			expectedErr: ErrEmptyInput,
		},
		{
			name:        "input too long",
			input:       strings.Repeat("a", MaxInputLength+1),
			expectedErr: ErrInputTooLong,
		},
		{
			name:        "invalid unicode",
			input:       "hello \xff world",
			expectedErr: ErrInvalidUnicode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToTitleCase(tt.input)
			if err != tt.expectedErr {
				t.Errorf("ToTitleCase(%q) error = %v, want %v", tt.input, err, tt.expectedErr)
			}
			if err != nil && result != "" {
				t.Errorf("ToTitleCase(%q) returned non-empty result on error: %q", tt.input, result)
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
			result, err := titleWord(tt.word, tt.isFirstOrLast)
			if err != nil {
				t.Errorf("titleWord(%q, %t) returned unexpected error: %v", tt.word, tt.isFirstOrLast, err)
				return
			}
			if result != tt.expected {
				t.Errorf("titleWord(%q, %t) = %q, want %q", tt.word, tt.isFirstOrLast, result, tt.expected)
			}
		})
	}
}

func TestTitleWordErrors(t *testing.T) {
	tests := []struct {
		name        string
		word        string
		expectedErr error
	}{
		{
			name:        "invalid unicode",
			word:        "hello\xff",
			expectedErr: ErrInvalidUnicode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := titleWord(tt.word, false)
			if err != tt.expectedErr {
				t.Errorf("titleWord(%q) error = %v, want %v", tt.word, err, tt.expectedErr)
			}
			if err != nil && result != "" {
				t.Errorf("titleWord(%q) returned non-empty result on error: %q", tt.word, result)
			}
		})
	}
}

func TestCapitalizeFirst(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "lowercase word",
			input:    "hello",
			expected: "Hello",
		},
		{
			name:     "already capitalized",
			input:    "Hello",
			expected: "Hello",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single character",
			input:    "a",
			expected: "A",
		},
		{
			name:     "unicode character",
			input:    "ñoño",
			expected: "Ñoño",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := capitalizeFirst(tt.input)
			if err != nil {
				t.Errorf("capitalizeFirst(%q) returned unexpected error: %v", tt.input, err)
				return
			}
			if result != tt.expected {
				t.Errorf("capitalizeFirst(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCapitalizeFirstErrors(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedErr error
	}{
		{
			name:        "invalid unicode",
			input:       "hello\xff",
			expectedErr: ErrInvalidUnicode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := capitalizeFirst(tt.input)
			if err != tt.expectedErr {
				t.Errorf("capitalizeFirst(%q) error = %v, want %v", tt.input, err, tt.expectedErr)
			}
			if err != nil && result != "" {
				t.Errorf("capitalizeFirst(%q) returned non-empty result on error: %q", tt.input, result)
			}
		})
	}
}
