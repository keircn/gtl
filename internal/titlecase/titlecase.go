package titlecase

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	MaxInputLength = 10000
)

var (
	ErrInputTooLong   = errors.New("input text exceeds maximum length")
	ErrInvalidUnicode = errors.New("input contains invalid unicode")
	ErrEmptyInput     = errors.New("input cannot be empty")
)

var SmallWords = map[string]bool{
	"a": true, "an": true, "and": true, "as": true, "at": true, "but": true,
	"by": true, "for": true, "if": true, "in": true, "nor": true, "of": true,
	"on": true, "or": true, "the": true, "to": true, "up": true, "yet": true,
	"so": true, "with": true,
}

func ToTitleCase(text string) (string, error) {
	if text == "" {
		return "", ErrEmptyInput
	}

	if len(text) > MaxInputLength {
		return "", ErrInputTooLong
	}

	if !utf8.ValidString(text) {
		return "", ErrInvalidUnicode
	}

	words := strings.Fields(text)
	if len(words) == 0 {
		return "", ErrEmptyInput
	}

	result := make([]string, len(words))

	for i, word := range words {
		titleWord, err := titleWord(word, i == 0 || i == len(words)-1)
		if err != nil {
			return "", err
		}
		result[i] = titleWord
	}

	return strings.Join(result, " "), nil
}

func titleWord(word string, isFirstOrLast bool) (string, error) {
	if word == "" {
		return word, nil
	}

	if !utf8.ValidString(word) {
		return "", ErrInvalidUnicode
	}

	lowerWord := strings.ToLower(word)

	if isFirstOrLast {
		return capitalizeFirst(lowerWord)
	}

	if SmallWords[lowerWord] {
		return lowerWord, nil
	}

	return capitalizeFirst(lowerWord)
}

func capitalizeFirst(word string) (string, error) {
	if word == "" {
		return word, nil
	}

	if !utf8.ValidString(word) {
		return "", ErrInvalidUnicode
	}

	runes := []rune(word)
	if len(runes) == 0 {
		return word, nil
	}

	runes[0] = unicode.ToUpper(runes[0])
	return string(runes), nil
}
